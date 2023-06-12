package data

import (
	"crypto/md5"
	"encoding/hex"
)

// MD5V : md5签名 that what v means
func MD5V(str []byte, b ...byte) string {
	//java快速转go-声明变量并赋值
	md5Instance := md5.New()
	md5Instance.Write(str)
	return hex.EncodeToString(md5Instance.Sum(b))
}
