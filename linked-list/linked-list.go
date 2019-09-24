/*
 * HomeWork-4: Doubly Linked List
 * Created on 20.09.19 23:04
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package linkedlist implements a doubly linked list.
package linkedlist

import "errors"

// ErrEmptyList is an error to return on an operation working on an empty list.
var ErrEmptyList = errors.New("list is empty")

// Item is a node of the doubly linked list
type Item struct {
	data       interface{}
	prev, next *Item
}

// List is a structure representing a doubly linked list.
type List struct {
	first, last *Item
	length      int
}

// First returns pointer to the node at front of the list.
func (l *List) First() *Item {
	return l.first
}

// Last returns pointer to the node at back of the list.
func (l *List) Last() *Item {
	return l.last
}

// Next returns a pointer to the next node.
func (i *Item) Next() *Item {
	return i.next
}

// Prev returns a pointer to the previous node.
func (i *Item) Prev() *Item {
	return i.prev
}

// Len returns length of the list.
func (l *List) Len() int {
	return l.length
}

// NewList creates new linked list with the data items.
func NewList(data ...interface{}) *List {
	l := &List{}
	for _, item := range data {
		l.PushBack(item)
	}
	return l
}

// Remove deletes item from the list.
func (l *List) Remove(i *Item) error {
	if l.length == 0 || i == nil {
		return ErrEmptyList
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
	return nil
}

// PushBack pushes item to end of the list.
func (l *List) PushBack(v interface{}) {
	node := &Item{data: v, prev: l.last}
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
		return nil, ErrEmptyList
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
	node := &Item{data: v, next: l.first}
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
		return nil, ErrEmptyList
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
