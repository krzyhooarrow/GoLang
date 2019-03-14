package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Employee struct {
	name string
}

func (e Employee) executeTask(t Task) int {
	if t.operation == "+" {
		return t.arg2 + t.arg1
	}
	if t.operation == "-" {
		return t.arg2 - t.arg1
	}
	if t.operation == "/" {
		if t.arg1 != 0 {
			return t.arg2 / t.arg1
		}
	}
	if t.operation == "*" {
		return t.arg2 * t.arg1
	}
	return 0
}

func pickupTask(e Employee) {
	sum := 1
	for sum < SimulationTime {
		sum += 1


		mutex := &sync.Mutex{}
		mutex.Lock()
		x:= len(taskList)


		if x > 0 && len(productList) < productListSize{


				fmt.Println("I picked up the task: ", e)
				Product := Product{e.executeTask(taskList[0]), e}
				productList = append(productList, Product)
				taskList = taskList[:copy(taskList[0:], taskList[1:])]
				mutex.Unlock()

		}else {
		mutex.Unlock()}

		time.Sleep(time.Duration(rand.Intn(EmployeeTaskPickupDelay)) * time.Second)
	}
}
