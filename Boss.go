package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func Boss() {
	sum := 1
	for sum < SimulationTime {
		sum += 1

		if len(taskList) < taskListSize {



			task := createTask(rand.Intn(TaskArgRange), rand.Intn(TaskArgRange), operator(rand.Intn(4)))
			atomic.AddUint64(&counter, 1)
			fmt.Println("The Boss have a new task!")

			mutex := &sync.Mutex{}
			mutex.Lock()

			taskList = append(taskList, task)

			mutex.Unlock()
		}
		time.Sleep(time.Duration(rand.Intn(NewTaskTime)) * time.Second)
	}
}
