// Package linkedlist implements a doubly linked list.
// It is replacement builtin golang list
// but more concurrency safe.
package linkedlist

import (
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/errors"
)

// Element is a node of the doubly linked list
type Element struct {
	data       interface{}
	prev, next *Element
}

// List is a structure representing a doubly linked list.
type List struct {
	// no need insert lock on every method right now
	// see unit tests and add needed lock after
	sync.RWMutex
	first, last *Element
	length      int
}

// Front returns pointer to the node at front of the list.
func (l *List) Front() *Element {
	l.Lock()
	defer l.Unlock()

	return l.first
}

// Back returns pointer to the node at back of the list.
func (l *List) Back() *Element {
	return l.last
}

// Next returns a pointer to the next node.
func (i *Element) Next() *Element {
	return i.next
}

// Prev returns a pointer to the previous node.
func (i *Element) Prev() *Element {
	return i.prev
}

// Len returns length of the list.
func (l *List) Len() int {
	return l.length
}

// Value returns data of the item.
func (i *Element) Value() interface{} {
	return i.data
}

// New creates new linked list with the data items.
func New(data ...interface{}) *List {
	l := &List{}

	for _, item := range data {
		l.PushBack(item)
	}

	return l
}

// Remove deletes item from the list.
func (l *List) Remove(i *Element) {
	if l.length == 0 || i == nil {
		return
	}

	if i.prev != nil {
		i.prev.next = i.next
	} else {
		l.first = i.next
	}

	if i.next != nil {
		i.next.prev = i.prev
	} else {
		l.last = i.prev
	}
	l.length--
}

// PushBack pushes item to end of the list.
func (l *List) PushBack(v interface{}) {
	node := &Element{data: v, prev: l.last}
	l.length++

	if l.first == nil {
		l.first, l.last = node, node
		return
	}

	l.last, l.last.next = node, node
}

// PopBack pops item from end of the list.
func (l *List) PopBack() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.ErrEmptyList
	}

	v := l.last
	l.last = l.last.prev

	if l.last == nil {
		l.first = nil
	} else {
		l.last.next = nil
	}
	l.length--

	return v.data, nil
}

// PushFront pushes item to begin of the list.
func (l *List) PushFront(v interface{}) {
	l.Lock()
	defer l.Unlock()

	node := &Element{data: v, next: l.first}
	l.length++

	if l.first == nil {
		l.first, l.last = node, node
		return
	}

	l.first, l.first.prev = node, node
}

// PopFront pops item from end of the list.
func (l *List) PopFront() (interface{}, error) {
	if l.length == 0 {
		return nil, errors.ErrEmptyList
	}

	v := l.first
	l.first = l.first.next

	if l.first == nil {
		l.last = nil
	} else {
		l.first.prev = nil
	}
	l.length--

	return v.data, nil
}
