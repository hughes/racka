package racka

import (
    "fmt"
    "net/http"
    "strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, world")
}

func UserDetail(w http.ResponseWriter, r *http.Request) {
    path := "/user/"
    lenPath := len(path)
    id_fragment := r.URL.Path[lenPath:]
    id, err := strconv.ParseInt(id_fragment, 10, 64)
    if err != nil {
        http.Error(
            w,
            fmt.Sprintf("Invalid user ID: %s", id_fragment),
            http.StatusBadRequest)
        return
    }
    response := fmt.Sprintf("%d", id)
    fmt.Fprintf(w, response)
}
