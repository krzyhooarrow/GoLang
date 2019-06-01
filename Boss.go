package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Boss struct {
	name    string
	surname string
}

func (b Boss) newTask() {

	for {

		task := createTask(rand.Intn(TaskArgRange), rand.Intn(TaskArgRange), operator(rand.Intn(2)))

		if version == 1 {
			fmt.Println("The Boss have a new task!", task)
		}

		inputT <- &task
		time.Sleep(NewTaskTime * time.Second)
	}
}
