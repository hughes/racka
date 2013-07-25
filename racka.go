package main

import (
	"fmt"
	"time"
)

type User struct {
	pk       int
	key      string
	email    string
	password string
}

type Day struct {
	date time.Time
	hours string
}

func main() {
	fmt.Println("Running server...")
}
