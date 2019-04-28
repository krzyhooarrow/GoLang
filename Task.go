package main

import (
	"fmt"
	"time"
)

type Task struct {
	arg1      int
	arg2      int
	operation string
	score     int
}

var inputT = make(chan *Task)
var outputT = make(chan *Task)
var printT = make(chan bool)

func storeTask() {
	var storage [taskListSize]Task
	var capatity = 0
	for {
		select {
		case load := <-taskChanGuard(capatity > 0, outputT):

			capatity--
			load.arg1 = storage[capatity].arg1
			load.arg2 = storage[capatity].arg2
			load.operation = storage[capatity].operation

		case write := <-taskChanGuard(capatity < taskListSize, inputT):
			storage[capatity] = *write
			capatity++
		case <-printT:
			fmt.Println(storage)

		default:
			time.Sleep(time.Second)
		}
	}
}

func taskChanGuard(b bool, c chan *Task) chan *Task {
	if !b {
		return nil
	}
	return c
}

func createTask(arg1 int, arg2 int, operation string) Task {
	return Task{arg1, arg2, operation, 0}
}

func operator(i int) (operator string) {
	if i == 0 {
		return "+"
	}
	if i == 1 {
		return "*"
	}
	if i == 2 {
		return "/"
	}
	if i == 3 {
		return "-"
	}
	return "Bad Operator"
}
