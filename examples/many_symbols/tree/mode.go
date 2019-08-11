package tree

import "fmt"

type (
    Tree struct {
        Left  *Tree
        Right *Tree
        Value string
    }
)

func NewTree(value string, left *Tree, right *Tree) *Tree {
    return &Tree{left, right, value}
}

func PrintTree(tree *Tree, offset int) {
    if tree == nil {
        return
    }

    for i := 0; i < offset; i++ {
        fmt.Print(" ")
    }

    fmt.Println("└──", tree.Value)

    nextOffset := offset + 2

    PrintTree(tree.Left, nextOffset)
    PrintTree(tree.Right, nextOffset)
}
