package server

import (
	"fmt"
	"gen-go/kvs"
)

//type string int

//type kvs.Node struct {
//	ID string
//	IP string
//	Port   string
//}

// omit warn
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
