package futuapi

import (
	"crypto/aes"
)

// ftaesEncrypt encrypts data using the Futu FTAES_ECB variant.
// This is NOT standard AES/ECB — it's a custom padding scheme:
//
//   1. Pad data to 16-byte boundary with null bytes (\x00)
//   2. AES/ECB encrypt the padded data
//   3. Append a 16-byte trailer: 15 null bytes + original length mod 16
//
// The trailer's last byte records how many bytes of the last 16-byte block
// were real data (0 means the data was already a multiple of 16).
func ftaesEncrypt(key []byte, plaintext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeEncryptionFailed, "FTAES key must be 16 bytes")
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeEncryptionFailed, "create AES cipher: "+err.Error())
	}

	// Pad plaintext to 16-byte boundary with null bytes
	padLen := 16 - (len(plaintext) % 16)
	if padLen == 16 {
		padLen = 0
	}
	padded := make([]byte, len(plaintext)+padLen)
	copy(padded, plaintext)
	// null bytes already zeroed

	// AES/ECB encrypt
	ciphertext := make([]byte, len(padded))
	block.Encrypt(ciphertext, padded)

	// Build 16-byte trailer: 15 null bytes + original length mod 16 as last byte
	trailer := make([]byte, 16)
	trailer[15] = byte(len(plaintext) % 16)

	// Append trailer
	result := append(ciphertext, trailer...)
	return result, nil
}

// ftaesDecrypt decrypts data using the Futu FTAES_ECB variant.
// Reverse of ftaesEncrypt:
//
//   1. Read trailer's last byte (origLen % 16)
//   2. Remove the 16-byte trailer
//   3. AES/ECB decrypt
//   4. Remove null byte padding (if origLen%16 != 0, strip last partial block padding)
func ftaesDecrypt(key []byte, ciphertext []byte) ([]byte, error) {
	if len(key) != 16 {
		return nil, NewError(CodeDecryptionFailed, "FTAES key must be 16 bytes")
	}

	if len(ciphertext) < 16 {
		return nil, NewError(CodeDecryptionFailed, "ciphertext too short")
	}

	// Extract the trailer
	trailer := ciphertext[len(ciphertext)-16:]
	origRemainder := int(trailer[15]) // original data length mod 16

	// Remove trailer
	ciphertext = ciphertext[:len(ciphertext)-16]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, NewError(CodeDecryptionFailed, "create AES cipher: "+err.Error())
	}

	// AES/ECB decrypt
	plaintext := make([]byte, len(ciphertext))
	block.Decrypt(plaintext, ciphertext)

	// Remove null-byte padding
	if origRemainder != 0 {
		// Data was not a multiple of 16; strip the last partial block's padding
		cutLen := 16 - origRemainder
		plaintext = plaintext[:len(plaintext)-cutLen]
	}
	// If origRemainder == 0, data was already a multiple of 16, no padding to strip

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