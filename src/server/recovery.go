package server

import (
	"context"
	"gen-go/kvs"
)

func recover(w *wal) {
	nodes := GetEveryone()
	for i, v := range nodes {
		if i != int(me) {
			com := NewIntraNodeComm(&v)
			var ctx context.Context
			com.GetHints(ctx, &kvs.Node{})
		}
	}
}
