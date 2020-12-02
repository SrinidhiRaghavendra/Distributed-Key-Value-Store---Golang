package server

import "gen-go/kvs"

func recover(w *wal) {
	nodes := GetEveryone()
	for i, v := range nodes {
		if i != int(me) {
			com := NewIntraNodeComm(&v)
			com.GetHints(&kvs.Node{})
		}
	}
}
