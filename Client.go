package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Customer struct {
	name string
}

func (C Customer) pickupProduct() {
	sum := 1
	for sum < SimulationTime {
		sum += 1

		mutex := &sync.Mutex{}
		mutex.Lock()
		x := len(ProductChannel)
		mutex.Unlock()
		if x > 0 {
			mutex.Lock()
			if version == 1 {
				fmt.Println("Client bought new product", <-ProductChannel)
			} else {
				<-ProductChannel
			}
			mutex.Unlock()
		}
		time.Sleep(time.Duration(rand.Intn(Clientinterval)) * time.Second)
	}
}
