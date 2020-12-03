package server

import (
	"bufio"
	"fmt"
	"gen-go/kvs"
	"log"
	"os"
	"strconv"
	"strings"
)

type wal struct {
	f     *os.File
	fread *os.File
	scan  *bufio.Scanner
}

func (w *wal) Init() {
	f, err := os.Create("./wal.bin")
	if err != nil {
		log.Fatal("WAL init failed", err)
	}
	w.f = f
}

func (w *wal) Begin() {
	fread, err := os.OpenFile("./wal-"+strconv.Itoa(int(me)), os.O_CREATE, 644)
	if err != nil {
		log.Fatal("WAL begin read failed", err)
	}
	if w.fread != nil {
		w.fread.Close()
	}

	w.fread = fread
	w.scan = bufio.NewScanner(w.fread)
	w.scan.Split(bufio.ScanLines)
}

func MarshalKVData(kvd *kvs.KVData) (s string) {
	return fmt.Sprintf("%v*KVSSEP*%v*KVSSEP*%v\n", kvd.Key, kvd.Value, kvd.Timestamp)

}

func UnMarshalKVData(d string) *kvs.KVData {
	l := strings.SplitN(d, "*KVSSEP*", 3)
	v, err := strconv.Atoi(l[0])
	if err != nil {
		log.Fatal("Error reading from WAL")
	}
	return &kvs.KVData{Key: int32(v), Value: l[1], Timestamp: l[2]}
}

func (w *wal) Read() (*kvs.KVData, int) {
	var d *kvs.KVData
	status := 0
	if w.scan.Scan() {
		d = UnMarshalKVData(w.scan.Text())
		status = 1
	}

	return d, status
}

func (w *wal) Put(d kvs.KVData) {
	str := MarshalKVData(&d)
	n, err := w.f.WriteString(str)
	if err != nil {
		log.Fatal("WAL write failed\n")
	}
	_ = n
	w.f.Sync()
}

func (w *wal) Close() {
	w.f.Close()
	w.fread.Close()
}
