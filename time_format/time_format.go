package main

import (
	"fmt"
	"time"
)

func doSomething() {
	for i := 0; i < 100; i++ {
		// Do nothing
		fmt.Println(i)
	}
	fmt.Println("OK")
}

func main() {

	beforeTime := time.Now()
	doSomething()
	afterTime := time.Now()

	fmt.Println(afterTime.Sub(beforeTime))
}
