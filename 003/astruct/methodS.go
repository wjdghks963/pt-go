package astruct

import "fmt"

// rect 구조체 정의
type rect struct {
	width  int
	height int
}

// rect 구조체의 메서드 정의
func (r rect) area() int {
	return r.width * r.height
}

func MakeMetStruct() {
	// rect 구조체의 인스턴스 생성
	r := rect{
		width:  5,
		height: 10,
	}

	// rect 구조체와 면적 출력
	fmt.Printf("rect: %+v\n", r)
	fmt.Printf("area: %d\n", r.area())
}