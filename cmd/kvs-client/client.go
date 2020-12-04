package main

import (
	"context"
	"flag"
	"fmt"
	"gen-go/kvs"
	"math/rand"
	"net"
	"os"
	"server"
	"strconv"
	"thrift/lib/go/thrift"
)

var _ = kvs.GoUnusedProtection__

var myserver *kvs.Node

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [config file]")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

var cmdusage string = "Bad format" +
	"\nget [key[0-255]] [consistency level[ONE|QUORUM]]\n" +
	"put\n [key[0-255]] [value] [consistency level[ONE|QUORUM]"

func main() {
	flag.Usage = Usage
	var host string
	var port int32
	var config string
	var trans thrift.TTransport
	flag.Usage = Usage
	flag.StringVar(&config, "c", "kvs.config", "Spcify the config file")
	var nodes [4]kvs.Node
	server.ReadConfig(config, &nodes)
	myserver = &nodes[rand.Intn(3)]
	host = myserver.IP
	port = myserver.Port
	portStr := strconv.Itoa(int(port))
	var err error
	trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}
	defer trans.Close()
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iprot := protocolFactory.GetProtocol(trans)
	oprot := protocolFactory.GetProtocol(trans)
	client := kvs.NewReplicaClient(thrift.NewTStandardClient(iprot, oprot))
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}
	fcons := func(s *string) kvs.ConsistencyLevel {
		switch *s {
		case "ONE":
			return kvs.ConsistencyLevel_ONE
		case "QUORUM":
			return kvs.ConsistencyLevel_QUORUM
		default:
			return kvs.ConsistencyLevel_INVALID
		}
	}
	for {
		var arg1 string
		var arg2 string
		var arg3 string
		var arg4 string
		var argn int
		fmt.Printf("kvs-shell $ ")
		argn, err := fmt.Scan(arg1, arg2, arg3, arg4)
		if err != nil {
			fmt.Println(cmdusage)
		}

		switch arg1 {
		case "get":
			if argn != 3 {
				fmt.Println(cmdusage)
				break
			}
			intarg1, err := strconv.Atoi(arg2)
			if err != nil {
				fmt.Println(cmdusage)
				break
			}
			intarg2 := fcons(&arg3)
			if intarg2 == kvs.ConsistencyLevel_INVALID {
				fmt.Println(cmdusage)
				break
			}
			fmt.Print(client.Get(context.Background(), int32(intarg1), intarg2))
		case "put":
			if argn != 4 {
				fmt.Println(cmdusage)
				break
			}
			intarg1, err := strconv.Atoi(arg2)
			if err != nil {
				fmt.Println(cmdusage)
				break
			}
			intarg2 := fcons(&arg4)
			fmt.Print(client.Put(context.Background(), int32(intarg1), arg3, intarg2))
			fmt.Print("\n")
		default:
		}
		fmt.Print("\n")
	}
}
