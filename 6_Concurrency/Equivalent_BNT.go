package main
import (
    "fmt"
    "golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
    walk(t, ch)
    close(ch)
}

func walk(t *tree.Tree, ch chan int) {
    if t == nil {
        return
    }
    walk(t.Left, ch)
    ch <- t.Value
    walk(t.Right, ch)
}

func Same(t1, t2 *tree.Tree) bool {
    ch1 := make(chan int, 10)
    ch2 := make(chan int, 10)
    go Walk(t1, ch1)
    go Walk(t2, ch2)
    for v1 := range ch1 {
        v2 := <-ch2
        if v1 != v2 {
            return false
        }
    }
    return true
}

func main() {
    tree1 := tree.New(10)
    tree2 := tree.New(10)
    fmt.Println(Same(tree1, tree2))
}