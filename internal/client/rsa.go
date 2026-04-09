package futuapi

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

// RSAEncrypt encrypts data using RSA public key (PKCS1v15)
func RSAEncrypt(publicKeyPEM string, data []byte) ([]byte, error) {
	// Parse PEM block
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	// Parse public key
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %w", err)
	}

	rsaPub, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("not an RSA public key")
	}

	// Encrypt data
	encrypted, err := rsa.EncryptPKCS1v15(rand.Reader, rsaPub, data)
	if err != nil {
		return nil, fmt.Errorf("encryption failed: %w", err)
	}

	return encrypted, nil
}

// GenerateRSAKeys generates a new RSA key pair for testing
func GenerateRSAKeys(bits int) (privateKeyPEM, publicKeyPEM string, err error) {
	// Generate key
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return "", "", fmt.Errorf("failed to generate key: %w", err)
	}

	// Marshal private key
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyPEM = string(pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}))

	// Marshal public key
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		return "", "", fmt.Errorf("failed to marshal public key: %w", err)
	}
	publicKeyPEM = string(pem.EncodeToMemory(&pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}))

	return privateKeyPEM, publicKeyPEM, nil
}

