package main

import (
    "fmt"
    "github.com/olehan/marker"
    "github.com/olehan/marker/output"
    "github.com/olehan/marker/path"
)

func main() {
    // RecoverBadTypeCast recovers only bad interface type casts
    // and executes the given callback in that scenario.
    defer output.RecoverBadTypeCast(func(err error) {
        fmt.Println("number type cast error:", err)
    })

    m, err := marker.Mark(
        path.RelativeToCaller("./target.go"),
        marker.NewValueMarker("Number", marker.Int16Value),
    )

    if err != nil {
        panic(err)
    }

    // This line will cause panic and RecoverBadTypeCast
    // will execute catch callback.
    _ = m.(*int) // change int to int16 to get success run.
}
