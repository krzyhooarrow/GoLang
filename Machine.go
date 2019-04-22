package main

import (
	"fmt"
	"sync"
	"time"
)

var MachineList [machineCount] Machine

type Machine struct {
	machineIdentifier int
	machineType       string // 1 = adder 2 = multi
	machineSpeed      int
	status            bool
}

func (m Machine) work(t Task, e Employee) {
	if m.machineType == t.operation {

		mutex := sync.Mutex{}
		mutex.Lock()
		m.status = true
		mutex.Unlock()
		if version == 1 {

		fmt.Println(e.status,e.name,"{",EmployersStatistics[e.identifier],"} used machine",m.machineIdentifier,"ZZZZ... working...ZZZZ ... ")}
		time.Sleep(time.Duration(m.machineSpeed) * time.Second)
		EmployersStatistics[e.identifier]++
		m.executeTask(t, e)
		mutex.Lock()
		m.status = false
		mutex.Unlock()
	}	// jak nie to porzuca zadanie
}


func (m Machine) executeTask(t Task, e Employee) {

	if t.operation == "+" {
		t.score = t.arg2 + t.arg1
	}
	if t.operation == "-" {
		t.score = t.arg2 - t.arg1
	}
	if t.operation == "/" {
		if t.arg1 != 0 {
			t.score = t.arg2 / t.arg1
		}
	}
	if t.operation == "*" {
		t.score = t.arg2 * t.arg1
	}
	e.getProduct(t)
}
