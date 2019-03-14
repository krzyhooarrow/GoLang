package main

import (
	"fmt"
	"sync"
	"time"
)

type Debugger struct {
	name string
}
func debugTask(e Debugger) {
	sum := 1
	for sum < SimulationTime {
		sum += 1

		mutex := &sync.Mutex{}
		mutex.Lock()
		x:=len(taskList)
		mutex.Unlock()
	if x>0 {


		mutex := &sync.Mutex{}
		mutex.Lock()
		taskList = taskList[:copy(taskList[0:], taskList[1:])]
		mutex.Unlock()
		fmt.Println("picking task")
	}
	time.Sleep(2 * time.Second)
}

	}