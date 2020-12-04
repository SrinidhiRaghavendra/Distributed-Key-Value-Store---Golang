package main

import (
	"fmt"
	"gen-go/kvs"
	"thrift/lib/go/thrift"
	"os"
	"server"
	"strconv"
)

func initManagers() {
	//Init Replica
	nodeID, _ := strconv.Atoi(os.Args[2])
	server.InitNodeInfo(os.Args[1], int32(nodeID))
	//Init Memstore
	server.MemstoreInit()
	//Init Hint manager
	server.HintInit()
	//Init WAL
	server.WalInit()
}

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory) error {
	var transport thrift.TServerTransport
	var err error

	initManagers()
	transport, err = thrift.NewTServerSocket(server.MeListenAddr)

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := server.NewKVSHandler()
	processor := kvs.NewReplicaProcessor(handler)
	serverp := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ")
	go server.Recover()
	return serverp.Serve()
}

//func main() {
//	fmt.Println("hello server xD\n");
//	a := server.GetReplicasForKey(4)
//	fmt.Println(a)
//}

func main() {
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	transportFactory := thrift.NewTTransportFactory()

	if err := runServer(transportFactory, protocolFactory); err != nil {
		fmt.Println("error running server:", err)
	}
}
