package main

import (
	"fmt"
	"time"
	"io"
	"crypto/sha1"
)

type User struct {
	pk       int
	email    string
	password string
}

type Day struct {
	date time.Time
	hours string
}

func generatePass(u User, p String) {
	h := sha1.New()
	io.WriteString(h, p)
	u.password = h.Sum()
}

func main() {
	fmt.Println("Running server...")
}
