package futuapi

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"os"
	"testing"
)

func TestRSAEncryptRoundTripWithPrivateKey(t *testing.T) {
	data, err := os.ReadFile("/etc/futu/keys/private_key.pem")
	if err != nil { t.Skip("no key: ", err) }
	plain := []byte(`{"c2s":{"clientVer":300,"clientID":"futuapi4go"}}`)

	cipher, err := RSAEncrypt(string(data), plain)
	if err != nil { t.Fatalf("encrypt: %v", err) }

	block, _ := pem.Decode(data)
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil { t.Fatalf("parse: %v", err) }

	dec, err := rsa.DecryptPKCS1v15(nil, priv, cipher)
	if err != nil { t.Fatalf("decrypt: %v", err) }

	if string(dec) != string(plain) {
		t.Fatalf("round-trip mismatch:\n got: %q\nwant: %q", dec, plain)
	}
	t.Logf("round-trip OK: %d bytes -> %d bytes -> %d bytes", len(plain), len(cipher), len(dec))
}
