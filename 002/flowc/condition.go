package flowc

import "fmt"


func condition() {

	if true {
		fmt.Println("Hello WORLD")
	}

	// 조건값의 기본값은 true 이기 떄문에 case가 true 인것을 검색
	switch  {
    case false:
            fmt.Println("false")
    case 2==4:
            fmt.Println("false")

    case 2==2:
            fmt.Println("true1")

    case 3==3:
            fmt.Println("true2")

    default:
            fmt.Println("default")
}

}