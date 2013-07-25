package racka

import (
	"crypto/sha1"
	"encoding/base64"
	"io"
	"time"
)

type User struct {
	pk       int
	email    string
	password string
}

type Day struct {
	date  time.Time
	hours string
}

func SetPassword(u *User, p string) {
	h := sha1.New()
	io.WriteString(h, p)
	password := base64.URLEncoding.EncodeToString(h.Sum(nil))
	u.password = password
}
