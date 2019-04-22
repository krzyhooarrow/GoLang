package main

import (
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)


var MainChannel = make(chan Task,taskListSize)

type Boss struct {
	name string
	surname string
}


func (b Boss)newTask() {
	sum := 1
	for sum < SimulationTime {


		if len(MainChannel) < taskListSize {


			task := createTask(rand.Intn(TaskArgRange), rand.Intn(TaskArgRange), operator(rand.Intn(4)))
			atomic.AddUint64(&counter, 1)
			//if version ==1 {
			//
			//	fmt.Println("The Boss have a new task!")
			//}
			mutex := &sync.Mutex{}
			mutex.Lock()
			MainChannel<-task
			mutex.Unlock()
		}
		time.Sleep(time.Duration(rand.Intn(NewTaskTime)) * time.Second)
	}
}
