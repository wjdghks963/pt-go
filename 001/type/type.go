package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {

	 a = 32
	 b = 31

	 fmt.Println(a)
	 // int
	 fmt.Printf("%T\n",a)

	 fmt.Println(b)
	 // main.hotdog
	 fmt.Printf("%T\n",b)

	printString() 
	printBoolean()
	printConstans()	  
	printIota()
	printBitShifter()
}