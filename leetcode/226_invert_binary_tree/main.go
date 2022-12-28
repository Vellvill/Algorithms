package _26_invert_binary_tree

import (
    "algorithms/base/binary_tree"
    "algorithms/base/random"
)

func main() {
    invertTree(random.GetRandIntTree(random.GetRandomInt(10), random.GetRandomInt(10)).Root)
}

func invertTree(root *binary_tree.BinaryNode) *binary_tree.BinaryNode {
    if root == nil {
        return nil
    }

    l := invertTree(root.LeftNode)
    r := invertTree(root.RightNode)

    root.LeftNode = r
    root.RightNode = l

    return root
}