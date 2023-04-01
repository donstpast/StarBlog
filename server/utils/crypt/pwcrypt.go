package crypt

import (
	"crypto/rand"
	"encoding/base64"
	"golang.org/x/crypto/scrypt"
	"log"
)

// 生成随机盐值
func GenerateSalt() string {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		log.Fatal(err)
	}
	return base64.StdEncoding.EncodeToString(salt)
}

// 使用随机盐值生成密码哈希值
func PwScrypt(password string) (string, string) {
	const keyLen = 11
	salt := GenerateSalt()
	PwHash, err := scrypt.Key([]byte(password), []byte(salt), 262144, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(PwHash)
	return fpw, salt
}

// 验证密码是否正确
func VerifyPassword(password string, hashedPassword string, salt string) bool {
	PwHash, err := scrypt.Key([]byte(password), []byte(salt), 262144, 8, 1, 11)
	if err != nil {
		log.Fatal(err)
	}
	fpw := base64.StdEncoding.EncodeToString(PwHash)
	return fpw == hashedPassword
}

//// PwScrypt 用户存储，密码加盐
//func PwScrypt(password string) string {
//	const keyLen = 11 //定义常量Keylen的值
//	salt := make([]byte, 8)
//	salt = []byte{19, 83, 128, 30, 117, 110, 214, 212}
//	PwHash, err := scrypt.Key([]byte(password), salt, 262144, 8, 1, keyLen)
//	if err != nil {
//		log.Fatal(err)
//	}
//	fpw := base64.StdEncoding.EncodeToString(PwHash)
//	return fpw
//}
