/*
 * HomeWork-4: Doubly Linked List
 * Created on 23.09.19 19:18
 * Copyright (c) 2019 - Eugene Klimov
 */

package linkedlist

import "testing"

// listActions calls a function of the linked list.
type listActions func(*testing.T, *List)

var newListTestCases = []struct {
	name      string
	in        []interface{}
	out       []interface{}
	expLength int
}{
	{
		name:      "10 items",
		in:        []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		out:       []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		expLength: 10,
	},
	{
		name:      "2 items",
		in:        []interface{}{1, 3},
		out:       []interface{}{1, 3},
		expLength: 2,
	},
	{
		name:      "no items",
		in:        []interface{}{},
		out:       []interface{}{},
		expLength: 0,
	},
	{
		name:      "1 item",
		in:        []interface{}{999},
		out:       []interface{}{999},
		expLength: 1,
	},
}

var RemoveTestCases = []struct {
	name        string
	in          []interface{}
	out         []interface{}
	RemoveNth   int
	expLength   int
	expectedErr error
}{
	{
		name:      "10 items, remove 5th",
		in:        []interface{}{1, 2, 3, 4, 5, 6, 7, 88, 99, 1000},
		RemoveNth: 5,
		out:       []interface{}{1, 2, 3, 4, 6, 7, 88, 99, 1000},
		expLength: 9,
	},
	{
		name:        "no items, no remove",
		in:          []interface{}{},
		RemoveNth:   1,
		out:         []interface{}{},
		expLength:   0,
		expectedErr: ErrEmptyList,
	},
	{
		name:      "1 items, remove it",
		in:        []interface{}{1000},
		RemoveNth: 1,
		out:       []interface{}{},
		expLength: 0,
	},
	{
		name:      "2 items, remove last",
		in:        []interface{}{99, 1000},
		RemoveNth: 2,
		out:       []interface{}{99},
		expLength: 1,
	},
	{
		name:      "2 items, remove first",
		in:        []interface{}{99, 1000},
		RemoveNth: 1,
		out:       []interface{}{1000},
		expLength: 1,
	},
}

var pushPopTestCases = []struct {
	name     string
	in       []interface{}
	actions  []listActions
	expected []interface{}
}{
	{
		name: "PushFront only",
		in:   []interface{}{},
		actions: []listActions{
			pushFront(4),
			pushFront(3),
			pushFront(2),
			pushFront(1),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PushBack only",
		in:   []interface{}{},
		actions: []listActions{
			pushBack(1),
			pushBack(2),
			pushBack(3),
			pushBack(4),
		},
		expected: []interface{}{1, 2, 3, 4},
	},
	{
		name: "PopFront some items",
		in:   []interface{}{1, 2, 3, 4},
		actions: []listActions{
			popFront(1, nil),
			popFront(2, nil),
		},
		expected: []interface{}{3, 4},
	},
	{
		name: "PopFront till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []listActions{
			popFront(1, nil),
			popFront(2, nil),
			popFront(3, nil),
			popFront(4, nil),
			popFront(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "PopBack some items",
		in:   []interface{}{1, 2, 3, 4},
		actions: []listActions{
			popBack(4, nil),
			popBack(3, nil),
		},
		expected: []interface{}{1, 2},
	},
	{
		name: "PopBack till empty",
		in:   []interface{}{1, 2, 3, 4},
		actions: []listActions{
			popBack(4, nil),
			popBack(3, nil),
			popBack(2, nil),
			popBack(1, nil),
			popBack(nil, ErrEmptyList),
		},
		expected: []interface{}{},
	},
	{
		name: "mixed actions",
		in:   []interface{}{2, 3},
		actions: []listActions{
			pushFront(1),
			pushBack(4),
			popFront(1, nil),
			popFront(2, nil),
			popBack(4, nil),
			popBack(3, nil),
			popBack(nil, ErrEmptyList),
			popFront(nil, ErrEmptyList),
			pushFront(8),
			pushBack(7),
			pushFront(9),
			pushBack(6),
		},
		expected: []interface{}{9, 8, 7, 6},
	},
}

func pushFront(arg interface{}) listActions {
	return func(t *testing.T, l *List) {
		l.PushFront(arg)
	}
}

func pushBack(arg interface{}) listActions {
	return func(t *testing.T, l *List) {
		l.PushBack(arg)
	}
}

func popFront(expected interface{}, expectedErr error) listActions {
	return func(t *testing.T, l *List) {
		v, err := l.PopFront()
		if err != expectedErr {
			t.Errorf("PopFront() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopFront() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}

func popBack(expected interface{}, expectedErr error) listActions {
	return func(t *testing.T, l *List) {
		v, err := l.PopBack()
		if err != expectedErr {
			t.Errorf("PopBack() returned wrong, expected no error, got= %v", err)
		}

		if expectedErr == nil && v != expected {
			t.Errorf("PopBack() returned wrong, expected= %v, got= %v", expected, v)
		}
	}
}
