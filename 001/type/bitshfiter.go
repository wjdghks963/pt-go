package main

import "fmt"

func printBitShifter(){
	x := 4
	fmt.Printf("%d\t\t%b\n",x,x) // 4 100


	y := x << 1 // 왼쪽으로 한칸 밈
	fmt.Printf("%d\t\t%b\n",y,y) // 8 1000


	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb )
	fmt.Printf("%d\t\t%b\n", gb, gb) 

}

