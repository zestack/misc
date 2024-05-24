package misc

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"strings"
)

func MD5(str string) string {
	hash := md5.Sum([]byte(str))
	//md5str := hex.EncodeToString(hash)
	md5str := fmt.Sprintf("%x", hash) //将[]byte转成16进制
	return strings.ToUpper(md5str)
}

func Sha1(str string) string {
	hash := sha1.Sum([]byte(str))
	//return hex.EncodeToString(hash)
	return fmt.Sprintf("%x", hash)
}

func Sha256(str string) string {
	hash := sha256.Sum256([]byte(str))
	//return hex.EncodeToString(hash)
	return fmt.Sprintf("%x", hash)
}

func Hmac(key, data string) string {
	h := hmac.New(md5.New, []byte(key))
	h.Write([]byte(data))
	//return hex.EncodeToString(h.Sum([]byte("")))
	return fmt.Sprintf("%x", h.Sum([]byte("")))
}
