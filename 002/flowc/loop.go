package flowc

import "fmt"


func printLoop () {

	for i:=0; i<=10; i++ {
		for j:=0; j<=10; j++ {
			fmt.Printf("OUT :: %d\t INNER :: %d\n" , i, j)
		}
	}

	x := 1
	for x <10 {
    	fmt.Println(x)
	}


	y := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(y)
		y++
	}


	for i := 33; i<=122; i++{
		fmt.Printf("%v\t%#x\t%#U\n",i,i,i)
	}
}