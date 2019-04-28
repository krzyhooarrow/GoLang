package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Employee struct {
	name            int
	status          bool // 1 cierpliwy , 2 niecierpliwy
	bufferedMachine Machine
	stat            int
}

type Map struct {
	value   int
	boolean bool
	task    Task
	e       Employee
}

func pickupTask(e Employee) {

	var read = &Task{}
	for {
		outputT <- read
		if version == 1 {
			fmt.Println("Employee"+strconv.Itoa(e.name), e.stat, e.status, "picked up the task: ", read)
		} // true = cierpliwy , false = niecierpliwy
		if read.operation == operator(e.bufferedMachine.machineType) {

			Map := Map{e.bufferedMachine.machineIdentifier, true, *read, e}
			if e.status {
				isAvaible(Map)
				//cierpliwy czeka az bedzie dostepne
			} else {

				for Map.boolean {
					// jak maszyna zajeta to idzie do innej
					Map.value = rand.Intn(MachineCounter)
					outputM <- &Map
					time.Sleep(WalkDelay * time.Second)
				}
			}
			// daje do maszyny zadanie
			Map.boolean = true

			e.stat++
			Emp2[e.name]++

			time.Sleep((MachineWorkTime + EmployeeTaskPickupDelay) * time.Second)
		} else {
			inputT <- read

		} // taski są tracone potem jeżeli nie mieszczą sie na liscie taskow
	}
}

func isAvaible(m Map) {
	for m.boolean {
		outputM <- &m
		time.Sleep(10 * time.Millisecond)

	}
}

func pickTask(employee Employee, task Task) {

	if version == 1 {
		fmt.Println("Putting task to storage", task)
	}
	Product := Product{task.score, "Employee" + strconv.Itoa(employee.name)}
	inputP <- &Product
}
