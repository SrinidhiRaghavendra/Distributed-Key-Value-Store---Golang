package server

import (
	"gen-go/kvs"
)

var nodes [4]kvs.Node
var me uint8

// TODO init me

func next(start uint) uint {
	return (start + 1) % 3
}

func GetReplicasForKey(key uint8) []*kvs.Node {
	pri := uint(key / 64)
	dumb := make([]*kvs.Node, 0)
	return append(dumb, &nodes[pri], &nodes[next(pri)], &nodes[next(pri+1)])
}

func GetEveryone() []kvs.Node {
	return nodes[:]
}

func InitNodeInfo(m uint8) {
	nodes[0] = kvs.Node{ID: 0, IP: "127.0.0.1", Port: 8080}
	nodes[1] = kvs.Node{ID: 1, IP: "127.0.0.1", Port: 8081}
	nodes[2] = kvs.Node{ID: 2, IP: "127.0.0.1", Port: 8082}
	nodes[3] = kvs.Node{ID: 3, IP: "127.0.0.1", Port: 8083}
	me = m
}
