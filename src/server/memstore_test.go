package server

import (
	"gen-go/kvs"
	"testing"
	"time"
)

func TestMemstore(t *testing.T) {
	MemstoreInit()
	MemstorePut(&kvs.KVData{Key: 3, Value: "qwe", Timestamp: time.Now().Unix()})
	d, ok := MemstoreGet(3)
	if !ok || d.Value != "qwe" {
		t.Printf("bad Value in memstore key=%v Value=%v", 3, d.Value)
	}

	d, ok = MemstoreGet(4)
	if ok {
		t.Fatal("bad Value in memstore 2")
	}
}

func TestMemstore2(t *testing.T) {
	MemstoreInit()
	t1 := time.Now()
	time.Sleep(4 * time.Second)
	t2 := time.Now()
	MemstorePut(&kvs.KVData{Key: 3, Value: "qwe", Timestamp: t2.Unix()})
	MemstorePut(&kvs.KVData{Key: 3, Value: "asd", Timestamp: t1.Unix()})

	d, ok := MemstoreGet(3)
	if !ok || d.Value != "qwe" {
		t.Printf("bad Value in memstore key=%v Value=%v", 3, d.Value)
	}

	time.Sleep(3 * time.Second)
	t2 = time.Now()
	MemstorePut(&kvs.KVData{Key: 3, Value: "asd", Timestamp: t2.Unix()})
	d, ok = MemstoreGet(3)
	if !ok || d.Value != "asd" {
		t.Printf("bad Value in memstore key=%v Value=%v", 3, d.Value)
	}
}
