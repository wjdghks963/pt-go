package ainterface

import "fmt"


type expense interface {
	cost() float64
}

type printer interface {
	print()
}


func (e email) cost() float64 {
	if !e.isSub{
		return 0.05 * float64(len(e.body))
	}
	return 0.01 * float64(len(e.body))
}

func (e email) print() {
	fmt.Println(e.body)
}

type email struct {
	isSub bool
	body string
}


func testMultiInter(e expense, p printer){
	fmt.Printf("Printing cost : $%.2f ..\n", e.cost)
	p.print()
	fmt.Println("================")
}