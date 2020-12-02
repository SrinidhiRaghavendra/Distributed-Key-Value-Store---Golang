package server

import (
	"fmt"
	"gen-go/kvs"
	"testing"
	"time"
)

func TestWal(t *testing.T) {
	var w wal
	defer w.Close()
	w.Init()
	d := kvs.KVData{Key: 1, Value: "one", Timestamp: time.Now().String()}
	w.Put(d)
	d = kvs.KVData{Key: 2, Value: "two", Timestamp: time.Now().String()}
	w.Put(d)
	d = kvs.KVData{Key: 3, Value: "three", Timestamp: time.Now().String()}
	w.Put(d)

	w.Begin()
	var ret int
	var dp *kvs.KVData
	for {
		dp, ret = w.Read()
		if ret == 0 {
			break
		}
		_ = dp
		fmt.Printf("read: %v %v %v\n", dp.Key, dp.Value, dp.Timestamp)
	}
}
