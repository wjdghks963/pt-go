// mutable

// 바이트 슬라이스로 이루어짐

package main

import "fmt"


func printString() {

	s := "Hello, World";
	fmt.Println(s) // Hello, World
	fmt.Printf("%T\n",s)  // string
	
	
	bs := []byte(s)
	fmt.Println(bs) // [72 101 108 108 111 44 32 87 111 114 108 100]
	fmt.Printf("%T\n",bs)  // []unit8
}