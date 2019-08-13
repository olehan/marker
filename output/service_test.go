package output

import "testing"

func TestRecoverBadTypeCast(t *testing.T) {
    called := false

    defer func() {
        if !called {
            t.FailNow()
        }
    }()

    defer RecoverBadTypeCast(func(err error) {
        called = true
    })

    var i Output = 0
    t.Log(i.(string))
}
