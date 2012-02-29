package main

type Buffer struct {
	i int
	size int
	data []int
}

func (b *Buffer) Push(k int) (bool) {
	for i:=0; i < b.i ; i++{
		if b.data[i] == k {
			return false
		}
	}

	if b.i == b.size {
		b.Pop()
		b.data = append(b.data, k)
	} else {
		b.data[b.i] = k
	}
	b.i++
	return true
}

func (b *Buffer) Pop() {
	b.data = b.data[1:]
	b.i--
}


func NewBuffer (size int) (*Buffer) {
	return &Buffer{0, size, make([]int, size, size)}
}
