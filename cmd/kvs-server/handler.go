package main

import (
	"gen-go/kvs"
)

type KVSHandler struct {
}

func NewKVSHandler() *KVSHandler {
	return &KVSHandler{}
}

func (h *KVSHandler) Get(key int32, cLevel kvs.ConsistencyLevel) (r string, err error) {
	return
}
func (h *KVSHandler) Put(key int32, value string, cLevel kvs.ConsistencyLevel) (err error) {
	return
}
func (h *KVSHandler) GetDataFromNode(key int32) (r *kvs.KVData, err error) {
	return
}
func (h *KVSHandler) PutDataInNode(data *kvs.KVData) (err error) {
	return
}
func (h *KVSHandler) GetHints(node *kvs.Node) (r []*kvs.KVData, err error) {
	return
}
