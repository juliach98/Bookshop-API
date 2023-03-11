package helpers

import (
	"crypto/md5"
	"encoding/hex"
)

func GetEncryptedPassword(password string) string {
	hash := md5.Sum([]byte(password))
	pass := hex.EncodeToString(hash[:])

	return pass
}
