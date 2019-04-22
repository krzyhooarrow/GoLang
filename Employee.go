package main

import (
	"math/rand"
	"sync"
	"time"
)

var counterEmp = 0
var ProductChannel = make(chan Product, productListSize)
var EmployersStatistics [EmployersCounter] int
var EmployersList [EmployersCounter] Employee

type Employee struct {
	name            string
	identifier      int
	status          string
	bufferedMachine Machine
}

func createEmployee(name string) Employee {
	e := Employee{}
	e.name = name
	e.status = chooseStatus()
	e.identifier = counterEmp
	EmployersStatistics[counterEmp] = 0
	counterEmp++
	e.bufferedMachine = MachineList[rand.Intn(machineCount)]
	EmployersList[counterEmp-1] = e
	return e
}

func chooseStatus() string {
	if rand.Intn(2) == 1 {
		return "patient"

	} else {
		return "inpatient"
	}
}
func changeMachine() Machine {
	return MachineList[rand.Intn(machineCount)]
}

func (e Employee) putTask(t Task) {
	for 1 < 2 {

		if e.status == "inpatient" {
			e.bufferedMachine = changeMachine()
			if e.bufferedMachine.status == false {
				time.Sleep(time.Duration(walkInterval) * time.Second)
				e.bufferedMachine.work(t, e)
				break
			} else {
				time.Sleep(time.Duration(1) * time.Second)
			}
		} else {
			if e.bufferedMachine.status == false {
				e.bufferedMachine.work(t, e)
				break
			} else {
				time.Sleep(time.Duration(1) * time.Second)
			}
		}
	}
}
func (e Employee) getProduct(t Task) {
	ProductChannel <- Product{t.score, e}
}

func pickupTask(e Employee) {

	sum := 1
	for sum < SimulationTime {
		mutex := &sync.Mutex{}
		mutex.Lock()
		x := len(MainChannel)

		if x > 0 && len(ProductChannel) < productListSize {

			e.putTask(<-MainChannel)
		}
		mutex.Unlock()

		time.Sleep(time.Duration(rand.Intn(EmployeeTaskPickupDelay)) * time.Second)
	}
}
