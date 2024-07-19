package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func Sha512(bytes []byte) string {
	h := sha512.New()
	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}