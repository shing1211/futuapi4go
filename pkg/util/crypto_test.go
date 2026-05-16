package util

import "testing"

func TestMD5Hex(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"hello", "5d41402abc4b2a76b9719d911017c592"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
		{"password123", "482c811da5d5b4bc6d497ffa98491e38"},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got := MD5Hex(tt.input); got != tt.want {
				t.Errorf("MD5Hex(%q) = %s, want %s", tt.input, got, tt.want)
			}
		})
	}
}

func TestEncryptPassword(t *testing.T) {
	// EncryptPassword is a convenience wrapper around MD5Hex
	got := EncryptPassword("hello")
	want := MD5Hex("hello")
	if got != want {
		t.Errorf("EncryptPassword() = %s, want %s", got, want)
	}
}
