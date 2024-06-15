package main

import "fmt"

const ( 
	// untyped
	conA = 32
	conB = 32.123
	// typed
	conC string= "GOGO"
)


func printConstans(){

	fmt.Println(conA);
	fmt.Println(conB);
	fmt.Println(conC);
}