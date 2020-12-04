package server

import (
	"gen-go/kvs"
	"sync"
)

type memstore struct {
	store map[int32]kvs.KVData
}

var m memstore
var memmtx sync.Mutex

func MemstoreInit() {
	m.store = make(map[int32]kvs.KVData)
}

func MemstorePut(d *kvs.KVData) {
	memmtx.Lock()
	defer memmtx.Unlock()
	if val, ok := m.store[d.Key]; ok {

		if val.Timestamp <= d.Timestamp {
			m.store[d.Key] = *d
		}
	} else {
		m.store[d.Key] = *d
	}
}

func MemstoreGet(key int32) (*kvs.KVData, bool) {
	var ret kvs.KVData
	e := false
	if val, ok := m.store[key]; ok {
		ret = val
		e = true
	}
	return &ret, e
}
