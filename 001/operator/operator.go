package main

import "fmt"

// 전역 변수 선언
var y = 100 + 24

// int 타입 z 전역 변수 선언 => default = 0
var z int

func main() {

	// 단축 변수 선언
	x := 42
	fmt.Println(x)

	// 값 대입
	x = 99
	fmt.Println(x)

	// 전역 스코프에 있는 y를 print
	fmt.Println(y)
	fmt.Printf("%T\n",y)

	z := "JJJJ"
	fmt.Println(z)
	fmt.Printf("%T\n",z)

}
