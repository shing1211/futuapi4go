package futuapi

import (
	"bytes"
	"crypto/rand"
	"testing"
)

func aesTestKey(t testing.TB) []byte {
	t.Helper()
	key := make([]byte, 16)
	if _, err := rand.Read(key); err != nil {
		t.Fatal(err)
	}
	return key
}

func TestFTAESEncryptDecrypt_RoundTrip(t *testing.T) {
	tests := []struct {
		name    string
		payload []byte
	}{
		{"empty", []byte{}},
		{"single byte", []byte{0x42}},
		{"16 bytes (exact block)", bytes.Repeat([]byte{0xFF}, 16)},
		{"17 bytes (one over)", bytes.Repeat([]byte{0xAB}, 17)},
		{"32 bytes (two blocks)", bytes.Repeat([]byte{0x01, 0x02}, 16)},
		{"1K random", func() []byte { b := make([]byte, 1024); rand.Read(b); return b }()},
		{"10K random", func() []byte { b := make([]byte, 10240); rand.Read(b); return b }()},
	}

	key := aesTestKey(t)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ciphertext, err := ftaesEncrypt(key, tt.payload)
			if err != nil {
				t.Fatalf("encrypt: %v", err)
			}

			plaintext, err := ftaesDecrypt(key, ciphertext)
			if err != nil {
				t.Fatalf("decrypt: %v", err)
			}

			if !bytes.Equal(plaintext, tt.payload) {
				t.Errorf("plaintext mismatch: got %x, want %x", plaintext, tt.payload)
			}
		})
	}
}

func TestFTAESEncrypt_Trailer(t *testing.T) {
	key := aesTestKey(t)
	payload := []byte("hello world")

	ciphertext, err := ftaesEncrypt(key, payload)
	if err != nil {
		t.Fatal(err)
	}

	// Ciphertext must have at least 16-byte trailer
	if len(ciphertext) < 16 {
		t.Fatal("ciphertext too short, missing trailer")
	}

	// Last byte of trailer = padLen
	trailer := ciphertext[len(ciphertext)-16:]
	padLen := int(trailer[15])
	expectedPadLen := 16 - (len(payload) % 16)
	if expectedPadLen == 16 {
		expectedPadLen = 0
	}
	if padLen != expectedPadLen {
		t.Errorf("padLen mismatch: got %d, want %d", padLen, expectedPadLen)
	}
}

func TestFTAESDecrypt_TooShort(t *testing.T) {
	key := aesTestKey(t)
	_, err := ftaesDecrypt(key, []byte{0x00, 0x01, 0x02})
	if err == nil {
		t.Error("expected error for ciphertext < 16 bytes")
	}
}

func TestFTAESEncrypt_InvalidKey(t *testing.T) {
	_, err := ftaesEncrypt([]byte{0x00, 0x01, 0x02}, []byte("data"))
	if err == nil {
		t.Error("expected error for key < 16 bytes")
	}
}

func TestFTAESDecrypt_InvalidKey(t *testing.T) {
	_, err := ftaesDecrypt([]byte{0x00, 0x01, 0x02}, []byte("data"))
	if err == nil {
		t.Error("expected error for key < 16 bytes")
	}
}

func TestAESPublicWrappers(t *testing.T) {
	key := aesTestKey(t)
	payload := []byte("test payload for public wrappers")

	enc, err := AESEncrypt(key, payload)
	if err != nil {
		t.Fatalf("AESEncrypt: %v", err)
	}

	dec, err := AESDecrypt(key, enc)
	if err != nil {
		t.Fatalf("AESDecrypt: %v", err)
	}

	if !bytes.Equal(dec, payload) {
		t.Errorf("round-trip failed: got %x, want %x", dec, payload)
	}
}

func TestFTAESMulitblockPadding(t *testing.T) {
	key := aesTestKey(t)
	// Exact multiples of block size need padLen=16 (special: no padding)
	sizes := []int{0, 1, 15, 16, 31, 32, 48, 64, 100, 256}
	for _, size := range sizes {
		t.Run("", func(t *testing.T) {
			payload := make([]byte, size)
			rand.Read(payload)

			enc, err := ftaesEncrypt(key, payload)
			if err != nil {
				t.Fatal(err)
			}

			dec, err := ftaesDecrypt(key, enc)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(dec, payload) {
				t.Errorf("size %d: round-trip mismatch", size)
			}
		})
	}
}

func BenchmarkFTAESEncrypt(b *testing.B) {
	key := aesTestKey(b)
	payload := make([]byte, 1024)
	rand.Read(payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ftaesEncrypt(key, payload)
	}
}

func BenchmarkFTAESDecrypt(b *testing.B) {
	key := aesTestKey(b)
	payload := make([]byte, 1024)
	rand.Read(payload)
	enc, _ := ftaesEncrypt(key, payload)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ftaesDecrypt(key, enc)
	}
}
