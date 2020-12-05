package main

import (
	"bufio"
	"context"
	"fmt"
	"gen-go/kvs"
	"os"
	"strconv"
	"thrift/lib/go/thrift"
)

var _ = kvs.GoUnusedProtection__

var myserver *kvs.Node

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " ip:port")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

var cmdusage string = "Bad format" +
	"\nget [key[0-255]] [consistency level[ONE|QUORUM]]\n" +
	"put\n [key[0-255]] [value] [consistency level[ONE|QUORUM]"

func main() {
	var host string
	var trans thrift.TTransport
	var err error
	if len(os.Args) != 2 {
		Usage()
		os.Exit(1)
	}
	host = os.Args[1]
	trans, err = thrift.NewTSocket(host)
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
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, err)
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
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		line := input.Text()
		argn, err := fmt.Sscanf(line, "%s %s %s %s", &arg1, &arg2, &arg3, &arg4)
		_ = err
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
			s, er := client.Get(context.Background(), int32(intarg1), intarg2)
			if er != nil {
				fmt.Println("Put consistency level not met.")
			} else {
				fmt.Printf(">> %v", s)
			}
			_ = er
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
			err3 := client.Put(context.Background(), int32(intarg1), arg3, intarg2)
			if err3 != nil {
				fmt.Println("Put consistency level not met.")
			} else {
				fmt.Println("Put successfully done.")
			}
		case "exit":
			os.Exit(0)
		default:
		}
		fmt.Print("\n")
	}
}
