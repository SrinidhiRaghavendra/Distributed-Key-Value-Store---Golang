package server

import (
	"bufio"
	"gen-go/kvs"
	"log"
	"os"
	"strconv"
)

var hints [4]*os.File

func HintInit() {
	var err error
	for i, _ := range hints {
		if i != int(me) {
			hints[i], err = os.Create("hints-" + strconv.Itoa(int(me)) + ":" + strconv.Itoa(int(i)))
			if err != nil {
				log.Fatalf("error creating hint file %v", "hints-"+strconv.Itoa(int(me))+":"+strconv.Itoa(int(i)))
			}
		}
	}
}

func HintFlush(id int32) {
	hints[id].Truncate(0)
}

func StoreHint(id int32, hint kvs.KVData) {
	str := MarshalKVData(&hint)
	n, err := hints[id].WriteString(str)
	if err != nil {
		log.Fatal("WAL write failed\n")
	}
	_ = n
	hints[id].Sync()
}

func GetHintsForNode(id int32) []*kvs.KVData {
	scan := bufio.NewScanner(hints[id])
	var d []*kvs.KVData
	for scan.Scan() {
		d = append(d, UnMarshalKVData(scan.Text()))
	}
	return d
}
