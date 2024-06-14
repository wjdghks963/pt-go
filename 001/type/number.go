// unit(8~64) -> 0~2^(8~63)

// int(8~64) -> -2^(7~63) ~ 2^(7~63)

// float(32 or 64)

package main

import "fmt"


func number() {

	var x int
	var y float64
	
	x = 32
	y = 32.12312

	fmt.Println(x)
	fmt.Println(y)

	// x = 32.123 -> error
}