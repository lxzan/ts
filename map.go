package ts

import (
	"sync"
)

type any map[string]interface{}

type Map struct {
	data  []any
	mutex []*sync.RWMutex
	Size  int
}

func NewMap() *Map {
	var s = 64
	obj := &Map{
		data:  make([]any, s),
		mutex: make([]*sync.RWMutex, s),
		Size:  s,
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
	var index = int(k[0]) % this.Size
	this.mutex[index].Lock()
	this.data[index][k] = v
	this.mutex[index].Unlock()
}

func (this *Map) Get(k string) (v interface{}, exist bool) {
	if k == "" {
		panic("key can't be empty")
	}
	var index = int(k[0]) % this.Size
	this.mutex[index].RLock()
	v, exist = this.data[0][k]
	this.mutex[index].RUnlock()
	return
}

