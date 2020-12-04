package server

import (
	"gen-go/kvs"
	"context"
	"errors"
)

type KVSHandler struct {
}

func NewKVSHandler() *KVSHandler {
	return &KVSHandler{}
}

func (h *KVSHandler) Get(c context.Context, key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	//1. Get all replicas for this key
	//2. Use Quorum Manager to get the consistent value
	// return 
	if(key < 0 || key > 255) {
		return "", errors.New("Range of key is only [0,255]")
	}
	replicaSet := GetReplicasForKey(uint8(key))
	quorumMgr := NewQuorumMgr()
	r, err = quorumMgr.Get(context.Background(), replicaSet, key, cLevel)
	return
}
func (h *KVSHandler) Put(c context.Context,key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return nil
}
func (h *KVSHandler) GetDataFromNode(c context.Context,key int32) (r *kvs.KVData, err error) {
	data, _ := MemstoreGet(key)
	return data, nil
}
func (h *KVSHandler) PutDataInNode(c context.Context,data *kvs.KVData) (err error) {
	return
}
func (h *KVSHandler) GetHints(c context.Context,node *kvs.Node) (r []*kvs.KVData, err error) {
	return
}
