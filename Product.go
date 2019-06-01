package main

import (
	"fmt"
	"time"
)

type Product struct {
	value      int
	identifier string
}

var inputP = make(chan *Product)
var outputP = make(chan *Product)
var printP = make(chan bool)

func storeProduct() {
	var storage [productListSize]Product
	var capatity = 0
	for {
		select {
		case load := <-productChanGuard(capatity > 0, outputP):
			capatity--
			load.identifier = storage[capatity].identifier
			load.value = storage[capatity].value

		case write := <-productChanGuard(capatity < productListSize, inputP):
			storage[capatity] = *write
			capatity++
		case <-printP:
			fmt.Println(storage)

		default:
			time.Sleep(time.Second)
		}
	}
}

func productChanGuard(b bool, c chan *Product) chan *Product {
	if !b {
		return nil
	}
	return c
}
