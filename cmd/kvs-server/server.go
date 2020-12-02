package main

import (
	"fmt"
	"gen-go/kvs"
	"thrift/lib/go/thrift"
)

func runServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory) error {
	var transport thrift.TServerTransport
	var err error

	transport, err = thrift.NewTServerSocket("0.0.0.0")

	if err != nil {
		return err
	}
	fmt.Printf("%T\n", transport)
	handler := NewKVSHandler()
	processor := kvs.NewReplicaProcessor(handler)
	serverp := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ")
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
