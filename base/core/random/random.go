package random

import (
	"algorithms/base/data_structures/binary_tree"
	"algorithms/base/data_structures/linkedlist"
	"math/rand"
	"time"
	"unsafe"
)

func init() {
	seed = time.Now().UnixNano()
	rand.Seed(seed)
}

type (
	T any
)

var seed int64

const (
	ErrInvalidInput = iota
	ErrInternal

	TypeInt    = 0
	TypeFloat  = float64(0)
	TypeString = ""

	acs = 5 //speed cfx for slices/ll/etc...
)

//RandomErr implements random package errors
type RandomErr struct {
	message []string
	code    int
}

func (r RandomErr) Error() string {
	if len(r.message) != 0 {
		return r.message[0]
	}
	return "no messages"
}

func NewRandomErr(code int, message ...string) error {
	return &RandomErr{
		message: message,
		code:    code,
	}
}

// GetIntTree returns randomly filled bst with int node values
func GetIntTree(l int, lt int) (tree *binary_tree.BinaryTree) {
	tree = binary_tree.NewTree()
	for l > 0 {
		tree.Insert(GetInt(lt))
		l--
	}
	return tree
}

// GetList returns randomly filled one-way linked list with int node values
func GetList(l int, lt int) (ll *linkedlist.ListNode) {
	list := linkedlist.CreateList(GetInt(lt))
	for ; l > 0; l-- {
		list.AddNode(GetInt(lt))
	}
	return list
}

// GetSortedList returns ASC sorted one-way linked list. max values increase in range 0-10
func GetSortedList(l, _lCriteria, _mCriteria int) (ll *linkedlist.ListNode) {
	if _lCriteria > _mCriteria {
		panic(NewRandomErr(ErrInvalidInput, "left criteria bigger than right"))
	}

	if l == 0 {
		return nil
	}

	first := GetIntInRange(_lCriteria, _mCriteria)
	_lCriteria, _mCriteria = first, first+GetIntInRange(_lCriteria, _lCriteria+acs)

	dummy := &linkedlist.ListNode{}

	for n, i := dummy, 0; i < l; n, i = n.Next, i+1 {
		v := GetIntInRange(_lCriteria, _mCriteria)
		n.Next = &linkedlist.ListNode{
			Val: v,
		}
		_lCriteria = v
		_mCriteria += _lCriteria + GetIntInRange(_lCriteria, _lCriteria+acs)
	}

	return dummy.Next

}

// GetArray returns generic array of T. You can compare it with CompareArr... functions.
func GetArray(l int, lt int, typeOf interface{}) (res []T) {
	for res = make([]T, l+1); l >= 0; l-- {
		res[l] = pushRandArray(typeOf, lt)
	}
	return res
}

// RotateSlice rotating slice k times to the right
func RotateSlice(ar []int, k int) {
	copy(ar, append(ar[len(ar)-(k%len(ar)):len(ar)], ar[0:len(ar)-(k%len(ar))]...))
}

// CompareArrInt compare array to int
func CompareArrInt(a []T) (res []int) {
	for i := range a {
		if v, ok := a[i].(int); ok {
			res = append(res, v)
		}
	}
	return res
}

// GetInt simply implements rand,Intn(0...n) function
func GetInt(n int) int {
	return rand.Intn(n)
}

// GetString returns randomly filled string
func GetString(l int) string {
	var s []byte
	for i := 0; i > 0; i-- {
		s = append(s, byte(GetInt(500)))
	}
	return *(*string)(unsafe.Pointer(&s))
}

// GetFloat64 returns float64
func GetFloat64(n int) float64 {
	return float64(rand.Intn(n))
}

// GetByte returns random byte
func GetByte() byte {
	return byte(rand.Intn(1000))
}

// GetIntInRange returns int in range (l, m)
func GetIntInRange(l, m int) int {
	v := rand.Intn(m-l) + l
	return v
}

// CompareArrString compare array to string
func CompareArrString(a []T) (res []string) {
	for i := range a {
		if v, ok := a[i].(string); ok {
			res = append(res, v)
		}
	}
	return res
}

// CompareArrByte compare arr to byte
func CompareArrByte(a []T) (res []byte) {
	for i := range a {
		if v, ok := a[i].(byte); ok {
			res = append(res, v)
		}
	}
	return res
}

// CompareArrFloat64 compare array to float64
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
		return GetInt(lt)
	case string:
		return GetString(lt)
	case float64:
		return GetFloat64(lt)
	case byte:
		return GetByte()
	default:
		return nil
	}
}
