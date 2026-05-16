package util

import (
	"crypto/md5"
	"fmt"
)

// MD5Hex computes the MD5 hash of the input string and returns it as a
// lowercase hex string.
//
// Example:
//
//	MD5Hex("hello") -> "5d41402abc4b2a76b9719d911017c592"
func MD5Hex(input string) string {
	h := md5.Sum([]byte(input))
	return fmt.Sprintf("%x", h)
}

// EncryptPassword hashes a password using MD5, producing the format required
// by the Futu OpenAPI login flow.
//
// This is a convenience wrapper around MD5Hex for the common case of
// hashing a trading password before passing it to UnlockTrade.
//
// Example:
//
//	hashed := EncryptPassword("myPassword123")
func EncryptPassword(password string) string {
	return MD5Hex(password)
}
