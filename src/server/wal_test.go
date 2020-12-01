package server

import (
	"fmt"
	"testing"
	"time"
)

func TestWal(t *testing.T) {
	var w wal
	defer w.Close()
	w.Init()
	d := KVData{key: 1, value: "one", ts: time.Now()}
	w.Put(d)
	d = KVData{key: 2, value: "two", ts: time.Now()}
	w.Put(d)
	d = KVData{key: 3, value: "three", ts: time.Now()}
	w.Put(d)

	w.Begin()
	var ret int
	var dp *KVData
	for {
		dp, ret = w.Read()
		if ret == 0 {
			break
		}
		_ = dp
		fmt.Printf("read: %v %v %v\n", dp.key, dp.value, dp.ts)
	}
}
