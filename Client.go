package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Customer struct {
	name string
}

func (C Customer) pickupProduct() {
	var Product = &Product{}

	for  {
		time.Sleep(time.Duration(rand.Intn(ClientInterval)) * time.Second)
			if version == 1 {
				outputP<-Product
				fmt.Println("Client bought new product", Product)
			} else {
				outputP<-Product
			}

	}
}
