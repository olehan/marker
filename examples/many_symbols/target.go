package main

import "github.com/olehan/marker/examples/many_symbols/tree"

var (
    Private = false
    Depth   = 11
    Date    = 11.11
    Version = "v1.10.1"
    Name    = "ebin"
    Author  = "spurdo sparde"
    Repo    = "https://github.com/olehan"
    Tree = tree.NewTree("root dep",
        tree.NewTree("left dep 1",
            tree.NewTree("left left dep 2", nil, nil),
            tree.NewTree("left right dep 2", nil, nil),
        ),
        tree.NewTree("right dep 1",
            tree.NewTree("right left dep 2", nil, nil),
            tree.NewTree("right right dep 2", nil, nil),
        ),
    )
)
