// Copyright 2026 shing1211
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package futuapi

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"math/big"
)

// RSAEncrypt encrypts data using RSA public key (PKCS1v15).
// Accepts a "PUBLIC KEY" (PKIX) PEM. For convenience during testing, also
// accepts an "RSA PRIVATE KEY" (PKCS1/PKCS8) PEM and extracts the public key.
//
// WARNING: Passing a private key PEM is NOT recommended in production.
// The private key material is loaded into memory only to extract the public
// key, which is unnecessary and introduces unnecessary risk. Always pass
// the public key PEM from your OpenD deployment.
//
// Data is encrypted in chunks compatible with Futu OpenD's protocol:
// each input chunk is at most (keySize - 11) bytes, producing a
// keySize-byte ciphertext block per chunk. This matches the chunked RSA
// encryption used by the Futu Python/C++ SDKs.
func RSAEncrypt(publicKeyPEM string, data []byte) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block")
	}

	var rsaPub *rsa.PublicKey

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err == nil {
		var ok bool
		rsaPub, ok = pub.(*rsa.PublicKey)
		if !ok {
			return nil, fmt.Errorf("not an RSA public key")
		}
	} else {
		logf("WARNING: RSAEncrypt received a private key PEM instead of a public key PEM. " +
			"This is convenient for testing but NOT recommended in production. " +
			"Use a PUBLIC KEY PEM from your OpenD deployment.")
		priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			privInterface, err := x509.ParsePKCS8PrivateKey(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("failed to parse as public key or private key: %w", err)
			}
			var ok bool
			priv, ok = privInterface.(*rsa.PrivateKey)
			if !ok {
				return nil, fmt.Errorf("not an RSA private key")
			}
			rsaPub = &priv.PublicKey
		} else {
			rsaPub = &priv.PublicKey
		}
	}

	keySize := (rsaPub.N.BitLen() + 7) / 8 // key size in bytes
	encChunkSize := keySize - 11
	if encChunkSize <= 0 {
		return nil, fmt.Errorf("RSA key size too small: %d bits", rsaPub.N.BitLen())
	}
	cipherChunkSize := keySize

	// Pre-allocate output buffer
	numBlocks := (len(data) + encChunkSize - 1) / encChunkSize
	if numBlocks == 0 {
		numBlocks = 1
	}
	result := make([]byte, numBlocks*cipherChunkSize)

	for i, outOff := 0, 0; i < len(data); i += encChunkSize {
		end := i + encChunkSize
		if end > len(data) {
			end = len(data)
		}
		chunk := data[i:end]

		// Build EM = 0x00 || 0x02 || PS || 0x00 || M
		em := make([]byte, cipherChunkSize)
		em[0] = 0x00
		em[1] = 0x02

		// PS: random non-zero bytes
		psLen := cipherChunkSize - len(chunk) - 3
		ps := em[2 : 2+psLen]
		if err := nonZeroRandomBytes(ps); err != nil {
			return nil, fmt.Errorf("random padding: %w", err)
		}

		em[2+psLen] = 0x00
		copy(em[2+psLen+1:], chunk)

		// Raw RSA encryption: c = m^e mod n
		e := big.NewInt(int64(rsaPub.E))
		m := new(big.Int).SetBytes(em)
		c := new(big.Int).Exp(m, e, rsaPub.N)
		cipher := c.Bytes()

		// Left-pad with zeros to keySize
		offset := cipherChunkSize - len(cipher)
		copy(result[outOff+offset:], cipher)

		outOff += cipherChunkSize
	}

	return result, nil
}

// nonZeroRandomBytes fills dst with random non-zero bytes.
// This matches the padding generation used by PKCS1v15.
func nonZeroRandomBytes(dst []byte) error {
	n := len(dst)
	buf := make([]byte, min(n, 64))
	for i := 0; i < n; {
		if _, err := io.ReadFull(rand.Reader, buf); err != nil {
			return err
		}
		for j := 0; j < len(buf) && i < n; j++ {
			if buf[j] != 0 {
				dst[i] = buf[j]
				i++
			}
		}
	}
	return nil
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
