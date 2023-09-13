package main

import (
	"fmt"
	"reflect"
)

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) Tail() T {
	if l.next == nil {
		return l.val
	}
	return l.next.Tail()
}

func (l *List[T]) LookUp(lv T) *List[T] {
	if reflect.DeepEqual(l.val, lv) { // types may be incomparable; use reflect
		return l
	}
	return l.next.LookUp(lv)
}

func (l *List[T]) Append(ln *List[T]) *List[T] {
	last := l.LookUp(l.Tail())
	last.next = ln
	return l //last
}

func (l *List[T]) PrintAll() {
	fmt.Println(l.val)
	if l.next != nil {
		l.next.PrintAll()
	} else {
		return
	}
}

