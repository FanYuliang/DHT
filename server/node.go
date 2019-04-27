package server

import (
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"context"
	"mp3/utils"
	"sync"
)

// Node can be embedded to have forward compatible implementations.
type Node struct {
	Name                string
	CoordinatorDelegate CoordinatorClient
	data                map[string]string
	lockMap  			map[string]*LockTuple
	lockMapLock			sync.RWMutex
}

type LockTuple struct {
	mutex         sync.RWMutex
	transactionID string
}
// NodeServer is the server API for Node service.
type NodeServer interface {
	ClientSet(context.Context, *SetParams) (*Feedback, error)
	ClientGet(context.Context, *GetParams) (*Feedback, error)
	CommitTransaction(context.Context, *Transaction) (*Feedback, error)
	AbortTransaction(context.Context, *Transaction) (*Feedback, error)
}

func (n *Node) Init() {
	n.lockMap = make(map[string]*LockTuple)
	n.data = make(map[string]string)
}

func (n *Node) ClientSet(ctx context.Context, req *SetParams) (*Feedback, error) {
	if *req.ServerIdentifier != n.Name {
		return nil, status.Error(codes.InvalidArgument, "Called the wrong node server")
	}
	//TODO: add WLOCK
	n.lockMapLock.Lock()
	_, ok := n.lockMap[*req.ObjectName]
	if !ok {
		n.lockMap[*req.ObjectName] = &LockTuple{
			mutex:         sync.RWMutex{},
			transactionID: *req.TransactionID,
		}
	} else {
		n.lockMap[*req.ObjectName].transactionID = *req.TransactionID
	}
	n.lockMapLock.Unlock()

	n.WLock(*req.ObjectName, *req.TransactionID)
	defer n.WUnLock(*req.ObjectName, *req.TransactionID)

	n.data[*req.ObjectName] = *req.Value
	resFeedback := &Feedback{}
	result := "OK"
	resFeedback.Message = &result
	return resFeedback, nil
}

func (n *Node) ClientGet(ctx context.Context, req *GetParams) (*Feedback, error) {
	if *req.ServerIdentifier != n.Name {
		return nil, status.Error(codes.InvalidArgument, "Called the wrong node server")
	}
	n.lockMapLock.RLock()
	_, ok := n.lockMap[*req.ObjectName]
	if !ok {
		resFeedback := &Feedback{}
		var result string
		result = "NOT FOUND"
		resFeedback.Message = &result
		//TODO: tell coordinator to abort the current transaction

		n.lockMapLock.RUnlock()
		return resFeedback, status.Error(codes.Aborted, "not found")
	}
	n.lockMapLock.RUnlock()


	n.RLock(*req.ObjectName, *req.TransactionID)
	defer n.RUnLock(*req.ObjectName, *req.TransactionID)

	resFeedback := &Feedback{}
	val, ok := n.data[*req.ObjectName]
	resFeedback.Message = &val
	return resFeedback, nil
}


func (*Node) CommitTransaction(ctx context.Context, req *Transaction) (*Feedback, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CommitTransaction not implemented")
}
func (n *Node) AbortTransaction(ctx context.Context, req *Transaction) (*Feedback, error) {
	n.abortTransaction(*req.Id)
	return nil, status.Errorf(codes.Unimplemented, "method AbortTransaction not implemented")
}

func (*Node) abortTransaction(transactionID string) {

}

func (n *Node) RLock(objectName string, transactionID string) {
	fmt.Println("RLock on ", objectName)
	tryLockParam := TryLockParam{}
	tryLockParam.TransactionID = &transactionID
	tryLockParam.ServerIdentifier = &n.Name
	tryLockParam.Object = &objectName
	lockType := "R"
	tryLockParam.LockType = &lockType
	feedback, err := n.CoordinatorDelegate.TryLock(context.Background(), &tryLockParam)
	utils.CheckError(err)
	if *feedback.Message == "Abort" {
		n.abortTransaction(transactionID)
	}
	n.lockMap[objectName].mutex.RLock()
}

func (n *Node) RUnLock(objectName string, transactionID string) {
	fmt.Println("RUnLock on ", objectName)
	reportUnlockParam := ReportUnLockParam{}
	reportUnlockParam.Object = &objectName
	reportUnlockParam.ServerIdentifier = &n.Name
	reportUnlockParam.TransactionID = &transactionID
	_, err := n.CoordinatorDelegate.ReportUnlock(context.Background(), &reportUnlockParam)
	utils.CheckError(err)
	n.lockMap[objectName].mutex.RUnlock()
}

func (n *Node) WLock(objectName string, transactionID string) {
	fmt.Println("WLock on ", objectName)
	tryLockParam := TryLockParam{}
	tryLockParam.TransactionID = &transactionID
	tryLockParam.ServerIdentifier = &n.Name
	lockType := "W"
	tryLockParam.LockType = &lockType
	tryLockParam.Object = &objectName
	feedback, err := n.CoordinatorDelegate.TryLock(context.Background(), &tryLockParam)
	utils.CheckError(err)
	if *feedback.Message == "Abort" {
		n.abortTransaction(transactionID)
	}
	n.lockMap[objectName].mutex.Lock()
}

func (n *Node) WUnLock(objectName string, transactionID string) {
	fmt.Println("WUnLock on ", objectName)
	reportUnlockParam := ReportUnLockParam{}
	reportUnlockParam.Object = &objectName
	reportUnlockParam.ServerIdentifier = &n.Name
	reportUnlockParam.TransactionID = &transactionID
	_, err := n.CoordinatorDelegate.ReportUnlock(context.Background(), &reportUnlockParam)
	utils.CheckError(err)
	n.lockMap[objectName].mutex.Unlock()
}
