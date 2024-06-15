package main

import "fmt"


const (
	iotaA = iota
	iotaB = iota
	iotaC = iota
)


const (
	_ = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)



func printIota(){
	fmt.Println(iotaA) // 0
	fmt.Println(iotaB) // 1
	fmt.Println(iotaC) // 2


	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb )
	fmt.Printf("%d\t\t%b\n", gb, gb) 

}