package random

import (
    "time"
    "math/rand"
    "unsafe"
    "algorithms/base/binary_tree"
)

var seed int64

func init() {
    seed = time.Now().UnixNano()
    rand.Seed(seed)
}

type T any

func GetRandIntTree(l int, lt int) (tree *binary_tree.BinaryTree) {
    tree = binary_tree.NewTree()
    for l > 0 {
        tree.Insert(GetRandomInt(lt))
        l--
    }
    return tree
}

func GetRandArray(l int, lt int, typeOf interface{}) (res []T) {
    for res = make([]T, l); l > 0; res = append(res, pushRandArray(typeOf, lt)) {
        l--
    }
    return res
}

func CompareArrInt(a []T) (res []int) {
    for i := range a {
        if v, ok := a[i].(int); ok {
            res = append(res, v)
        }
    }
    return res
}

func CompareArrString(a []T) (res []string) {
    for i := range a {
        if v, ok := a[i].(string); ok {
            res = append(res, v)
        }
    }
    return res
}

func CompareArrByte(a []T) (res []byte) {
    for i := range a {
        if v, ok := a[i].(byte); ok {
            res = append(res, v)
        }
    }
    return res
}

func CompareArrFloat64(a []T) (res []float64) {
    for i := range a {
        if v, ok := a[i].(float64); ok {
            res = append(res, v)
        }
    }
    return res
}

func pushRandArray(typeOf interface{}, lt int) T {
    switch typeOf.(type) {
    case int:
        return GetRandomInt(lt)
    case string:
        return GetRandomString()
    case float64:
        return GetRandomFloat64(lt)
    case byte:
        return GetRandomByte()
    default:
        return nil
    }
}

func GetRandomInt(n int) int {
    return rand.Intn(n)
}

func GetRandomString() T {
    var s []byte
    for i := rand.Intn(5) + 1; i > 0; i-- {
        s = append(s, byte(i))
    }
    return *(*string)(unsafe.Pointer(&s))
}

func GetRandomFloat64(n int) float64 {
    return float64(rand.Intn(n))
}

func GetRandomByte() byte {
    return byte(rand.Intn(1000))
}
