package main

import (
	"container/list"
	_ "syscall"
)

type line struct {
	li      *list.List
	head    int
	tail    int
	MAXSIZE int
}

func createLine(length int) *line {
	L := new(line)
	L.head = 0
	L.tail = 0
	L.MAXSIZE = length
	return L
}

func (L *line) Push(v interface{}) {
	L.li.PushBack(v)
	L.tail = (L.tail + 1) % L.MAXSIZE
}

func (L *line) Pop(v interface{}) {
	L.head = (L.head + 1) % L.MAXSIZE
}

func (L *line) getFront() interface{} {
	return L.li.Front()
}
