package main

import (
    "fmt"
    "github.com/olehan/marker"
    "github.com/olehan/marker/examples/basic/config"
    "github.com/olehan/marker/path"
)

func main() {
    m, err := marker.Mark(
        path.RelativeToCaller("./target.go"),
        marker.NewMarker("Config", config.Config{}),
    )

    if err != nil {
        panic(err)
    }

    // Symbol's always pointer.
    cfg := m.(*config.Config)

    fmt.Println("name:", cfg.Name)
    fmt.Println("version:", cfg.Version)
}
