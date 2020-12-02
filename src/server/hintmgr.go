package server

import "gen-go/kvs"

var hints [4][]kvs.KVData

func StoreHint(id NodeID, hint kvs.KVData) {
	hints[id] = append(hints[id], hint)
}

func GetHintsForNode(id NodeID) []kvs.KVData {
	return hints[id]
}
