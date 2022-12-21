package main

import (
    "algorithms/base/binary_tree"
    "algorithms/base/random"
    "fmt"
)

var (
    start = 1 << 32
)

func main() {
    fmt.Println(isValidBST(random.GetRandIntTree(2, 10).Root))
}

func isValidBST(root *binary_tree.BinaryNode) bool {
    return isValidBSTRecursion(root, start, start)
}

func isValidBSTRecursion(root *binary_tree.BinaryNode, min, max int) bool {
    if root == nil {
        return true
    } else if min != start && root.Data <= min || max != start && root.Data >= max {
        return false
    }
    return isValidBSTRecursion(root.LeftNode, min, root.Data) && isValidBSTRecursion(root.RightNode, root.Data, max)
}
