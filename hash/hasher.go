package hash

import (
	"crypto/md5"
	"encoding/hex"
)

type md5Hash struct {
	password string
	salt string
}

func NewMd5Hash(password, salt string) *md5Hash {
	return &md5Hash{
		password: password,
		salt: salt,
	}
}

func (h *md5Hash) GetSalt() string {
	return h.salt
}

func (h *md5Hash) Hash() string {
	hasher := md5.New()
	hasher.Write([]byte(h.password + h.salt))
	return hex.EncodeToString(hasher.Sum(nil))
}