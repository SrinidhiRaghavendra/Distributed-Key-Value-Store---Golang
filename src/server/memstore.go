package server

import "gen-go/kvs"

type memstore struct {
	store map[int32]kvs.KVData
}

func (m *memstore) Init() {
	m.store = make(map[int32]kvs.KVData)
}

func (m *memstore) put(d *kvs.KVData) {
	m.store[d.Key] = *d
}

func (m *memstore) get(key int32) (*kvs.KVData, bool) {
	var ret kvs.KVData
	e := false
	if val, ok := m.store[key]; ok {
		ret = val
		e = true
	}
	return &ret, e
}
