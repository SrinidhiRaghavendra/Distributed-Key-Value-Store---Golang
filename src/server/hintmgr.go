package server

import (
	"bufio"
	"gen-go/kvs"
	"log"
	"os"
	"strconv"
	"sync"
)

var hints [4]*os.File
var hintmtx sync.Mutex

func HintInit() {
	var err error
	for i, _ := range hints {
		if i != int(me) {
			hints[i], err = os.OpenFile("hints-"+strconv.Itoa(int(me))+":"+strconv.Itoa(int(i)), os.O_CREATE|os.O_RDWR, 0644)
			if err != nil {
				log.Printf("error creating hint file %v", "hints-"+strconv.Itoa(int(me))+":"+strconv.Itoa(int(i)))
			}
		}
	}
}

func HintFlush(id int32) {
	hints[id].Truncate(0)
}

func StoreHint(id int32, hint kvs.KVData) {
	str := MarshalKVData(&hint)

	hintmtx.Lock()
	defer hintmtx.Unlock()
	n, err := hints[id].WriteString(str)
	if err != nil {
		log.Fatal("WAL write failed\n")
	}
	_ = n
	hints[id].Sync()
}

func GetHintsForNode(id int32) []*kvs.KVData {
	f, err := os.Open("hints-" + strconv.Itoa(int(me)) + ":" + strconv.Itoa(int(id)))
	if err != nil {
		log.Printf("error opening hint file %v", "hints-"+strconv.Itoa(int(me))+":"+strconv.Itoa(int(id)))
	}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanLines)
	var d []*kvs.KVData
	for scan.Scan() {
		x := UnMarshalKVData(scan.Text())
		d = append(d, x)
	}
	return d
}
