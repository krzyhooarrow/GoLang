package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)
var Emp[EmployersCounter] Employee
var Emp2[EmployersCounter] int

func printTime() {
	var counter = 0
	for {
		fmt.Println("Seconds passed: ", counter)
		time.Sleep(time.Second)
		counter++
	}
}

func userInterface() {
	var in int
	for {


		fmt.Scan(&in)
		if in<EmployersCounter {
		fmt.Println("Employee",Emp[in].name,Emp[in].stat,Emp2[in])

		}else if in == 20 {
			printT<-true
			printM<-true
			printP<-true
		}

	}

}


func mainEvent() {

	go storeProduct()
	go storeTask()
	go storeMachine()

	boss := Boss{"Krzysiu", "Ibisz"}
	go boss.newTask()

	for i := 0; i < MachineCounter; i++ {
		inputM <- &Machine{i, rand.Intn(2), false}
	}

	for i := 0; i < EmployersCounter; i++ {
		employer := Employee{i, rand.Intn(2) != 0, storage[rand.Intn(MachineCounter)],0}
		Emp[i] = employer
		go pickupTask(employer)
	}

	for i := 0; i < ClientCounter; i++ {
		Client := Customer{"Client" + strconv.Itoa(i)}
		go Client.pickupProduct()
	}
}

func main() {

	if len(os.Args) > 1 { //tryb gadatliwy jezeli podamy dowolny argument
		version = 1
		mainEvent()

		printTime()

	} else { // tryb spokojny bez podawania argumentow
		version = 2
		mainEvent()

		userInterface()
	}





}
