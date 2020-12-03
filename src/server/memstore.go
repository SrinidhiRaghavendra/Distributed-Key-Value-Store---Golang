package server

import (
	"gen-go/kvs"
	"log"
	"time"
)

type memstore struct {
	store map[int32]kvs.KVData
}

var m memstore

func MemstoreInit() {
	m.store = make(map[int32]kvs.KVData)
}

func MemstorePut(d *kvs.KVData) {
	if val, ok := m.store[d.Key]; ok {
		oldts, err := time.Parse(time.UnixDate, val.Timestamp)
		if err != nil {
			log.Fatalf("Error parsing timestamp %v", val)
		}
		newts, derr := time.Parse(time.UnixDate, d.Timestamp)
		if derr != nil {
			log.Fatalf("Error parsing timestamp %v", d.Timestamp)
		}
		if !oldts.After(newts) {
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
