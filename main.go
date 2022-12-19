package main

import (
    "algorithms/array"
)

var custom struct {
    error
}

func main() {
    a := array.MakeArray()
    a.Append(1, 2, 3, 4, 5, 6, 7)
}

func test() {}
