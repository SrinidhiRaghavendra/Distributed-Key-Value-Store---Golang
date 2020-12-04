package server

import (
	"gen-go/kvs"
	"context"
)

type QuorumMgr struct {

}

func NewQuorumMgr() *QuorumMgr {
	return &QuorumMgr{}
}

/**
TODO: Implement Two/Three Phase commit later
*/
func (q *QuorumMgr) Put(c context.Context, replicaSet []*kvs.Node, data *kvs.KVData, cLevel kvs.ConsistencyLevel) *kvs.SystemException {
	var successCount int
	var intraNodeComms []*IntraNodeComm
	for _, node := range replicaSet {
		nodeComms := NewIntraNodeComm(node)
		err := nodeComms.Setuprep()
		if(err == nil) {
			successCount += 1
		}
		intraNodeComms = append(intraNodeComms, nodeComms)
	}
	if(successCount >= int(cLevel)) {
		for _, nodeComms := range intraNodeComms {
			err := nodeComms.PutDataInNode(c, data)
			if(err != nil) {
				//Store hint for this failed node 
				StoreHint(nodeComms.n.ID, *data)
			}
		}
		return nil
	} else {
		execption := kvs.NewSystemException()
		*execption.Message = "Consistency Level required is not met. Hence th eoperation is a failure. Please try again!"
		return execption
	}
}

func (q *QuorumMgr) Get(c context.Context, replicaSet []*kvs.Node, key int32, cLevel kvs.ConsistencyLevel) (string, *kvs.SystemException) {
	var latestData kvs.KVData 
	var latestCount int
	for _, node := range replicaSet {
		nodeComms := NewIntraNodeComm(node)
		kvData, err := nodeComms.GetDataFromNode(c, key)
		if(err != nil) {
			if(latestData.Value == kvData.Value) {
				latestCount += 1
				continue
			}
			if(latestData.Timestamp < kvData.Timestamp) {
				latestData = *kvData
				latestCount = 1
			}
		}
	}
	if(latestCount >= int(cLevel)) {
		return latestData.Value, nil
	} else {
		execption := kvs.NewSystemException()
		*execption.Message = "Consistency Level required is not met. Hence th eoperation is a failure. Please try again!"
		return "", execption
	}
}
