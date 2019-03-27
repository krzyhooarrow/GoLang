package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var ProductChannel = make(chan Product,50)

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
		x:= len(MainChannel)


		if x > 0 && len(ProductChannel) < productListSize{

			if version==1 {

				fmt.Println("I picked up the task: ", e)
			}
				Product := Product{e.executeTask(<-MainChannel), e}
				ProductChannel<-Product

				mutex.Unlock()

		}else {
		mutex.Unlock()}

		time.Sleep(time.Duration(rand.Intn(EmployeeTaskPickupDelay)) * time.Second)
	}
}
