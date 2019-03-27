package main

import (
	"bufio"
	"fmt"
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
	sum := 1
	for sum < SimulationTime {
		sum += 1
		buf := bufio.NewReader(os.Stdin)
		fmt.Print("> ")
		sentence,err:= buf.ReadBytes('\n')
		if err!=nil {
			fmt.Println(err)
		}else if sentence!=nil{

			fmt.Println("Products list size: ", len(ProductChannel))
			fmt.Println("Task list size : ", len(MainChannel))}



		}

}

func main() {

	if len(os.Args)>1 { //tryb gadatliwy jezeli podamy dowolny argument
	version = 1
		boss := Boss{"Krzysiu", "Ibisz"}
		go boss.newTask()
		for i := 0; i < EmployersCounter; i++ {
			employer := Employee{"Employee " + strconv.Itoa(i)}
			go pickupTask(employer)
		}
		for i := 0; i < ClientCounter; i++ {
			Client := Customer{"Client1 " + strconv.Itoa(i)}
			go Client.pickupProduct()
		}

		go printLists()
		printTime()

	} else { // tryb spokojny bez podawania argumentow

	boss := Boss{"Krzysiu", "Ibisz"}
	version = 2
		go boss.newTask()
		for i := 0; i < EmployersCounter; i++ {
			employer := Employee{"Employee " + strconv.Itoa(i)}
			go pickupTask(employer)
		}
		for i := 0; i < ClientCounter; i++ {
			Client := Customer{"Employee " + strconv.Itoa(i)}
			go Client.pickupProduct()
		}

		userInterface()
	}
}
