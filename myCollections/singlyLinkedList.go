package myCollections

import (
	"errors"
	"fmt"
)

func SaYHallo() {
	fmt.Println("Hello, from myCollections!!!!!")
}

type SinglyLinkedList struct {
	Head *NodeSinglyLinkedList
}

type NodeSinglyLinkedList struct {
	Val  int64
	Next *NodeSinglyLinkedList
}

// NewSinglyLinkedList возвращает объект "односвязный список"
func NewSinglyLinkedList(val int64) *SinglyLinkedList {
	return &SinglyLinkedList{Head: &NodeSinglyLinkedList{
		Val:  val,
		Next: nil,
	}}
}

// Prepend добавляет элемент в начало списка
func (list *SinglyLinkedList) Prepend(val int64) {
	list.Head = &NodeSinglyLinkedList{
		Val:  val,
		Next: list.Head,
	}
}

// Print печает в консоль весь односвязный список
func (list *SinglyLinkedList) Print() {
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

// Append добавляет элемент в конец односвязного списка
func (list *SinglyLinkedList) Append(val int64) {
	v := list.Head
	for {
		if v.Next != nil {
			v = v.Next
		} else {
			break
		}
	}
	v.Next = &NodeSinglyLinkedList{
		Val:  val,
		Next: nil,
	}
}

// RemoveByValue удаляет элемент по значению
func (list *SinglyLinkedList) RemoveByValue(val int64) {
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
}

// RemoveByIndex удаляет элемент по индексу
func (list *SinglyLinkedList) RemoveByIndex(index uint64) {
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
func (list *SinglyLinkedList) Find(val int64) int64 {
	v := list.Head
	var ind int64 = 0
	for v != nil {
		if v.Val == val {
			return ind
		} else {
			ind++
			v = v.Next
		}
	}
	return -1
}

// Size возвращает количество элементов в односвязном списке
func (list *SinglyLinkedList) Size() uint64 {
	var count uint64
	v := *list.Head
	for {
		count++
		if v.Next != nil {
			v = *v.Next
		} else {
			break
		}
	}
	return count
}

// GetValue возвращает количество элементов в односвязном списке
func (list *SinglyLinkedList) GetValue(index uint64) (int64, error) {
	var count uint64
	v := *list.Head
	for &v != nil {
		if count == index {
			return v.Val, nil
		} else if v.Next != nil {
			v = *v.Next
			count++
		} else {
			return 0, errors.New(fmt.Sprintf("Not found element with index %d", index))
		}
	}
	return 0, errors.New(fmt.Sprintf("Not found element with index %d", index))
}
