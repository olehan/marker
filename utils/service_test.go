package utils

import "testing"

func TestRandomStringRunes(t *testing.T) {
    v1 := RandomStringRunes(10)
    v2 := RandomStringRunes(10)
    if v1 == v2 {
        t.FailNow()
    }

    expectedLen := 5
    v1 = RandomStringRunes(expectedLen)
    if len(v1) != expectedLen {
        t.FailNow()
    }

    expectedLen = 100
    v1 = RandomStringRunes(expectedLen)
    if len(v1) != expectedLen {
        t.FailNow()
    }
}
