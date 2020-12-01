package server
import "testing"

func testReplicaCorrect(ni []*NodeInfo, ids []NodeID, t *testing.T) {
	for i,v := range ids {
		if (v != ni[i].nodeid) {
			t.Errorf("Replica did not match %v %v", v, ni[i].nodeid)
		}
	}
}
func TestReplica(t *testing.T) {
	InitNodeInfo()
	rep := GetReplicasForKey(255)
	exp := []NodeID{3,1,2}
	testReplicaCorrect(rep, exp, t)
	rep = GetReplicasForKey(63)
	exp = []NodeID{0,1,2}
	testReplicaCorrect(rep, exp, t)
	rep = GetReplicasForKey(128)
	exp = []NodeID{2,0,1}
	testReplicaCorrect(rep, exp, t)
}
