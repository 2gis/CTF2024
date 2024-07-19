package utils

import (
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(bytes []byte) string {
	h := sha512.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}

func Hash(b []byte) []byte {
	hasher := sha256.New()
	hasher.Write([]byte(b))
	return hasher.Sum(nil)
}