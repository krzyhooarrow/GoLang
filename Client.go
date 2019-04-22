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

	for 1<3 {

		mutex := &sync.Mutex{}
		mutex.Lock()
		x := len(ProductChannel)
		mutex.Unlock()

		if x > 0 {
			mutex.Lock()
			if version == 1 {
				p:= <-ProductChannel
				fmt.Println("Client bought new product {", p.value , p.identifier.name,"}")
			} else {
				<-ProductChannel
			}
			mutex.Unlock()
		}
		time.Sleep(time.Duration(rand.Intn(Clientinterval)) * time.Second)
	}
}
