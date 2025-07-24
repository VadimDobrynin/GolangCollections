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

func NewSinglyLinkedList(val int64) *SinglyLinkedList {
	return &SinglyLinkedList{Head: &NodeSinglyLinkedList{
		Val:  val,
		Next: nil,
	}}
}

// Prepend добавляет элемент в начало списка
func (head *SinglyLinkedList) Prepend(val int64) {
	head.Head = &NodeSinglyLinkedList{
		Val:  val,
		Next: head.Head,
	}
}

// Print печает в консоль весь односвязный список
func (head *SinglyLinkedList) Print() {
	v := *head.Head
	for true {
		fmt.Println(v.Val)
		if v.Next != nil {
			v = *v.Next
		} else {
			break
		}
	}
}

// Append добавляет элемент в конец односвязного списка
func (head *SinglyLinkedList) Append(val int64) {
	l := NodeSinglyLinkedList{
		Val:  val,
		Next: nil,
	}
	v := head.Head
	for true {
		if v.Next != nil {
			v = v.Next
		} else {
			break
		}
	}
	v.Next = &l
}

// RemoveByValue удаляет элемент по значению
func (head *SinglyLinkedList) RemoveByValue(val int64) {
	if head.Head.Val == val {
		head.Head = head.Head.Next
		return
	}
	v := head.Head
	for true {
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
func (head *SinglyLinkedList) RemoveByIndex(index uint64) {
	if index == 0 {
		head.Head = head.Head.Next
		return
	}
	var ind uint64 = 1
	v := head.Head
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
func (head *SinglyLinkedList) Find(val int64) int64 {
	v := head.Head
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
func (head *SinglyLinkedList) Size() uint64 {
	var count uint64
	v := *head.Head
	for true {
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
func (head *SinglyLinkedList) GetValue(index uint64) (int64, error) {
	var count uint64
	v := *head.Head
	for &v != nil {
		if count == index {
			return v.Val, nil
		} else if v.Next != nil {
			v = *v.Next
			count++
		} else {
			return 0, errors.New("Not found element with index ")
		}
	}
	return 0, errors.New("Not found element with index ")
}
