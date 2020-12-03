package server

import "gen-go/kvs"

var hints [4][]kvs.KVData

func StoreHint(id int32, hint kvs.KVData) {
	hints[id] = append(hints[id], hint)
}

func GetHintsForNode(id int32) []kvs.KVData {
	return hints[id]
}
