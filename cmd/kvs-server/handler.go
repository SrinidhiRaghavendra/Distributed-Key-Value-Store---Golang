package main

import (
	"gen-go/kvs"
	"context"
)

type KVSHandler struct {
}

func NewKVSHandler() *KVSHandler {
	return &KVSHandler{}
}

func (h *KVSHandler) Get(c context.Context, key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	return
}
func (h *KVSHandler) Put(c context.Context,key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return
}
func (h *KVSHandler) GetDataFromNode(c context.Context,key int32) (r *kvs.KVData, err error) {
	return
}
func (h *KVSHandler) PutDataInNode(c context.Context,data *kvs.KVData) (err error) {
	return
}
func (h *KVSHandler) GetHints(c context.Context,node *kvs.Node) (r []*kvs.KVData, err error) {
	return
}
