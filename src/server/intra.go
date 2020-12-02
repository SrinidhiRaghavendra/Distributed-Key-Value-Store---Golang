package server

import (
	"gen-go/kvs"
	"log"
	"net"
	"os"
	"thrift/lib/go/thrift"
)

type IntraNodeComm struct {
	rep   kvs.Replica
	n     *NodeInfo
	trans *thrift.TSocket
}

func NewIntraNodeComm(node *NodeInfo) *IntraNodeComm {
	return &IntraNodeComm{n: node}
}

func (h *IntraNodeComm) setuprep() {
	var err error
	h.trans, err = thrift.NewTSocket(net.JoinHostPort(h.n.ipaddr, h.n.port))
	if err != nil {
		log.Fatalf("failed to create socket to %v:%v [%v]", h.n.ipaddr, h.n.port, err)
		os.Exit(1)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	h.rep = kvs.NewReplicaClientFactory(h.trans, protocolFactory)
	if err := h.trans.Open(); err != nil {
		log.Fatalf("Error opening socket to %v:%v [%v]", h.n.ipaddr, h.n.port, err)
		os.Exit(1)
	}
}

func (h *IntraNodeComm) closerep() {
	h.trans.Close()
}

func (h *IntraNodeComm) Get(key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	return
}
func (h *IntraNodeComm) Put(key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return
}

func (h *IntraNodeComm) GetDataFromNode(key int32) (r *kvs.KVData, err error) {
	h.setuprep()
	defer h.closerep()
	r, err = h.rep.GetDataFromNode(key)
	return
}
func (h *IntraNodeComm) PutDataInNode(data *kvs.KVData) (err error) {
	h.setuprep()
	defer h.closerep()
	err = h.rep.PutDataInNode(data)
	return
}
func (h *IntraNodeComm) GetHints(node *kvs.Node) (r []*kvs.KVData, err error) {
	h.setuprep()
	defer h.closerep()
	r, err = h.rep.GetHints(node)
	return
}
