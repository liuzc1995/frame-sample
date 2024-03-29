package model

import (
	"crypto/md5"
	"encoding/hex"
)

//md5校验
func GeneratePasswordHash(pwd string) string {
	hasher := md5.New()
	hasher.Write([]byte(pwd))
	pwdHash := hex.EncodeToString(hasher.Sum(nil))
	return pwdHash
}
