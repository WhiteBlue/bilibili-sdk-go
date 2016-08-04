package utils

import (
	"crypto/md5"
	"encoding/hex"
)

func Md5(formal string) string {
	h := md5.New()
	h.Write([]byte(formal))
	return hex.EncodeToString(h.Sum(nil))
}
