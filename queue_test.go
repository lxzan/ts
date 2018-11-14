package ts

import "testing"

func TestQueue_Push(t *testing.T) {
	Q := NewQueue()
	Q.Push(1)
	Q.Push(2)
	Q.Push(3)

	for Q.Length() > 0 {
		node := Q.Front()
		v, _ := node.(int)
		println(v)
	}
}

func BenchmarkQueue_Push(b *testing.B) {
	var stack = NewQueue()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}
