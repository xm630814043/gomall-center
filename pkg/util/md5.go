package util

import "crypto/md5"

import "encoding/hex"

// EncodeMD5 加密
func EncodeMD5(str string) string {
	m := md5.New()
	m.Write([]byte(str))
	return hex.EncodeToString(m.Sum(nil))
}
