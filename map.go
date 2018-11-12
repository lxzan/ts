package ts

import (
	"sync"
	"sync/atomic"
)

type any map[string]interface{}

type Map struct {
	data  []any
	mutex []*sync.RWMutex
	size  int
	len   int64
}

func NewMap() *Map {
	var s = 64
	obj := &Map{
		data:  make([]any, s),
		mutex: make([]*sync.RWMutex, s),
		size:  64,
		len:   0,
	}
	for i := 0; i < s; i++ {
		obj.data[i] = any{}
		obj.mutex[i] = &sync.RWMutex{}
	}
	return obj
}

func (this *Map) Set(k string, v interface{}) {
	if k == "" {
		panic("key can't be empty")
	}
	var index = int(k[0]) % this.size
	this.mutex[index].Lock()
	this.data[index][k] = v
	atomic.AddInt64(&this.len, 1)
	this.mutex[index].Unlock()
}

func (this *Map) Get(k string) (v interface{}, exist bool) {
	if k == "" {
		panic("key can't be empty")
	}
	var index = int(k[0]) % this.size
	this.mutex[index].RLock()
	v, exist = this.data[index][k]
	this.mutex[index].RUnlock()
	return
}
