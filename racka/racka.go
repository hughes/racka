package main

import (
	"github.com/hughes/racka"
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Running server!")
	http.HandleFunc("/", racka.Index)
	http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
