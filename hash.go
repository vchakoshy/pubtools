package pubtools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// HashLib struct
type HashLib struct{}

// Hash returns hash
func Hash() *HashLib {
	return &HashLib{}
}

// MD5 returns md5 hash of string
func (h *HashLib) MD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// Base64 returns base64 hash of string
func (h *HashLib) Base64(text string) string {
	data := []byte(text)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}
