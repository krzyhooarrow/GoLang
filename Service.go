package main

import (
	"fmt"
	"time"

	//"time"
)

type Service struct {
	name string
}

var receiveMessage = make(chan int)
var returnEmployee = make(chan *ServiceEmployee)
var printEmployers = make(chan bool)

func (s Service) work() {
	for ; ; {

		select {
		case in := <-receiveMessage:
			{
				go fixMachine(<-returnEmployee, in)
			}
		case <-printEmployers:
			{
				for i:=0;i<ServiceEmployeeCounter ;i++  {
					emp:=<-returnEmployee
					fmt.Println(emp)
					returnEmployee<-emp
				}
			}
		}

	}
}

func fixMachine(employee *ServiceEmployee, machineID int) {
	if version == 1 {

		fmt.Println("The service send ", employee, " to fix the Machine ", machineID)
	}
	time.Sleep((ServiceEmployeeWalkDelay) * time.Second)

	backdoor <- machineID
	if version == 1 {
		fmt.Println("Work was done. ", employee, " is back!")
	}
	returnEmployee <- employee // wraca

}

//
type ServiceEmployee struct {
	name string
}
