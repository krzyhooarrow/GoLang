package main

import (
	"fmt"
	"strconv"
	"time"
)

var counter uint64
var taskList []Task
var productList []Product

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

		fmt.Println(productList)
			fmt.Println(taskList)
		fmt.Println("Products list size: ", len(productList))
		fmt.Println("Task list size : ", len(taskList))
		time.Sleep(2 * time.Second)

	}

}

func main() {

	go Boss()

	for i := 0; i < EmployersCounter; i++ {
		employer := Employee{"Employee " + strconv.Itoa(i)}
		go pickupTask(employer)
		//debuger := Debugger{strconv.Itoa(i)}
	//	go debugTask(debuger)
	}

	//go printLists()
	printTime()

}
