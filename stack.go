package ts

import (
	"sync"
	"sync/atomic"
)

type listNode struct {
	val  interface{}
	prev *listNode
	next *listNode
}

type Stack struct {
	mutex *sync.RWMutex
	first *listNode
	last  *listNode
	len   int64
}

func NewStack() *Stack {
	var list = &Stack{
		first: &listNode{},
		last:  &listNode{},
		mutex: &sync.RWMutex{},
	}
	return list
}

func (this *Stack) Push(v interface{}) {
	this.mutex.Lock()
	var node = &listNode{
		val:  v,
		prev: this.last,
	}
	this.last.next = node
	if this.len == 1 {
		this.first = &listNode{
			val:  this.last.val,
			next: node,
		}
	}
	this.last = node
	atomic.AddInt64(&this.len, 1)
	this.mutex.Unlock()
}

func (this *Stack) Pop() interface{} {
	this.mutex.Lock()
	var val interface{}
	if this.last.prev != nil {
		val = this.last.val
		prev := this.last.prev
		prev.next = nil
		this.last = prev
		atomic.AddInt64(&this.len, -1)
	}
	this.mutex.Unlock()
	return val
}

func (this *Stack) Length() int64 {
	var len int64 = 0
	this.mutex.RLock()
	len = this.len
	this.mutex.RUnlock()
	return len
}
