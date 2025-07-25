package myCollections

import (
	"errors"
	"fmt"
)

type LinkedList struct {
	Head *NodeLinkedList
	size uint64
}

type NodeLinkedList struct {
	Val  int64
	Prev *NodeLinkedList
	Next *NodeLinkedList
}

// NewSinglyLinkedList возвращает объект "односвязный список"
func NewLinkedList(val int64) *LinkedList {
	return &LinkedList{Head: &NodeLinkedList{
		Val:  val,
		Prev: nil,
		Next: nil,
	},
		size: 1}
}

// Prepend добавляет элемент в начало списка
func (list *LinkedList) Prepend(val int64) {
	list.Head = &NodeLinkedList{
		Val:  val,
		Next: list.Head,
	}
	list.size++
}

// Print печает в консоль весь двусвязный список(в прямом направлении)
func (list *LinkedList) Print() {
	v := *list.Head
	for {
		fmt.Println(v.Val)
		if v.Next != nil {
			v = *v.Next
		} else {
			break
		}
	}
}

// Append добавляет элемент в конец двусвязного списка
func (list *LinkedList) Append(val int64) {
	v := list.Head
	for {
		if v.Next != nil {
			v = v.Next
		} else {
			break
		}
	}
	v.Next = &NodeLinkedList{
		Val:  val,
		Prev: v,
		Next: nil,
	}
	list.size++
}

// RemoveByValue удаляет элемент по значению
func (list *LinkedList) RemoveByValue(val int64) {
	if list.Head.Val == val {
		list.Head = list.Head.Next
		return
	}
	v := list.Head
	for {
		if v.Next != nil {
			if v.Next.Val == val {
				v.Next = v.Next.Next
				break
			}
			v = v.Next
		} else {
			break
		}
	}
	list.size--
}

// RemoveByIndex удаляет элемент по индексу
func (list *LinkedList) RemoveByIndex(index uint64) {
	if index == 0 {
		list.Head = list.Head.Next
		return
	}
	var ind uint64 = 1
	v := list.Head
	for true {
		if v.Next != nil {
			if ind == index {
				v.Next = v.Next.Next
				list.size--
				break
			}
			v = v.Next
			ind++
		} else {
			break
		}
	}
}

// Find принимает на вход значение элемента списка и возвращает его индекс,
// если элемент не найден, возвращает -1
func (list *LinkedList) Find(val int64) int64 {
	v := list.Head
	for ind := int64(0); v != nil; ind++ {
		if v.Val == val {
			return ind
		}
		v = v.Next
	}
	return -1
}

// Size возвращает количество элементов в односвязном списке
func (list *LinkedList) Size() uint64 {
	return list.size
	//var count uint64
	//v := list.Head
	//for v != nil {
	//	count++
	//	v = v.Next
	//}
	//return count
}

// GetValue возвращает элемент по индексу в двусвязном
func (list *LinkedList) GetValue(index uint64) (int64, error) {
	v := list.Head
	for count := uint64(0); &v != nil; count++ {
		if count == index {
			return v.Val, nil
		}
		v = v.Next
	}
	return 0, errors.New(fmt.Sprintf("Not found element with index %d", index))
}
