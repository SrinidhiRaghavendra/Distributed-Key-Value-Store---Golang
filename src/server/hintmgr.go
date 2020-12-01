package server

var hints [4][]KVData

func StoreHint(id NodeID, hint KVData) {
	hints[id] = append(hints[id], hint)
}

func GetHintsForNode(id NodeID) []KVData {
	return hints[id]
}
