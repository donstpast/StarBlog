package crypt

import (
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

// PwScrypt 用户存储，密码加盐
func PwScrypt(password string) string {
	const keyLen = 11 //定义常量Keylen的值
	salt := make([]byte, 8)
	salt = []byte{19, 83, 128, 30, 117, 110, 214, 212}
	PwHash, err := scrypt.Key([]byte(password), salt, 262144, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(PwHash)
	return fpw
}
