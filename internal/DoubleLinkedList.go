package internal

import "fmt"

type DoubleLinkedListNode[T interface{}] struct {
	PreNode  *DoubleLinkedListNode[T]
	NextNode *DoubleLinkedListNode[T]
	Key      string
	Value    T
}

type DoubleLinkedList[T interface{}] struct {
	head *DoubleLinkedListNode[T]
	tail *DoubleLinkedListNode[T]
	size int
}

func NewDoubleLinkedList[T interface{}]() *DoubleLinkedList[T] {
	head := &DoubleLinkedListNode[T]{
		PreNode:  nil,
		NextNode: nil,
	}
	tail := &DoubleLinkedListNode[T]{
		PreNode:  nil,
		NextNode: nil,
	}

	head.NextNode = tail
	tail.PreNode = head

	dll := &DoubleLinkedList[T]{
		size: 0,
		head: head,
		tail: tail,
	}
	return dll
}

// add to tail
func (dll *DoubleLinkedList[T]) Add(key string, value T) {
	node := &DoubleLinkedListNode[T]{
		PreNode:  dll.tail.PreNode,
		NextNode: dll.tail,
		Value:    value,
		Key:      key,
	}
	dll.tail.PreNode.NextNode = node
	dll.tail.PreNode = node
	dll.size++
}

func (dll *DoubleLinkedList[T]) DeleteFirst() {
	if dll.size == 0 {
		return
	}
	dll.head.NextNode.PreNode = dll.head
	dll.head.NextNode = dll.head.NextNode.NextNode
	dll.size--
}

func (dll *DoubleLinkedList[T]) DeleteLast() {
	if dll.size == 0 {
		return
	}
	dll.tail.PreNode.NextNode = dll.tail
	dll.tail.PreNode = dll.tail.PreNode.PreNode
	dll.size--
}

func (dll *DoubleLinkedList[T]) Delete(key string) {
	if dll.size == 0 {
		return
	}
	node := dll.head.NextNode
	for node != dll.tail {
		if node.Key == key {
			node.PreNode.NextNode = node.NextNode
			node.NextNode.PreNode = node.PreNode
			dll.size--
			return
		}
		node = node.NextNode
	}
}

func (dll *DoubleLinkedList[T]) GetFirst() *DoubleLinkedListNode[T] {
	if dll.size == 0 {
		return nil
	}
	return dll.head.NextNode
}

func (dll *DoubleLinkedList[T]) GetLast() *DoubleLinkedListNode[T] {
	if dll.size == 0 {
		return nil
	}
	return dll.tail.PreNode
}

func (dll *DoubleLinkedList[T]) Traverse() {
	node := dll.head.NextNode
	for node != dll.tail {
		fmt.Println(node.Value)
		node = node.NextNode
	}
}
