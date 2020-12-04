package server

import (
	"context"
	"gen-go/kvs"
	"log"
)

func Recover() {
	WalBegin()
	for {
		v, s := WalRead()
		if s == 0 {
			break
		}
		MemstorePut(v)
	}

	nodes := GetEveryone()
	for i, v := range nodes {
		if i != int(me) {
			com := NewIntraNodeComm(&v)
			var ctx context.Context
			h, err := com.GetHints(ctx, &kvs.Node{ID: me})
			if err != nil {
				log.Printf("Recovery warning: Error fetching hints from %v (%v)\n", v.ID, err)
			} else {
				for _, data := range h {
					WalPut(*data)
					MemstorePut(data)
				}
			}
		}
	}
}
