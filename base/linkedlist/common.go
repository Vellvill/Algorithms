package linkedlist

type LinkedList struct {
    Value int
    Next  *LinkedList
}

func CreateList(v int) *LinkedList {
    return &LinkedList{Value: v, Next: nil}
}

func (l *LinkedList) AddNode(v int) {
    for n := l; n != nil; n = n.Next {
        if n.Next == nil {
            new := CreateList(v)
            n.Next = new
            break
        }
    }
}

func (l *LinkedList) MakeCycle() *LinkedList {
    for n := l; n != nil; n = n.Next {
        if n.Next == nil {
            n.Next = l.FindHead()
            break
        }
    }
    return l
}

func (l *LinkedList) ToSlice() []int {
    out := make([]int, 0)
    for n := l; n != nil; n = n.Next {
        out = append(out, n.Value)
    }
    return out
}

func (l *LinkedList) FindHead() *LinkedList {
    var Head *LinkedList
    for n := l; n != nil; n = n.Next {
        Head = n
    }
    return Head
}

func (l *LinkedList) Delete(v int) {
    for n := l; n != nil; n = n.Next {
        if n.Value == v {
            n.Value = n.Next.Value
            n.Next = n.Next.Next
        }
    }
}

func (l *LinkedList) Reverse() *LinkedList {

    cur := l
    var prev, next *LinkedList

    for cur != nil {
        //копируем указатель на сл элемент
        next = cur.Next

        //разворачиваем следующий указатель
        cur.Next = prev

        //идём дальше по листу
        prev = cur
        cur = next
    }

    return prev

}

func (l *LinkedList) AddNodeByValue(WhichToAdd, Adder int) {
    for n := l; n != nil; n = n.Next {
        if n.Value == WhichToAdd {
            new := CreateList(Adder)
            new.Next = n.Next
            n.Next = new
            break
        }
    }
}

func (l *LinkedList) AddNodeToBegin(v int) {
    new := CreateList(v)
    new.Next = l
}
