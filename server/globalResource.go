// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package server

import (
	"fmt"
	"mp3/utils"
	"sync"
)

// cat generic_ccmap.go | genny gen "Key=string Value=*blockchain.Transaction" > [targetName].go
type CoordinatorDelegate interface{
	AddDependency(fromA string, toB string) bool
	DeleteDependency(fromA string, toB string)
	CheckDeadlock(initTransactionID string) bool
}

// StringDictionary the set of Items
type ResourceMap struct {
	items map[string]*ResourceObject //key: serverIdentifier + "_" + objectName, value: [transactionID + "_" + lockType]
	mutex sync.RWMutex
}


type TransactionUnit struct {
	transactionID 	string
	lockType		string
}

func (d *ResourceMap) Init() {
	d.items = make(map[string]*ResourceObject)
}

// Has returns true if the key exists in the ccmap
func (d *ResourceMap) Has(k string) bool {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	_, ok := d.items[k]
	return ok
}

// Get returns the value associated with the key
func (d *ResourceMap) Get(k string) *ResourceObject {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return d.items[k]
}

// Clear removes all the items from the ccmap
func (d *ResourceMap) Clear() {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	d.items = make(map[string]*ResourceObject)
}

// Size returns the amount of elements in the ccmap
func (d *ResourceMap) Size() int {
	d.mutex.RLock()
	defer d.mutex.RUnlock()
	return len(d.items)
}

func (d *ResourceMap) ConstructKey(param TryLockParam) string {
	return utils.Concatenate(*param.ServerIdentifier, "_", *param.Object)
}

func (d *ResourceMap) TryLockAt(param TryLockParam, coordinator *Coordinator) bool {
	resourceKey := d.ConstructKey(param)
	if coordinator.CheckDeadlock(param) {
		return false
	}

	if d.items[resourceKey] == nil {
		d.items[resourceKey] = new(ResourceObject)
		d.items[resourceKey].Init()
	}
	if *param.LockType == "R" && d.items[resourceKey].lockHolders.Has(TransactionUnit{transactionID:*param.TransactionID, lockType:"W"}) {
		//if lock holder is W, then for the read lock request with same transaction id, give it directly
		return true
	}

	//first put into waiting queue
	if *param.LockType == "W" && d.items[resourceKey].lockHolders.Has(TransactionUnit{transactionID:*param.TransactionID, lockType:"R"}) {
		d.items[resourceKey].UnlockHolder(TransactionUnit{*param.TransactionID, "R"})
		d.items[resourceKey].AppendToUpgradeList(TransactionUnit{*param.TransactionID, *param.LockType})
	} else {
		d.items[resourceKey].AppendToWaitingQueue(TransactionUnit{*param.TransactionID, *param.LockType})
	}

	if d.Has(resourceKey) {
		for _, holder := range d.Get(resourceKey).lockHolders.GetItems() {
			if !(holder.lockType == "R" && *param.LockType == "R") {
				coordinator.transactionDependency.Set(*param.TransactionID, holder.transactionID)
			}
		}
	}
	d.items[resourceKey].PrintContent()
	d.items[resourceKey].mutex.Lock()
	fmt.Println("*param.TransactionID:", *param.TransactionID)
	fmt.Println("d.items[resourceKey].GetNextTarget()", d.items[resourceKey].GetNextTarget(false))
	for *param.TransactionID !=  d.items[resourceKey].GetNextTarget(false).transactionID {
		d.items[resourceKey].cond.Wait()
	}
	currUnit := d.items[resourceKey].GetNextTarget(true)
	if !d.items[resourceKey].abortList.Remove(TransactionUnit{*param.TransactionID,*param.LockType}) {
		d.items[resourceKey].lockHolders.Append(currUnit)
	}
	d.items[resourceKey].mutex.Unlock()

	coordinator.transactionDependency.Delete(*param.TransactionID)
	return true
}

func(d *ResourceMap) AbortAllRelatedTo(transactionID string) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	for k, _ := range d.items {
		wFlag := false
		rFlag := false
		if d.items[k].upgradeList.Remove(TransactionUnit{transactionID:transactionID, lockType:"W"}) {
			wFlag = true
		}
		if d.items[k].waitingQueue.Remove(TransactionUnit{transactionID:transactionID, lockType:"W"}) {
			wFlag = true
		}
		if d.items[k].waitingQueue.Remove(TransactionUnit{transactionID:transactionID, lockType:"R"}) {
			rFlag = true
		}
		if wFlag {
			d.items[k].AppendToAbortList(TransactionUnit{transactionID:transactionID, lockType:"W"})
		}
		if rFlag {
			d.items[k].AppendToAbortList(TransactionUnit{transactionID:transactionID, lockType:"R"})
		}
		fmt.Println("HERERER!")
		d.items[k].cond.Broadcast()
	}

}