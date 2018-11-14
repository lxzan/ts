package ts

import "testing"

func TestList_Push(t *testing.T) {
	list := NewStack()
	list.Push(1)
	list.Push(2)
	list.Push(3)

	for list.Length() > 0 {
		node := list.Pop()
		v, _ := node.(int)
		println(v)
	}
}

func BenchmarkStack_Push(b *testing.B) {
	var stack = NewStack()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}
