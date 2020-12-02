package server

import (
	"gen-go/kvs"
	"testing"
	"time"
)

func TestMemstore(t *testing.T) {
	var st memstore
	st.Init()
	st.put(&kvs.KVData{Key: 3, Value: "qwe", Timestamp: time.Now().String()})
	d, ok := st.get(3)
	if !ok || d.Value != "qwe" {
		t.Fatalf("bad Value in memstore key=%v Value=%v", 3, d.Value)
	}

	d, ok = st.get(4)
	if ok {
		t.Fatal("bad Value in memstore 2")
	}
}
