package futuapi

import (
	"crypto/aes"
	"crypto/cipher"
)

// ftaesEncrypt encrypts data using the Futu FTAES_ECB variant.
// Padding: null bytes (\x00) to 16-byte boundary.
// Trailer: 16-byte block where last byte = original len % 16 (how many nulls were added).
//
// Layout:
//   padded     = plaintext + (\x00 * padLen)   where padLen = (16 - len%16) % 16
//   encrypted  = AES_ECB_encrypt(padded)
//   result     = encrypted + trailer(16 bytes: \x00*15 + [padLen])
func ftaesEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeEncryptionFailed, "FTAES key must be 16 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeEncryptionFailed, "create AES cipher: "+err.Error())
	}

	// Null-byte padding: append \x00 bytes to reach 16-byte boundary.
	// Special case: if plaintext is empty (0 bytes), we still need a full 16-byte
	// block of null padding so the trailer can encode padLen=16 unambiguously.
	// For non-empty plaintext that is already 16-byte aligned, padLen=0 (no padding).
	padLen := 16 - (len(plaintext) % 16)
	if len(plaintext) > 0 && padLen == 16 {
		padLen = 0
	}
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	// \x00 bytes are already zeroed

	// AES/ECB encrypt — iterate over all 16-byte blocks
	ciphertext := make([]byte, len(padded))
	bs := block.BlockSize()
	for i := 0; i < len(padded); i += bs {
		block.Encrypt(ciphertext[i:], padded[i:])
	}

	// Build 16-byte trailer: 15 null bytes + original remainder as last byte.
	// If len(plaintext) > 0 && len(plaintext) % 16 == 0, trailer[15] = 0 (no padding).
	// If len(plaintext) == 0, trailer[15] = 16 (full 16-byte null block).
	trailer := make([]byte, 16)
	trailer[15] = byte(padLen)

	// Append trailer
	result := append(ciphertext, trailer...)
	return result, nil
}

// ftaesDecrypt decrypts data using the Futu FTAES_ECB variant.
//
//   1. Validate ciphertext length is at least 16 bytes and the encrypted portion
//      (excluding 16-byte trailer) is a multiple of the AES block size.
//   2. Read trailer's last byte → original remainder (how many nulls were padded)
//   3. Remove the 16-byte trailer
//   4. AES/ECB decrypt
//   5. Strip padLen null bytes from the end
//
// Returns ErrNotEncrypted if the data does not match FTAES format (not a full
// block after trailer removal). Callers should fall back to treating the body
// as plaintext when this error is returned.
func ftaesDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeDecryptionFailed, "FTAES key must be 16 bytes")
	}

	if len(ciphertext) < 16 {
		return nil, NewError(CodeDecryptionFailed, "ciphertext too short")
	}

	// FTAES format: encrypted_data (must be n*16 bytes) + 16-byte trailer
	encLen := len(ciphertext) - 16
	if encLen <= 0 || encLen%16 != 0 {
		return nil, ErrNotEncrypted
	}

	// Extract trailer: 16-byte block appended after ciphertext
	// trailer[15] = original remainder (how many null bytes were appended during encrypt)
	trailer := ciphertext[len(ciphertext)-16:]
	padLen := int(trailer[15]) // 0 means original was 16-byte aligned

	// Remove trailer
	ciphertext = ciphertext[:encLen]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeDecryptionFailed, "create AES cipher: "+err.Error())
	}

	// AES/ECB decrypt — iterate over all 16-byte blocks
	plaintext := make([]byte, len(ciphertext))
	bs := block.BlockSize()
	for i := 0; i < len(ciphertext); i += bs {
		block.Decrypt(plaintext[i:], ciphertext[i:])
	}

	// Strip null-byte padding: remove last padLen bytes (which were all \x00)
	if padLen > 0 && padLen <= 16 {
		plaintext = plaintext[:len(plaintext)-padLen]
	}

	return plaintext, nil
}

// AESEncrypt encrypts plaintext using AES/ECB with FTAES null-byte padding scheme.
// Returns ciphertext with appended 16-byte trailer.
func AESEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	return ftaesEncrypt(key, plaintext)
}

// AESDecrypt decrypts ciphertext using AES/ECB with FTAES null-byte padding scheme.
// Input should include the 16-byte trailer; plaintext is returned without padding.
func AESDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	return ftaesDecrypt(key, ciphertext)
}

func aesCBCDecrypt(key []byte, iv []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeDecryptionFailed, "AES_CBC key must be 16 bytes")
	}
	if len(iv) != 16 {
		return nil, NewError(CodeDecryptionFailed, "AES_CBC IV must be 16 bytes")
	}
	if len(ciphertext)%16 != 0 {
		return nil, NewError(CodeDecryptionFailed, "AES_CBC ciphertext must be multiple of 16 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeDecryptionFailed, "create AES cipher: "+err.Error())
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	plaintext := make([]byte, len(ciphertext))
	mode.CryptBlocks(plaintext, ciphertext)
	return plaintext, nil
}

func aesCBCEncrypt(key []byte, iv []byte, plaintext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeEncryptionFailed, "AES_CBC key must be 16 bytes")
	}
	if len(iv) != 16 {
		return nil, NewError(CodeEncryptionFailed, "AES_CBC IV must be 16 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeEncryptionFailed, "create AES cipher: "+err.Error())
	}
	padLen := 16 - (len(plaintext) % 16)
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	for i := len(plaintext); i < len(padded); i++ {
		padded[i] = byte(padLen)
	}
	ciphertext := make([]byte, len(padded))
	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, padded)
	return ciphertext, nil
}

func aes256Encrypt(key []byte, plaintext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, NewError(CodeEncryptionFailed, "AES-256 key must be 32 bytes")
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeEncryptionFailed, "create AES-256 cipher: "+err.Error())
	}
	padLen := 16 - (len(plaintext) % 16)
	if padLen == 16 && len(plaintext) > 0 {
		padLen = 0
	}
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	ciphertext := make([]byte, len(padded))
	for i := 0; i < len(padded); i += 16 {
		block.Encrypt(ciphertext[i:i+16], padded[i:i+16])
	}
	trailer := make([]byte, 16)
	trailer[15] = byte(padLen)
	return append(ciphertext, trailer...), nil
}

func aes256Decrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, NewError(CodeDecryptionFailed, "AES-256 key must be 32 bytes")
	}
	encLen := len(ciphertext) - 16
	if encLen <= 0 || encLen%16 != 0 {
		return nil, ErrNotEncrypted
	}
	trailer := ciphertext[len(ciphertext)-16:]
	padLen := int(trailer[15])
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeDecryptionFailed, "create AES-256 cipher: "+err.Error())
	}
	plaintext := make([]byte, encLen)
	for i := 0; i < encLen; i += 16 {
		block.Decrypt(plaintext[i:i+16], ciphertext[i:i+16])
	}
	if padLen > 0 && padLen <= 16 {
		plaintext = plaintext[:len(plaintext)-padLen]
	}
	return plaintext, nil
}