package main

import (
    "github.com/hughes/racka"
    "fmt"
    "net/http"
    "os"
)

func main() {
    fmt.Println("Running server!")
    
    // static files
    http.Handle("/static/", http.StripPrefix("/static",
        http.FileServer(http.Dir("./static"))))

    // everything else
    http.HandleFunc("/user/", racka.UserDetail)
    http.HandleFunc("/", racka.Index)

    http.ListenAndServe(":" + os.Getenv("PORT"), nil)
}
