package server

import (
	"bufio"
	"gen-go/kvs"
	"log"
	"os"
	"strconv"
	"sync"
)

type wal struct {
	f     *os.File
	fread *os.File
	scan  *bufio.Scanner
}

var wo wal
var w *wal
var walmtx sync.Mutex

func WalInit() {
	w = &wo
	f, err := os.OpenFile("./wal-"+strconv.Itoa(int(me)), os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal("WAL init failed", err)
	}
	w.f = f
}

func WalBegin() {
	fread, err := os.Open("./wal-" + strconv.Itoa(int(me)))
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

func WalRead() (*kvs.KVData, int) {
	var d *kvs.KVData
	status := 0
	if w.scan.Scan() {
		d = UnMarshalKVData(w.scan.Text())
		status = 1
	}

	return d, status
}

func WalPut(d *kvs.KVData) {
	str := MarshalKVData(d)
	walmtx.Lock()
	defer walmtx.Unlock()
	n, err := w.f.WriteString(str)
	if err != nil {
		log.Print("WAL write failed\n")
		log.Print(err)
	}
	_ = n
	w.f.Sync()
}

func WalClose() {
	w.f.Close()
	w.fread.Close()
}
