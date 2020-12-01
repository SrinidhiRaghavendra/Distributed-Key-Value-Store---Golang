package server

import (
	"bufio"
	"fmt"
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
	fread, err := os.Open("./wal.bin")
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

func (w *wal) Read() (*KVData, int) {
	var d KVData
	var l []string
	status := 0
	if w.scan.Scan() {
		l = strings.SplitN(w.scan.Text(), ",", 3)
		v, err := strconv.Atoi(l[0])
		if err != nil {
			log.Fatal("Error reading from WAL")
		}
		d.key = uint8(v)
		d.value = l[1]
		fmt.Printf("Read() getting %v\n", l[2])
		d.ts.UnmarshalText([]byte(l[2]))
		status = 1
	}

	return &d, status
}

func (w *wal) Put(d KVData) {
	str := d.String()
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
