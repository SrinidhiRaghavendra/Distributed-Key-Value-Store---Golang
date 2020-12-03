package server

import (
	"fmt"
	"gen-go/kvs"
	"testing"
	"time"
)

func TestWal(t *testing.T) {
	WalInit()
	defer WalClose()
	WalInit()
	d := kvs.KVData{Key: 1, Value: "one", Timestamp: time.Now().Unix()}
	WalPut(d)
	d = kvs.KVData{Key: 2, Value: "two", Timestamp: time.Now().Unix()}
	WalPut(d)
	d = kvs.KVData{Key: 3, Value: "three", Timestamp: time.Now().Unix()}
	WalPut(d)

	WalBegin()
	var ret int
	var dp *kvs.KVData
	for {
		dp, ret = WalRead()
		if ret == 0 {
			break
		}
		_ = dp
		fmt.Printf("read: %v %v %v\n", dp.Key, dp.Value, dp.Timestamp)
	}
}
