package server

import (
	"testing"
	"time"
)

func TestMemstore(t *testing.T) {
	var st memstore
	st.Init()
	st.put(&KVData{3, "qwe", time.Now()})
	d, ok := st.get(3)
	if !ok || d.value != "qwe" {
		t.Fatalf("bad value in memstore key=%v value=%v", 3, d.value)
	}

	d, ok = st.get(4)
	if ok {
		t.Fatal("bad value in memstore 2")
	}
}
