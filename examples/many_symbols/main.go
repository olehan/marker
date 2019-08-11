package main

import (
    "fmt"
    "github.com/olehan/marker"
    "github.com/olehan/marker/examples/many_symbols/tree"
    "github.com/olehan/marker/path"
    "reflect"
)

func main() {
    outputMap, err := marker.MarkMany(
        path.RelativeToCaller("./target.go"),
        marker.NewValueMarker("Version", marker.StringValue),
        marker.NewValueMarker("Depth", marker.IntValue),
        marker.NewValueMarker("Date", marker.Float64Value),
        marker.NewValueMarker("Author", marker.StringValue),
        marker.NewValueMarker("Repo", marker.StringValue),
        marker.NewValueMarker("Name", marker.StringValue),
        marker.NewValueMarker("Private", marker.BoolValue),
        marker.NewMarker("Tree", tree.NewTree("", nil, nil)),
    )

    if err != nil {
        panic(err)
    }

    for key, value := range outputMap {
        fmt.Printf("> key: %s | value: %v\n", key, reflect.ValueOf(value).Elem())
    }

    // NewTree returns pointer, the symbol is always pointer to a value,
    // so we have this pointer that points to a pointer situation.
    t := outputMap["Tree"].(**tree.Tree)
    tree.PrintTree(*t, 0)
}
