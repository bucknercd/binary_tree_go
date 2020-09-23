package main
import (
    "fmt"
)

type Node struct {
    Value int
    Left *Node
    Right *Node
}

func (n *Node) Add(val int) {
    if n != nil {
        if val < n.Value {
            if n.Left == nil {
                n.Left = &Node{Value: val}
            } else {
                n.Left.Add(val)
            }
        } else if val > n.Value {
            if n.Right == nil {
                n.Right = &Node{Value: val}
            } else {
                n.Right.Add(val)
            }
        }
    } else {
        n = &Node{Value: val}
    }
}

func Preorder(n *Node) {
    if n != nil {
        fmt.Println(n.Value)
        Preorder(n.Left)
        Preorder(n.Right)
    }
}

func InOrder(n *Node) {
    if n != nil {
        InOrder(n.Left)
        fmt.Println(n.Value)
        InOrder(n.Right)
    }
}

func main() {
    var tree Node
    vals := []int{1,44,22,4,7,99}
    for _, val := range vals {
        tree.Add(val)
    }
    fmt.Println(tree)
    fmt.Println("PREORDER:")
    Preorder(&tree)
    fmt.Println("INORDER:")
    InOrder(&tree)
}
