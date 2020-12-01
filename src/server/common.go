package server

import (
	"fmt"
	"log"
	"net"
	"time"
)

type KVData struct {
	key   uint8
	value string
	ts    time.Time
}

type NodeID int

type NodeInfo struct {
	nodeid NodeID
	ipaddr net.IP
	port   int
}

func (n KVData) String() string {
	t, err := n.ts.MarshalText()
	if err != nil {
		log.Fatal("KVData timestamp conversion failed")
	}
	return fmt.Sprintf("%v,%v,%v\n", n.key, n.value, string(t))
}

type ksio interface {
	put(*KVData)
	get() *KVData
}

type ksmodule interface {
	Init()
}
