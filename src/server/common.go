package server

import (
	"fmt"
	"gen-go/kvs"
)

type NodeID int

type NodeInfo struct {
	nodeid NodeID
	ipaddr string
	port   string
}

func MarshalKVData(n kvs.KVData) string {
	return fmt.Sprintf("%v,%v,%v\n", n.Key, n.Value, n.Timestamp)
}

type ksio interface {
	put(*kvs.KVData)
	get() *kvs.KVData
}

type ksmodule interface {
	Init()
}
