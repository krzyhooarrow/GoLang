package main

import (
	"fmt"
	"time"
)

type Machine struct {
	machineIdentifier int
	machineType       int // 1 = adder 2 = multi
	machineStatus     bool
}

var inputM = make(chan *Machine)
var outputM = make(chan *Map)
var printM = make(chan bool)


var storage [MachineCounter]Machine
func storeMachine() {
	var capatity = 0
	for {
		select {
		case in := <-machineChanGuard(capatity < MachineCounter, inputM):
			storage[capatity] = *in
			capatity++

		case out := <-outputM:
			if storage[out.value].machineStatus {
				out.boolean = true
			} else {

				storage[out.value].machineStatus = true
				out.boolean = false
				go func() {
					out.task.score=calculateValue(out.task)
					time.Sleep(MachineWorkTime * time.Second)
					pickTask(out.e,out.task)
					storage[out.value].machineStatus = false
				}()
			}
		case <-printM:
			fmt.Println(storage)
		}
	}
}


func calculateValue(t Task) int {
	if t.operation == "+" {
		return  t.arg2 + t.arg1
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

func machineChanGuard(b bool, c chan *Machine) chan *Machine {
	if !b {
		return nil
	}
	return c
}
