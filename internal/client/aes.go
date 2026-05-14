package futuapi

import (
	"crypto/aes"
)

// ftaesEncrypt encrypts data using the Futu FTAES_ECB variant.
// Padding scheme: PKCS7-style (padding_len repeated padding_len times).
// Trailer: 16-byte block where last byte = original padding_len (1-16, or 0 for 16-byte-aligned).
//
// Layout:
//   encrypted = AES_ECB_encrypt(plaintext + padding_len_repeated * padding_len)
//   result = encrypted + trailer(16 bytes: [0]*15 + [padding_len])
func ftaesEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeEncryptionFailed, "FTAES key must be 16 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeEncryptionFailed, "create AES cipher: "+err.Error())
	}

	// PKCS7-style padding: repeat padding_len value, padding_len times
	padLen := 16 - (len(plaintext) % 16)
	if padLen == 16 && len(plaintext) > 0 {
		padLen = 0
	}
	if padLen == 0 && len(plaintext) == 0 {
		padLen = 16
	}
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	if padLen > 0 {
		for i := len(plaintext); i < len(padded); i++ {
			padded[i] = byte(padLen)
		}
	}

	// AES/ECB encrypt
	ciphertext := make([]byte, len(padded))
	block.Encrypt(ciphertext, padded)

	// Build 16-byte trailer: 15 null bytes + padding_len as last byte
	trailer := make([]byte, 16)
	trailer[15] = byte(padLen)

	// Append trailer
	result := append(ciphertext, trailer...)
	return result, nil
}

// ftaesDecrypt decrypts data using the Futu FTAES_ECB variant.
// Trailer last byte = padding_len (1-16, or 0 means data was 16-byte aligned).
//
//   1. Read trailer's last byte (padding_len)
//   2. Remove the 16-byte trailer
//   3. AES/ECB decrypt
//   4. Strip PKCS7-style padding: remove last padding_len bytes
func ftaesDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeDecryptionFailed, "FTAES key must be 16 bytes")
	}

	if len(ciphertext) < 16 {
		return nil, NewError(CodeDecryptionFailed, "ciphertext too short")
	}

	// Extract the trailer
	trailer := ciphertext[len(ciphertext)-16:]
	padLen := int(trailer[15]) // padding_len (1-16, or 0)

	// Remove trailer
	ciphertext = ciphertext[:len(ciphertext)-16]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeDecryptionFailed, "create AES cipher: "+err.Error())
	}

	// AES/ECB decrypt
	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)

	// Strip PKCS7-style padding
	if padLen > 0 && padLen <= 16 {
		plaintext = plaintext[:len(plaintext)-padLen]
	}

	return plaintext, nil
}

// AESEncrypt encrypts plaintext using AES/ECB with FTAES padding scheme.
// Returns ciphertext with appended 16-byte length trailer.
func AESEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	return ftaesEncrypt(key, plaintext)
}

// AESDecrypt decrypts ciphertext using AES/ECB with FTAES padding scheme.
// Input should include the 16-byte trailer; plaintext is returned without padding.
func AESDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	return ftaesDecrypt(key, ciphertext)
}