package futuapi

import (
	"crypto/aes"
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

	// Null-byte padding: append \x00 bytes to reach 16-byte boundary
	padLen := 16 - (len(plaintext) % 16)
	if padLen == 16 {
		padLen = 0 // len(plaintext) is already a multiple of 16
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
	// If len(plaintext) % 16 == 0, trailer[15] = 0 (no padding was added).
	trailer := make([]byte, 16)
	trailer[15] = byte(padLen)

	// Append trailer
	result := append(ciphertext, trailer...)
	return result, nil
}

// ftaesDecrypt decrypts data using the Futu FTAES_ECB variant.
//
//   1. Read trailer's last byte → original remainder (how many nulls were padded)
//   2. Remove the 16-byte trailer
//   3. AES/ECB decrypt
//   4. Strip padLen null bytes from the end
func ftaesDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeDecryptionFailed, "FTAES key must be 16 bytes")
	}

	if len(ciphertext) < 16 {
		return nil, NewError(CodeDecryptionFailed, "ciphertext too short")
	}

	// Extract trailer: 16-byte block appended after ciphertext
	// trailer[15] = original remainder (how many null bytes were appended during encrypt)
	trailer := ciphertext[len(ciphertext)-16:]
	padLen := int(trailer[15]) // 0 means original was 16-byte aligned

	// Remove trailer
	ciphertext = ciphertext[:len(ciphertext)-16]

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