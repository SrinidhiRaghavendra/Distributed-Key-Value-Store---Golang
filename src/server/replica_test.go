package server

import (
	"gen-go/kvs"
	"testing"
)

func testReplicaCorrect(ni []*kvs.Node, ids []int32, t *testing.T) {
	for i, v := range ids {
		if v != ni[i].ID {
			t.Errorf("Replica did not match %v %v", v, ni[i].ID)
		}
	}
}
func TestReplica(t *testing.T) {
	InitNodeInfo(0)
	rep := GetReplicasForKey(255)
	exp := []int32{3, 1, 2}
	testReplicaCorrect(rep, exp, t)
	rep = GetReplicasForKey(63)
	exp = []int32{0, 1, 2}
	testReplicaCorrect(rep, exp, t)
	rep = GetReplicasForKey(128)
	exp = []int32{2, 0, 1}
	testReplicaCorrect(rep, exp, t)
}
