package server

import (
	"gen-go/kvs"
	"container/list"
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
	var failedNodes List = list.New()
	var successCount int
	var intraNodeComms [len(replicaSet)]IntraNodeComm
	for index, node := range replicaSet {
		nodeComms := NewIntraNodeComm(node)
		err := nodeComms.Setuprep()
		if(err == nil) {
			successCount += 1
		}
		intraNodeComms[index] = nodeComms
	}
	if(successCount >= cLevel) {
		for index, nodeComms := range intraNodeComms {
			err := nodeComms.PutDataInNode(c, data)
			if(err != nil) {
				//Store hint for this failed node 
				StoreHint(node.GetID(), data)
			}
		}
		return nil
	} else {
		execption := kvs.NewSystemException()
		execption.Message = "Consistency Level required is not met. Hence th eoperation is a failure. Please try again!"
		return execption
	}
}

func (q *QuorumMgr) Get(c context.Context, replicaSet []*kvs.Node, int32 key, cLevel kvs.ConsistencyLevel) (string, *kvs.SystemException) {
	var latestData kvs.KVData 
	var latestCount int
	successCount := 0
	for index, node := range replicaSet {
		nodeComms := NewIntraNodeComm(node)
		kvData, err := nodeComms.GetDataFromNode(c, key)
		if(err != nil) {
			if(latestData.Value == kvData.Value) {
				latestCount += 1
				continue
			}
			if(latestData.TimeStamp < kvData.TimeStamp) {
				latestData = *kvData
				latestCount = 1
			}
		}
	}
	if(latestCount >= cLevel) {
		return (latestData.Value, nil)
	} else {
		execption := kvs.NewSystemException()
		execption.Message = "Consistency Level required is not met. Hence th eoperation is a failure. Please try again!"
		return (nil, execption)
	}
}
