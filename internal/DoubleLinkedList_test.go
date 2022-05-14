package internal

import "testing"

func TestAdd(t *testing.T) {

	dll := NewDoubleLinkedList[string]()

	dll.Add("a", "s111_a")
	dll.Add("b", "s111_b")
	dll.Add("c", "s111_c")
	dll.Add("d", "s111_d")

	dll.Traverse()

	if dll.GetLast().Value != "s111_d" {
		t.Error("GetFirst error")
	}

	dll.DeleteLast()

	dll.Traverse()

	if dll.GetLast().Value != "s111_c" {
		t.Error("GetLast error")
	}

	dll.DeleteFirst()

	dll.Traverse()

	if dll.GetFirst().Value != "s111_b" {
		t.Error("GetFirst error")
	}
}
