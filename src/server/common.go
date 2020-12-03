package server

import (
	"fmt"
	"gen-go/kvs"
	"log"
	"strconv"
	"strings"
)

type kvsio interface {
	put(*kvs.KVData)
	get() *kvs.KVData
}

type kvsmodule interface {
	Init()
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
