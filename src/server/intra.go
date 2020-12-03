package server

import (
	"context"
	"gen-go/kvs"
	"log"
	"net"
	"os"
	"strconv"
	"thrift/lib/go/thrift"
)

type IntraNodeComm struct {
	rep   kvs.Replica
	n     *kvs.Node
	trans *thrift.TSocket
}

func NewIntraNodeComm(node *kvs.Node) *IntraNodeComm {
	return &IntraNodeComm{n: node}
}

func (h *IntraNodeComm) setuprep() {
	var err error
	h.trans, err = thrift.NewTSocket(net.JoinHostPort(h.n.IP, strconv.Itoa(int(h.n.Port))))
	if err != nil {
		log.Fatalf("failed to create socket to %v:%v [%v]", h.n.IP, h.n.Port, err)
		os.Exit(1)
	}
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	h.rep = kvs.NewReplicaClientFactory(h.trans, protocolFactory)
	if err := h.trans.Open(); err != nil {
		log.Fatalf("Error opening socket to %v:%v [%v]", h.n.IP, h.n.Port, err)
		os.Exit(1)
	}
}

func (h *IntraNodeComm) closerep() {
	h.trans.Close()
}

func (h *IntraNodeComm) Get(c context.Context, key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	return
}
func (h *IntraNodeComm) Put(c context.Context, key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return
}

func (h *IntraNodeComm) GetDataFromNode(c context.Context, key int32) (r *kvs.KVData, err error) {
	h.setuprep()
	defer h.closerep()
	r, err = h.rep.GetDataFromNode(c, key)
	return
}
func (h *IntraNodeComm) PutDataInNode(c context.Context, data *kvs.KVData) (err error) {
	h.setuprep()
	defer h.closerep()
	err = h.rep.PutDataInNode(c, data)
	return
}
func (h *IntraNodeComm) GetHints(c context.Context, node *kvs.Node) (r []*kvs.KVData, err error) {
	h.setuprep()
	defer h.closerep()
	r, err = h.rep.GetHints(c, node)
	return
}
