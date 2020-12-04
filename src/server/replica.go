package server

import (
	"gen-go/kvs"
	"log"
	"os"
	"bufio"
	"strings"
	"strconv"
)

var nodes [4]kvs.Node
var me int32
var MeListenAddr string
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

func InitNodeInfo(configFile string, m int32) {
	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var index int32 = 0
	//NOTE: The config file is assumed to be in correct format, can add validations later
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Split(text, ",")
		port,_ := strconv.ParseInt(parts[1], 10, 32)
		nodes[index] = kvs.Node{ID: index, IP: parts[0], Port: int32(port)}
		if(index == m) {
			MeListenAddr = parts[0]+":"+parts[1]
		}
		index += 1

	}
	file.Close()
	me = m
}
