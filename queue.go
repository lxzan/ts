package ts

import (
	"sync"
	"sync/atomic"
)

type Queue struct {
	mutex *sync.RWMutex
	first *listNode
	last  *listNode
	len   int64
}

func NewQueue() *Queue {
	var list = &Queue{
		first: &listNode{},
		last:  &listNode{},
		mutex: &sync.RWMutex{},
	}
	return list
}

func (this *Queue) Push(v interface{}) {
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

func (this *Queue) Front() interface{} {
	this.mutex.Lock()
	var val = this.first.val
	if this.first.next != nil {
		val = this.first.val
		next := this.first.next
		next.prev = nil
		this.first = next
	}
	if val != nil {
		atomic.AddInt64(&this.len, -1)
	}
	this.mutex.Unlock()
	return val
}

func (this *Queue) Length() int64 {
	var len int64 = 0
	this.mutex.RLock()
	len = this.len
	this.mutex.RUnlock()
	return len
}
