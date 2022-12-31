package main

import (
	"algorithms/base/core/random"
	"algorithms/base/data_structures/binary_tree"
	"algorithms/base/data_structures/queue"
	"fmt"
)

func main() {
	a := random.GetIntTree(10, 10)
	lca := lowestCommonAncestor(a.Root, a.Root.RightNode, a.Root.LeftNode)
	fmt.Println(lca.Data)
	lcaRCA := LCARec(a.Root, a.Root.RightNode, a.Root.LeftNode)
	fmt.Println(lcaRCA.Data)
}

func lowestCommonAncestor(root, p, q *binary_tree.BinaryNode) *binary_tree.BinaryNode {
	qu := queue.NewQueue(root) // создаем новую очередь
	for qu.GetLen() > 0 {      // интерируемся пока длинна очередь > 0
		node := qu.Pop() //выбиваем первый элемент
		if node == nil { //если нил то возвращаем пустую ноду
			return node
		}

		if p.Data > node.Data && q.Data > node.Data { //если p и q даты больше чем дата текущей ноды, то идем вправо
			qu.Push(p.RightNode)
		} else if p.Data < node.Data && q.Data < node.Data { //если же меньше, то влево
			qu.Push(p.LeftNode)
		} else {
			return node //при любом другом случае возвращаем предка, так как ноды на разных ветвях и это значит что текущая нода - ближайший предок
		}
	}
	return nil
}

func LCARec(root, p, q *binary_tree.BinaryNode) *binary_tree.BinaryNode { //рекурсивно решение
	if root == nil || root == p || root == q { //если мы дошли до нуля
		return root
	}

	l := lowestCommonAncestor(root.LeftNode, p, q)  //левая нода
	r := lowestCommonAncestor(root.RightNode, p, q) //правая нода

	if l != nil && r != nil { //если обе ноды не равны нулю, то мы нашли ноды и слева и справа, что означает что рут дерева - lca
		return root
	}

	if l == nil { //если левая нода нил, то слева нет нужных нод и нужные ноды справа, поэтому возвращаем первую найденную с правой части
		return r
	}
	return l //иначе с левой части
}
