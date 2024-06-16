package astruct

import "fmt"

func makeBaseStruct(){


	// 익명 구조체 -> 인스턴스 바로 생성
	// myCar := struct {
	// 	Make string
	// 	Model string
	// } {
	// 	Make: "tesla",
	// 	Model: "model 3",
	// }


	type car struct {
		make string
		model string
	}
	
	type truck struct {
		// car의 구조체를 상속 받아 사용합니다.
		car
		bdeSize int
	}


	myTruck := truck{
		bdeSize: 10,
		car: car{
			make: "kia",
			model: "bigModel",
		},
	}

	fmt.Printf("myTruck Name :: %v\n", myTruck.model) // bigModel

}