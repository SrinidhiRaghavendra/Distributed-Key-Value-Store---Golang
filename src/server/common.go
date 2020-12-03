package server

import (
	"gen-go/kvs"
)

//type string int

//type kvs.Node struct {
//	ID string
//	IP string
//	Port   string
//}

type ksio interface {
	put(*kvs.KVData)
	get() *kvs.KVData
}

type ksmodule interface {
	Init()
}
