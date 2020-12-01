package server

import (
	"fmt"
	"net"
)

var nodes [4]NodeInfo

func (n NodeInfo) String() string {
	return fmt.Sprintf("[nodeid=%v ipaddr=%v port=%v]\n", n.nodeid, n.ipaddr, n.port)
}

func next(start uint) uint {
	return (start + 1) % 3
}

func GetReplicasForKey(key uint8) []*NodeInfo {
	pri := uint(key / 64)
	dumb := make([]*NodeInfo, 0)
	return append(dumb, &nodes[pri], &nodes[next(pri)], &nodes[next(pri+1)])
}

func GetEveryone() []NodeInfo {
	return nodes[:]
}

func InitNodeInfo() {
	nodes[0] = NodeInfo{nodeid: 0, ipaddr: net.IPv4(127, 0, 0, 1), port: 8080}
	nodes[1] = NodeInfo{nodeid: 1, ipaddr: net.IPv4(127, 0, 0, 1), port: 8081}
	nodes[2] = NodeInfo{nodeid: 2, ipaddr: net.IPv4(127, 0, 0, 1), port: 8082}
	nodes[3] = NodeInfo{nodeid: 3, ipaddr: net.IPv4(127, 0, 0, 1), port: 8083}
}
