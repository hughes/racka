package racka

import (
    "testing"
)

func TestGeneratePass(t *testing.T) {
    u := User{}
    SetPassword(&u, "asdf")
    if u.password != "PaVBVZkYqAjCQCu6UBL2xgsnZhw=" {
        t.Error("Failed to hash password: " + u.password)
    }
}
