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

func (h *IntraNodeComm) Setuprep() (err error) {
	if(me != h.n.ID) {
		if(h.trans != nil && h.trans.IsOpen()){
			return
		}
		var err error
		h.trans, err = thrift.NewTSocket(net.JoinHostPort(h.n.IP, strconv.Itoa(int(h.n.Port))))
		if err != nil {
			log.Printf("failed to create socket to %v:%v [%v]", h.n.IP, h.n.Port, err)
		}
		protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
		h.rep = kvs.NewReplicaClientFactory(h.trans, protocolFactory)
		if err := h.trans.Open(); err != nil {
			log.Printf("Error opening socket to %v:%v [%v]", h.n.IP, h.n.Port, err)
			return err
		}
	} else {
		h.rep = NewKVSHandler()
	}
	return
}

func (h *IntraNodeComm) closerep() {
	if(me != h.n.ID){
		h.trans.Close()
	}
}

func (h *IntraNodeComm) Get(c context.Context, key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	return
}
func (h *IntraNodeComm) Put(c context.Context, key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return
}

func (h *IntraNodeComm) GetDataFromNode(c context.Context, key int32) (r *kvs.KVData, err error) {
	err = h.Setuprep()
	if(err != nil) {
		return nil, err
	}
	defer h.closerep()
	r, err = h.rep.GetDataFromNode(c, key)
	return
}
func (h *IntraNodeComm) PutDataInNode(c context.Context, data *kvs.KVData) (err error) {
	err = h.Setuprep()
	if(err != nil) {
		return err
	}
	defer h.closerep()
	err = h.rep.PutDataInNode(c, data)
	return
}
func (h *IntraNodeComm) GetHints(c context.Context, node *kvs.Node) (r []*kvs.KVData, err error) {
	err = h.Setuprep()
	if(err != nil) {
		return nil, err
	}
	defer h.closerep()
	r, err = h.rep.GetHints(c, node)
	return
}
