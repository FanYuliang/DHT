// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/cheekybits/genny

package server

import (
	"fmt"
	"os"
	"sync"
)

//cat generic_ccmap.go | genny gen "Key=string TransactionUnit=*blockchain.Transaction" > [targetName].go

// TransactionUnitList the set of Items
type TransactionUnitList struct {
	items []transactionUnit
	lock  sync.RWMutex
	firstReaderLoc int
}

// Set adds a new item to the tail of the list
func (d *TransactionUnitList) Append(v transactionUnit) {
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.items == nil {
		d.items = make([]transactionUnit, 1)
	}
	if v.lockType == "W" {
		s := make([]transactionUnit, 1)
		s[0] = v
		d.items = append(s, d.items...)
		d.firstReaderLoc += 1
	} else if v.lockType == "R" {
		d.items = append(d.items, v)
	} else {
		fmt.Println("LOCK TYPE is neither W nor R")
		os.Exit(7)
	}
}

// GetTransactionToCommit front
func (d *TransactionUnitList) Pop(lockType string) transactionUnit {
	d.lock.Lock()
	defer d.lock.Unlock()
	if lockType == "W" {
		if d.firstReaderLoc == 0 {
			fmt.Println("No writer available!")
			os.Exit(8)
		}
		d.firstReaderLoc -= 1
		res := d.items[0]
		d.items = d.items[1:]
		return res
	} else if lockType == "R" {
		if d.firstReaderLoc == len(d.items){
			fmt.Println("No reader available!")
			os.Exit(9)
		}
		res := d.items[len(d.items)-1]
		d.items = d.items[:len(d.items)-1]
		return res
	} else {
		res := d.items[0]
		if res.lockType == "W" {
			d.firstReaderLoc -= 1
		}
		d.items = d.items[1:]
		return res
	}
}

func (d *TransactionUnitList) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return len(d.items)
}

func (d* TransactionUnitList) Get(idx int) *transactionUnit {
	d.lock.RLock()
	defer d.lock.RUnlock()
	if idx > len(d.items){
		return nil
	}
	return &d.items[idx]
}