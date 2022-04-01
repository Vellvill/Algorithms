package linkedlist

import (
	"algorithms/common"
	"testing"
)

func TestLinkedList(t *testing.T) {
	t.Run("to_slice", func(t *testing.T) {
		l14 := CreateList(1)
		l14.AddNode(2)
		l14.AddNode(3)
		l14.AddNode(4)

		l11 := CreateList(1)

		l12 := CreateList(1)
		l12.AddNode(1)

		l13 := CreateList(1)
		l13.AddNodeToBegin(2)

		common.Equal(t, l11.ToSlice(), []int{1})
		common.Equal(t, l12.ToSlice(), []int{1, 1})
		common.Equal(t, l13.ToSlice(), []int{1})
		common.Equal(t, l14.ToSlice(), []int{1, 2, 3, 4})
	})
	t.Run("find_head", func(t *testing.T) {
		l11 := CreateList(1)
		l11.AddNode(2)
		l11.AddNode(3)

		common.Equal(t, l11.FindHead(), &LinkedList{3, nil})
	})

	t.Run("add_to_begin", func(t *testing.T) {
		l11 := CreateList(1)
		l11.AddNode(2)
		l11.AddNode(3)
		l11.AddNodeToBegin(4)

		common.Equal(t, l11.ToSlice(), []int{})
	})
}
