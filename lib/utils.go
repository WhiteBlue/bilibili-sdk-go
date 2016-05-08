package lib

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)


//simple Md5 encrypt
func Md5(formal string) string {
	h := md5.New()
	h.Write([]byte(formal))
	return hex.EncodeToString(h.Sum(nil))
}


//数字判断
func IsNumber(target string) bool {
	_, err := strconv.Atoi(target);
	if err != nil {
		return false
	}
	return true
}

