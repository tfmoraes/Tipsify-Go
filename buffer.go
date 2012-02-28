package main

import "container/list"

type Buffer struct {
	i int
	size int
	data *list.List
}

func (b *Buffer) Push(k int) (bool) {
	for e:=b.data.Front(); e != nil ; e = e.Next(){
		if e.Value == k {
			return false
		}
	}

	if b.i == b.size {
		b.Pop()
	}
	b.data.PushBack(k)
	b.i++
	return true
}

func (b *Buffer) Pop() {
	b.i--
	b.data.Remove(b.data.Front())
}


func NewBuffer (size int) (*Buffer) {
	return &Buffer{0, size, list.New()}
}
