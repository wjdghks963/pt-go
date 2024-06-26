package main

import (
	"fmt"
	"time"
)


func sendEmail(message string)  {
	
   	go func ()  {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received : '%s'\n", message)
	}()

fmt.Printf("Email received : '%s'\n", message)

}

func coTest(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func coMain() {
	coTest("Hello there Stacy!")
	coTest("Hi there John!")
	coTest("Hey there Jane!")
}