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
	items []TransactionUnit
	lock  sync.RWMutex
	firstReaderLoc int
}

func (d * TransactionUnitList) GetItems() []TransactionUnit {
	res := make([]TransactionUnit, 0)
	for _, v := range d.items {
		res = append(res, v)
	}
	return res
}

// Set adds a new item to the tail of the list
func (d *TransactionUnitList) Append(v TransactionUnit) {
	if d.Has(v) {
		return
	}
	d.lock.Lock()
	defer d.lock.Unlock()
	if d.items == nil {
		d.items = make([]TransactionUnit, 0)
	}
	if v.lockType == "W" {
		s := make([]TransactionUnit, 1)
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

func (d *TransactionUnitList) Pop(lockType string, modified bool) TransactionUnit {
	d.lock.Lock()
	defer d.lock.Unlock()
	if lockType == "W" {
		if d.firstReaderLoc == 0 {
			fmt.Println("No writer available!")
			os.Exit(8)
		}
		res := d.items[0]
		if modified {
			d.firstReaderLoc -= 1
			d.items = d.items[1:]
		}
		return res
	} else if lockType == "R" {
		if d.firstReaderLoc == len(d.items){
			fmt.Println("No reader available!")
			os.Exit(9)
		}
		res := d.items[len(d.items)-1]
		if modified {
			d.items = d.items[:len(d.items)-1]
		}
		return res
	} else {
		res := d.items[0]
		if modified {
			if res.lockType == "W" {
				d.firstReaderLoc -= 1
			}
			d.items = d.items[1:]
		}
		return res
	}
}

func (d *TransactionUnitList) Size() int {
	d.lock.RLock()
	defer d.lock.RUnlock()
	return len(d.items)
}

func (d* TransactionUnitList) Get(idx int) *TransactionUnit {
	d.lock.RLock()
	defer d.lock.RUnlock()
	if idx > len(d.items){
		return nil
	}
	return &d.items[idx]
}


func (d *TransactionUnitList) PrintContent(title string) {
	d.lock.RLock()
	defer d.lock.RUnlock()
	fmt.Println("-----------------")
	fmt.Println(title)
	for _, v := range d.items {
		fmt.Println(v.transactionID, " ", v.lockType)
	}
	fmt.Println("=================")
}

func (d *TransactionUnitList) Remove(unit TransactionUnit) bool {
	d.lock.Lock()
	defer d.lock.Unlock()
	for i, v := range d.items {
		if v.transactionID == unit.transactionID && v.lockType == unit.lockType {
			d.items = append(d.items[:i], d.items[i+1:]...)
			if unit.lockType == "W" {
				d.firstReaderLoc -= 1
			}
			return true
		}
	}
	return false
}

func (d *TransactionUnitList) Has(unit TransactionUnit) bool {
	d.lock.RLock()
	defer d.lock.RUnlock()
	for _, v := range d.items {
		if v.transactionID == unit.transactionID && v.lockType == unit.lockType {
			return true
		}
	}
	return false
}