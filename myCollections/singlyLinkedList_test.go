package myCollections

import (
	//"myCollections"
	"testing"
)

func TestGetValue(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	var want int64 = 777
	got, _ := ml.GetValue(0)
	if want != got {
		t.Errorf("ml.GetValue(0) = %d; want %d", got, want)
	}
}

func TestPrepend(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	ml.Prepend(4)
	ml.Prepend(3)
	ml.Prepend(2)
	ml.Prepend(1)
	ml.Prepend(0)
	var want [5]int64 = [5]int64{0, 1, 2, 3, 4}
	var got [5]int64
	got[0], _ = ml.GetValue(0)
	got[1], _ = ml.GetValue(1)
	got[2], _ = ml.GetValue(2)
	got[3], _ = ml.GetValue(3)
	got[4], _ = ml.GetValue(4)
	if want != got {
		t.Errorf("ml.GetValue(0) = %d; want %d", got[0], want[0])
		t.Errorf("ml.GetValue(1) = %d; want %d", got[1], want[1])
		t.Errorf("ml.GetValue(2) = %d; want %d", got[2], want[2])
		t.Errorf("ml.GetValue(3) = %d; want %d", got[3], want[3])
		t.Errorf("ml.GetValue(4) = %d; want %d", got[4], want[4])
	}
}

func TestAppend(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	ml.Append(1)
	ml.Append(2)
	ml.Append(3)
	ml.Append(4)
	ml.Append(5)
	var want [5]int64 = [5]int64{777, 1, 2, 3, 4}
	var got [5]int64
	got[0], _ = ml.GetValue(0)
	got[1], _ = ml.GetValue(1)
	got[2], _ = ml.GetValue(2)
	got[3], _ = ml.GetValue(3)
	got[4], _ = ml.GetValue(4)
	if want != got {
		t.Errorf("ml.GetValue(0) = %d; want %d", got[0], want[0])
		t.Errorf("ml.GetValue(1) = %d; want %d", got[1], want[1])
		t.Errorf("ml.GetValue(2) = %d; want %d", got[2], want[2])
		t.Errorf("ml.GetValue(3) = %d; want %d", got[3], want[3])
		t.Errorf("ml.GetValue(4) = %d; want %d", got[4], want[4])
	}
}

func TestRemoveByValue(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	ml.Append(1)
	ml.Append(2)
	ml.Append(3)
	ml.Append(4)
	ml.Append(5)
	var want [2]int64 = [2]int64{2, 3}
	var got [2]int64
	ml.RemoveByValue(777)
	ml.RemoveByValue(1)
	ml.RemoveByValue(4)
	got[0], _ = ml.GetValue(0)
	got[1], _ = ml.GetValue(1)
	if want != got {
		t.Errorf("ml.GetValue(0) = %d; want %d", got[0], want[0])
		t.Errorf("ml.GetValue(1) = %d; want %d", got[1], want[1])
	}
}

func TestRemoveByIndex(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	ml.Append(1)
	ml.Append(2)
	ml.Append(3)
	ml.Append(4)
	ml.Append(5)
	var want [3]int64 = [3]int64{1, 3, 4}
	var got [3]int64
	ml.RemoveByIndex(0)
	ml.RemoveByIndex(4)
	ml.RemoveByIndex(1)
	got[0], _ = ml.GetValue(0)
	got[1], _ = ml.GetValue(1)
	got[2], _ = ml.GetValue(2)
	if want != got {
		t.Errorf("ml.GetValue(0) = %d; want %d", got[0], want[0])
		t.Errorf("ml.GetValue(1) = %d; want %d", got[1], want[1])
		t.Errorf("ml.GetValue(2) = %d; want %d", got[2], want[2])
	}
}

func TestSize(t *testing.T) {
	ml := NewSinglyLinkedList(777)
	ml.Append(1)
	ml.Append(2)
	ml.Append(3)
	ml.Append(4)
	ml.Append(5)
	var want uint64 = 6
	got := ml.Size()
	if want != got {
		t.Errorf("ml.Size = %d; want %d", got, want)
	}
}
