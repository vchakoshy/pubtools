package pubtools

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
)

// GetMD5Hash returns md5 hash of string
func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

// GetBase64Hash returns base64 hash of string
func GetBase64Hash(text string) string {
	data := []byte(text)
	str := base64.StdEncoding.EncodeToString(data)
	return str
}
