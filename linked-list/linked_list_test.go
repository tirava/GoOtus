/*
 * HomeWork-4: Doubly Linked List
 * Created on 20.09.19 23:14
 * Copyright (c) 2019 - Eugene Klimov
 */

package linkedlist

import (
	"bytes"
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	for _, tc := range newListTestCases {
		newList := NewList(tc.in...)
		t.Run(tc.name, func(t *testing.T) {
			checkLinkedList(t, newList, tc.out)
		})
		actLength := newList.Len()
		if tc.expLength != actLength {
			t.Errorf("%s\n\tExpected length: %v\n\tGot: %v", tc.name, tc.expLength, actLength)
		}
	}
}

func TestPushPop(t *testing.T) {
	for _, tc := range pushPopTestCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := NewList(tc.in...)
			for _, ac := range tc.actions {
				ac(t, actual)
			}
			checkLinkedList(t, actual, tc.expected)
		})
	}
}

func TestRemove(t *testing.T) {
	for _, tc := range RemoveTestCases {
		newList := NewList(tc.in...)
		t.Run(tc.name, func(t *testing.T) {
			err := newList.Remove(newList.getNthItem(tc.RemoveNth))
			if err != tc.expectedErr {
				t.Errorf("Remove() returned wrong, expected = %v, got = %v", tc.expectedErr, err)
			}
			checkLinkedList(t, newList, tc.out)
		})
		actLength := newList.Len()
		if tc.expLength != actLength {
			t.Errorf("%s\n\tExpected length: %v\n\tGot: %v", tc.name, tc.expLength, actLength)
		}
	}
}

// checkLinkedList checks that the linked list is correct.
func checkLinkedList(t *testing.T, l *List, expected []interface{}) {
	// check that length and elements are correct (scan once from begin -> end)
	item, count, i := l.First(), 0, 0
	for ; item != nil && i < len(expected); item, count, i = item.Next(), count+1, i+1 {
		if item.data != expected[i] {
			t.Errorf("wrong value from %d-th item, expected = %v, got = %v", i, expected[i], item.data)
		}
	}
	if !(item == nil && i == len(expected)) {
		t.Errorf("expected %d items, got = %d", len(expected), count)
	}

	// if items are the same, need to examine links too (next & prev)
	switch {
	case l.First() == nil && l.Last() == nil: // empty list
		return
	case l.First() != nil && l.Last() != nil && l.First().Next() == nil: // 1 element
		valid := l.First() == l.Last() &&
			l.First().Next() == nil &&
			l.First().Prev() == nil &&
			l.Last().Next() == nil &&
			l.Last().Prev() == nil

		if !valid {
			t.Errorf("expected only 1 item and no links, got = %v", l.debugString())
		}
	}

	// > 1 element
	if l.First().Prev() != nil {
		t.Errorf("expected First.Prev() == nil, got = %v", l.First().Prev())
	}

	prev := l.First()
	cur := l.First().Next()
	for i := 0; cur != nil; i++ {
		if !(prev.Next() == cur && cur.Prev() == prev) {
			t.Errorf("%d-th item's links is wrong", i)
		}
		prev = cur
		cur = cur.Next()
	}

	if l.Last().Next() != nil {
		t.Errorf("expected Last().Next() == nil, got = %v", l.Last().Next())
	}
}

// debugString prints the linked list with both node's Val, next & prev pointers.
func (l *List) debugString() string {
	buf := bytes.NewBuffer([]byte{'{'})
	buf.WriteString(fmt.Sprintf("First()= %p; ", l.First()))

	for cur := l.First(); cur != nil; cur = cur.Next() {
		buf.WriteString(fmt.Sprintf("[Prev() = %p, Data = %p (%v), Next() = %p] <-> ", cur.Prev(), cur, cur.data, cur.Next()))
	}

	buf.WriteString(fmt.Sprintf("; Last() = %p; ", l.Last()))
	buf.WriteByte('}')

	return buf.String()
}

// getNthItem helps getting n-th item for Remove tests
func (l *List) getNthItem(n int) *Item {
	for i := l.First(); i != nil; i = i.next {
		n--
		if n == 0 {
			return i
		}
	}
	return nil
}
