package main

import (
	"fmt"
	"time"
)

func regularFunction() {
	fmt.Print("Just entered regularFunction()")
	time.Sleep(5 * time.Second)
}

func goroutineFunction() {
	fmt.Println("Just enetered goroutineFunction()")
	time.Sleep(10 * time.Second)
	fmt.Println("goroutineFunction finished its work")
}

func main() {
	go goroutineFunction()
	fmt.Println("In main one lin ebelow goroutineFunction()")
	regularFunction()
	fmt.Println("IN main one lin ebelow regularFunction()")
}
