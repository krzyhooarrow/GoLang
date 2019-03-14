package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Customer struct
{
	name string

}

func (C Customer)pickupProduct(){

	mutex:=&sync.Mutex{}
	mutex.Lock()
	x:= len(productList)
	mutex.Unlock()
	if x>0 {
		mutex.Lock()
		productList = productList[:copy(taskList[0:], taskList[1:])]
		mutex.Unlock()
		fmt.Println("Client bought new product")

	}
	time.Sleep(time.Duration(rand.Intn(Clientinterval)) * time.Second)
}