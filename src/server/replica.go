package server

import (
	"bufio"
	"gen-go/kvs"
	"log"
	"os"
	"strconv"
	"strings"
)

var nodes [4]kvs.Node
var me int32
var MeListenAddr string

// TODO init me

func next(start uint) uint {
	return (start + 1) % uint(len(nodes))
}

func GetReplicasForKey(key uint8) []*kvs.Node {
	pri := uint(key / 64)
	dumb := make([]*kvs.Node, 0)
	return append(dumb, &nodes[pri], &nodes[next(pri)], &nodes[next(pri+1)])
}

func GetEveryone() []kvs.Node {
	return nodes[:]
}

func ReadConfig(configFile string, nodes *[4]kvs.Node) {
	file, err := os.Open(configFile)
	defer file.Close()
	if err != nil {
		log.Fatalf("Failed to open config file %v\n", configFile)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	index := int32(0)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ",")
		port, _ := strconv.ParseInt(parts[1], 10, 32)
		(*nodes)[index] = kvs.Node{ID: index, IP: parts[0], Port: int32(port)}
		index += 1
	}
	return
}

func InitNodeInfo(configFile string, m int32) {
	ReadConfig(configFile, &nodes)
	MeListenAddr = nodes[m].IP + ":" + strconv.Itoa(int(nodes[m].Port))
	me = m
}
