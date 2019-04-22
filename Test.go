package main

import (
	"fmt"
	"math/rand"
	"os"

	"strconv"
	"time"
)

var counter uint64

func printTime() {
	for i := 0; i < SimulationTime; i++ {
		fmt.Println("Seconds passed: ", i)
		time.Sleep(time.Second)
	}
}

func printLists() {
	sum := 1
	for sum < SimulationTime {
		sum += 1
		fmt.Println("Products list size: ", len(ProductChannel))
		fmt.Println("Task list size : ", len(MainChannel))
		time.Sleep(5 * time.Second)
	}
}

func userInterface() {
	var age int
	for 1< 2 {

		fmt.Print("> ")
		fmt.Scan(&age)

			if age<EmployersCounter {
				fmt.Println(EmployersList[age], EmployersStatistics[age])
			}
			fmt.Println("Products list size: ", len(ProductChannel))
			fmt.Println("Task list size : ", len(MainChannel))
		}
	}


func loudMode() {
	go printLists()
	printTime()
}

func quietMode() {
	userInterface()
}

func main() {

	boss := Boss{"Krzysiu", "Ibisz"}
	go boss.newTask()
	for i := 0; i < machineCount; i++ {
		if rand.Intn(2) == 1 {
			mach := Machine{i, "+", rand.Intn(avgMachineWorkTime), false}
			MachineList[i] = mach
		} else {
			mach := Machine{i, "*", rand.Intn(avgMachineWorkTime), false}
			MachineList[i] = mach
		}

	}
	for i := 0; i < EmployersCounter; i++ {
		employer := createEmployee("Employee" + strconv.Itoa(i))
		go pickupTask(employer)
	} // pracownik dostaje randomowo przydzielona maszyne

	for i := 0; i < ClientCounter; i++ {
		Client := Customer{"Client" + strconv.Itoa(i)}
		go Client.pickupProduct()
	}



	if len(os.Args) > 1 { //tryb gadatliwy jezeli podamy dowolny argument
		version = 1
		loudMode()
	} else { // tryb spokojny bez podawania argumentow
		version = 2
		quietMode()
	}

}
