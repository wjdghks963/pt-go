# GO

## 모듈

Go 모듈은 패키지 관리와 의존성 관리를 위한 Go의 최신 시스템입니다.
모듈은 관련된 Go 패키지의 모음으로, 특정 프로젝트에서 사용할 수 있는 코드와 의존성을 정의합니다.
모듈은 go.mod 파일에 의해 관리됩니다.

1.  모듈 초기화

    - go mod init <module-path> 사용시 새로운 모듈을 초기화
    - <module-path> 는 일반적으로 프로젝트가 호스팅되는 위치

2.  go.mod

    - 모듈의 기본 정보 포함. 모듈 경로, go 버전, 의존성 명시

3.  의존성 관리:
    - go get <package> 명령어를 사용하여 패키지를 추가하거나 업데이트합니다.
    - go mod tidy 명령어는 사용되지 않는 패키지를 제거하고 필요한 패키지를 추가합니다.
    - go install: 빌드된 파일을 환경 변수에 맞춰 지정된 위치($GOBIN)에 배치합니다.

npm의 package.json과 같은 역할

## 명령어

1. run : 빌드 & 실행
2. build : 빌드 -> 패키지 이름 바이너리 파일을 return -> 이 빌드 파일은 어떤 os에서든 사용 가능함
3. install : 환경 변수에 맞춰 빌드된 파일을 해당 위치에 떨굼 => GOBIN

만약 go의 환경변수를 찾지 못해 install 한 바이너리를 사용하지 못한다면

```go
echo 'export GOPATH=$HOME/go' >> ~/.zshrc # 또는 ~/.bashrc, ~/.profile
echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.zshrc # 또는 ~/.bashrc, ~/.profile
source ~/.zshrc # 또는 ~/.bashrc, ~/.profile
```

전역으로 설치된 바이너리를 삭제하고 싶다면

```go
// bin의 파일이 있는 장소
ls $HOME/go/bin
```

## import

함수를 import하기 위해선 해당하는 함수 명은 대문자로 시작해야합니다.

export 하는 코드

```go
package astruct

import "fmt"

// 공개된 함수 (다른 패키지에서 접근 가능)
func MakeMetStruct() {
	type Rect struct {
		Width  int
		Height int
	}

	// Rect 구조체의 메서드 정의
	func (r Rect) Area() int {
		return r.Width * r.Height
	}

	// Rect 구조체의 인스턴스 생성
	r := Rect{
		Width:  5,
		Height: 10,
	}

	// Rect 구조체와 면적 출력
	fmt.Printf("Rect: %+v\n", r)
	fmt.Printf("Area: %d\n", r.Area())
}

// 비공개 함수 (같은 패키지 내에서만 접근 가능)
func privateFunction() {
	fmt.Println("This is a private function")
}
```

사용하는 코드

```go
// main.go
package main

import (
	"myproject/astruct" // asutrct 패키지 가져오기
)

func main() {
	astruct.MakeMetStruct() // 공개된 함수 호출
	// astruct.privateFunction() // 오류: 비공개 함수는 호출할 수 없음
}

// go.mod
module github.com/myproject/astruct

go 1.22.4

require (
    github.com/junghwan/myproject/astructs v0.0.0
)

replace github.com/junghwan/myproject/astruct v0.0.0 => ./astruct
```

# 변수 값 타입

## 변수 선언

### short declaration operator

```go
x := 23
```

### var

var는 전역 스코프에서 변수를 사용하기 위한 예약어입니다.

```go
var x = 11
var z int
```

### const

상수는 const로 선언합니다.

```go
const Pi = 3.14
const Greeting = "Hello World"
```

## 타입

정적 타입을 가졌기 때문에 선언시에 타입을 지정합니다.

```go
var s string = "GO GO";
var z double
```

### 사용자 정의 타입 만들기

type 을 사용해 전용 타입을 만들 수 있음

```go
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

    // !오류
    // int에 hodog 타입을 대입할 수 없음
     a = b
}
```

### 타입 변환

go에서는 타입 캐스팅이라 하지 않고 타입 변환이라합니다.

위의 코드에서 변환을 사용하면 a=b 에서 발생한 오류를 없앨 수 있습니다.

```go
a = int(b)
```

### 타입 종류와 개념

원시 타입은 기본 값을 가지고 있습니다.

ex) bool = true, string = "", int = 0 ...

**iota**

상수 선언에 사용되는 특별한 식별자로, 상수 값을 자동으로 증가시키는 데 사용
iota는 const 선언 블록 내에서 첫 번째 상수부터 0부터 시작하여 각각의 상수마다 1씩 증가합니다. 이를 통해 반복적인 상수 값을 쉽게 정의할 수 있습니다.

<br>

# 루프 / 조건문

## 반복문 루프

go는 while do 와 같은 반복문 대신 for 를 사용합니다.

for 루프는 초기화, 조건, 후처리 부분을 모두 생략할 수 있습니다.

1. 일반 for loop

```go
	for i:=0; i<=10; i++ {
		for j:=0; j<=10; j++ {
			fmt.Printf("OUT :: %d\t INNER :: %d\n" , i, j)
		}
	}
```

2. while 스타일 (짧은 구문)

```go
    x := 1
    for x <10 {
        fmt.Println(x)
    }
```

```go
	y := 1
	for {
		if x > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
```

3. range

슬라이스, 배열, 맵, 채널 등을 순회할 때 유용하게 사용됩니다.

```go
nums := []int{2, 3, 4}
for i, num := range nums {
    fmt.Println("index:", i, "value:", num)
}
```

## 조건문

```go

if true {
    fmt.Print("TRUE !!")
}
```

### 초기화 조건문

조건문 안에서 변수를 초기화하면 스코프를 제한 시킬 수 있습니다.

```go
if x:32; x == 2{
    fmt.Println("x is 2")
}
```

### switch

스위치는 흐름제어를 위해 사용합니다

`fallthrough` 를 사용한다면 성공 case에 걸리지 않고 하위 case를 실행합니다.

```go
// 조건값의 기본값은 true 이기 떄문에 case가 true 인것을 검색
switch  {
    case false:
            fmt.Println("false")
    case 2==4:
            fmt.Println("false")

    case 2==2:
            fmt.Println("true1")
            fallthrough
    case 3==3:
            fmt.Println("true2")

    default:
            fmt.Println("default")
}
```

# 구조체

## struct

### 익명 구조체

일반적인 경우 익명 구조체를 사용하는 경우는 피해야합니다.

익명 구조체는 이름이 없는 구조체로, 일회성 사용이나 간단한 데이터 그룹화를 위해 유용합니다. 예를 들어, JSON payload를 처리하거나 간단한 보안 관련 데이터를 임시로 저장할 때 사용됩니다.

1회성으로 사용하는 경우는 사용하기 유용합니다.
ex) 보안관련, json payload 관련 등...

```go
	// 익명 구조체 -> 인스턴스 바로 생성
	myCar := struct {
		Make string
		Model string
	} {
		Make: "tesla",
		Model: "model 3",
	}
```

### 타입 지정 구조체

Go는 객체 지향 언어가 아니지만, 구조체와 인터페이스를 사용하여 객체 지향 프로그래밍과 유사한 스타일을 구현할 수 있습니다.
구조체를 사용하여 데이터와 메서드를 그룹화할 수 있으며, 다른 구조체를 포함하여 “상속”과 유사한 기능을 제공할 수 있습니다.

```go
type car struct {
    make string
    model string
}

type truck struct {
    // car의 구조체를 상속 받아 사용합니다.
    car
    bdeSize int
}

// 인스턴스 생성
myTruck := truck{
	bdeSize: 10,
	car: car{
		make: "kia",
		model: "bigModel",
		},
	}

```

### 구조체 포함(Embedding)과 상속의 차이

Go에서의 포함(embedding)은 명시적으로 재사용성을 높이는 방법이며, 명시적이지 않은 상속과 달리 더 직관적입니다.
포함된 구조체의 메서드나 필드를 오버라이딩하는 기능은 없지만, 포함된 구조체를 통해 비슷한 기능을 제공할 수 있습니다.

<br>

## interface

인터페이스를 명시적으로 선언하지 않고, 구조체가 인터페이스에 정의된 메서드를 모두 구현하면 해당 인터페이스를 자동으로 구현한 것으로 간주됩니다. 이를 “구조적 타이핑”이라고 합니다.

1. 구조적 타이핑: Go에서는 타입 선언 없이도 인터페이스를 구현할 수 있습니다. 인터페이스에 정의된 모든 메서드를 구현하면 해당 타입은 자동으로 그 인터페이스를 구현합니다.

2. 명시적 구현 선언이 필요 없음: 다른 언어와 달리, Go에서는 인터페이스를 구현한다고 명시적으로 선언할 필요가 없습니다. 필요한 메서드만 구현하면 됩니다.

3. 다양한 활용: 이 방법을 통해 코드의 유연성을 높이고, 다양한 타입을 처리할 수 있는 함수와 메서드를 작성할 수 있습니다.

```go
type Msg interface {
	getMsg() string
}

type Phone struct {
	model string
	category string
}

func (p Phone ) getMsg() string {
	return fmt.Sprintf("MODEL : %v, CATEGORY : %v", d.model, d.category)
}


type Watch struct {
	model string
	category string
}

func (w Watch ) getMsg() string {
	return fmt.Sprintf("MODEL : %v, CATEGORY : %v", d.model, d.category)
}

func sendMsg(m Msg) {
	fmt.Println(msg.getMsg())
}

func main(){
	phone := Phone{
		model:"IPhone",
		category:"phone"
	}

	watch := Watch{
		model:"apple watch",
		category:"watch"
	}


	sendMsg(phone)
	sendMsg(watch)
}
```

### 타입 어설션

타입 어설션을 사용하면 인터페이스 변수가 실제로 어떤 구체적인 타입을 갖는지 확인할 수 있습니다.

```go
type email struct {
	address string
}

type sms struct {
	number string
}

type message interface {
	getMessage() string
}

func (e email) getMessage() string {
	return "Email: " + e.address
}

func (s sms) getMessage() string {
	return "SMS: " + s.number
}

func main() {
	var msg message

	msg = email{address: "example@example.com"}

	// 타입 어설션
	if em, ok := msg.(email); ok {
		fmt.Println("Email address is:", em.address)
	} else {
		fmt.Println("Not an email")
	}

	msg = sms{number: "123-456-7890"}

	// 타입 어설션
	if sm, ok := msg.(sms); ok {
		fmt.Println("SMS number is:", sm.number)
	} else {
		fmt.Println("Not an SMS")
	}
}
```

# 에러

Go에서 에러는 error 인터페이스로 표현됩니다. error 인터페이스는 Error 메서드를 가지며, 이는 에러 메시지를 반환합니다.

```go
type error interface {
    Error() string
}
```

```go

user, err := getUser()
if err != nil {
	fmt.Println(err)
	return
}

profile, err  := getUserProfile(user.ID)
if err != nil {
	fmt.Println(err)
	return
}

```

## 커스텀 에러

에러는 errors 패키지의 New 함수나 fmt.Errorf 함수를 사용하여 생성할 수 있습니다.

1. New 함수 사용

```go
import (
    "errors"
    "fmt"
)

// errors.New를 사용하여 에러 생성
err := errors.New("an error occurred")

// fmt.Errorf를 사용하여 에러 생성
err = fmt.Errorf("an error occurred: %v", someValue)
```

2. error 인터페이스를 구현하는 사용자 정의 타입을 사용

```go
// 사용자 정의 에러 타입 정의
type divideError struct {
	dividend float64
	divisor  float64
}

// error 인터페이스 구현
func (de divideError) Error() string {
	return fmt.Sprintf(
		"cannot divide %v by %v",
		de.dividend, de.divisor,
	)
}
```

# Slice

슬라이스(Slice)는 동적 배열과 유사한 구조로, 크기가 **가변적**입니다. 슬라이스는 배열보다 더 유연하고 사용하기 쉽습니다.

따라서 직접적으로 array에 접근하는 일은 거의 없을 것입니다.

## slice 생성 방법

```go
	// 배열로부터 슬라이스 생성
	arr := [5]int{1, 2, 3, 4, 5}
	slice1 := arr[1:4] // 슬라이스: [2 3 4]

	// make 함수를 사용하여 슬라이스 생성
	slice2 := make([]int, 3) // 길이가 3인 슬라이스: [0 0 0]

	// 리터럴을 사용하여 슬라이스 생성
	slice3 := []int{6, 7, 8} // 슬라이스: [6 7 8]

	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)
```

## 슬라이스의 속성

슬라이스는 세 가지 속성을 가집니다:

- 길이(Length): 슬라이스에 포함된 요소의 개수
- 용량(Capacity): 슬라이스의 기저 배열에서 시작 위치부터 끝까지의 요소 개수
- 기저 배열(Underlying array): 슬라이스가 참조하는 배열

```go
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Length:", len(slice))   // 출력: Length: 5
	fmt.Println("Capacity:", cap(slice)) // 출력: Capacity: 5
```

## 함수 스프레드 연산자

```go
func printSlice(sl ...string){
	for i < len(sl){
		fmt.Prirntln(sl[i])
	}
}

// 가변 인자 함수 호출
printSlice("apple", "banana", "cherry")

// 슬라이스를 가변 인자로 전달할 때는 스프레드 연산자(...) 사용
fruits := []string{"apple", "banana", "cherry"}
printSlice(fruits...)
```

## append

1.  용량 내에서 확장:

    슬라이스의 용량이 충분하다면, append는 원본 슬라이스의 기저 배열을 수정합니다. 이 경우 원본 슬라이스와 새로운 슬라이스는 동일한 배열을 참조합니다.

2.  용량 초과:

    슬라이스의 용량이 부족하면, append는 새로운 배열을 할당하고, 기존 요소를 복사한 후 새로운 요소를 추가합니다. 이 경우 원본 슬라이스와 새로운 슬라이스는 다른 배열을 참조합니다.

```go
a := make([]int, 3)
fmt.Println("len of a:", len(a))
// len of a: 3
fmt.Println("cap of a:", cap(a))
// cap of a: 3
fmt.Println("appending 4 to b from a")
// appending 4 to b from a
b := append(a, 4)
fmt.Println("b:", b)
// b: [0 0 0 4]
fmt.Println("addr of b:", &b[0])
// addr of b: 0x44a0c0
fmt.Println("appending 5 to c from a")
// appending 5 to c from a
c := append(a, 5)
fmt.Println("addr of c:", &c[0])
// addr of c: 0x44a180
fmt.Println("a:", a)
// a: [0 0 0]
fmt.Println("b:", b)
// b: [0 0 0 4]
fmt.Println("c:", c)
// c: [0 0 0 5]
```

**! 주의점**

append 함수를 사용할 때 반환된 슬라이스를 원래 슬라이스에 다시 할당하지 않으면, 다음과 같은 문제가 발생할 수 있습니다:

1. 원래 슬라이스가 변경되지 않음: append 함수가 반환하는 새로운 슬라이스를 사용하지 않으면, 원래 슬라이스는 변경되지 않습니다. -> 유연하지 않음
2. 혼란과 버그 발생 가능성: 코드를 유지보수하거나 읽을 때 혼란이 생기고, 버그가 발생할 가능성이 높아집니다.

따라서 append를 사용할 땐 새로운 변수에 할당하면 안되고 원본 슬라이스에 재할당해야합니다.

GOOD

```go
slice := []int{1, 2, 3}
	fmt.Println("Before append:", slice) // 출력: Before append: [1 2 3]

	// append 함수의 결과를 원래 슬라이스에 다시 할당
	slice = append(slice, 4)
	fmt.Println("After append:", slice)  // 출력: After append: [1 2 3 4]
```

BAD

```go
	slice := []int{1, 2, 3}
	fmt.Println("Before append:", slice) // 출력: Before append: [1 2 3]

	// append 함수의 결과를 다른 변수에 할당
	newSlice := append(slice, 4)
	fmt.Println("Original slice after append:", slice) // 출력: Original slice after append: [1 2 3]
	fmt.Println("New slice after append:", newSlice)   // 출력: New slice after append: [1 2 3 4]
```

## range

range는 슬라이스, 배열, 맵, 문자열 등의 요소를 반복(iterate)하는 데 사용됩니다.

두 개의 값을 반환합니다: 인덱스와 해당 인덱스의 요소.

```go
	// slice에 대한 range
	numbers := []int{1, 2, 3, 4, 5}

	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num)
	}


	// map에 대한 range
	dict := map[string]int{"one": 1, "two": 2, "three": 3}

	for key, value := range dict {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}


	// 문자열에 대한 range
	str := "hello"

	for i, ch := range str {
		fmt.Printf("Index: %d, Character: %c\n", i, ch)
	}
```

<br>

# Map

### 맵 생성

Js의 object와 python의 dict 과 같은 기능입니다.

빈값이면 nil을 반환합니다.

```go
ages := make(map[string]int)
ages["J"] = 32
ages["C"] = 22


ages := map[string]int{
	"J":32,
	"C":22
}

```

### 맵 요소 추가 및 삭제

```go
	ages := map[string]int{
		"J": 32,
		"C": 22,
	}

	// 요소 추가
	ages["A"] = 27

	// 요소 삭제
	delete(ages, "C")
```

### 맵 요소 접근

요소 접근은 키 값으로 할 수 있으며 만약 값의 존재 유무에 대한 처리를 결정 할 수 있습니다.

1. 존재 여부 확인 후 조건문
2. 맵의 값에 대한 자료형으로 기본값 받기

```go
	ages := map[string]int{
		"J": 32,
		"C": 22,
	}

		// 존재 여부 확인
	ageX, exists := ages["X"]
	if exists {
		fmt.Println("Age of X:", ageX)
	} else {
		fmt.Println("X does not exist in the map") // 출력: X does not exist in the map
	}

	// 키가 존재하지 않을 때 기본 값 반환
	ageY := ages["Y"]
	fmt.Println("Age of Y:", ageY) // 출력: Age of Y: 0 (int의 제로 값)

```

### 주의점

Go의 맵(map)에서 **키는 반드시 비교 가능한 타입**이어야 하며, 일반 자료형을 사용해야 하는 이유는 Go 언어의 타입 시스템과 비교 연산의 특성 때문입니다.

- 숫자 타입(int, float 등)
- 문자열(string)
- 불리언(bool)
- 포인터(pointer)
- 인터페이스(interface)
- 구조체(struct) (모든 필드가 비교 가능한 경우)
- 배열(array) (모든 요소가 비교 가능한 경우)

일반 자료형을 사용하는 이유

1. 비교 가능성:
   - 맵의 키는 내부적으로 해시 테이블을 사용하여 저장됩니다. 해시 테이블은 키의 해시 값을 계산하고, 해시 값이 같은 키들에 대해 == 연산자를 사용하여 비교합니다.
   - 일반 자료형(숫자, 문자열 등)은 비교 연산이 정의되어 있어, 해시 테이블의 키로 사용하기에 적합합니다.
2. 일관성과 안정성:
   - 비교 가능한 타입은 일관된 동작을 보장하며, 키 비교 시 예상치 못한 결과를 피할 수 있습니다.
   - 예를 들어, 구조체 타입을 키로 사용할 때 구조체의 모든 필드가 비교 가능해야 하며, 구조체의 비교 연산이 정의되어 있어야 합니다.

# Advanced Functions (일급 함수 & 고차 함수)

## 일급 함수 (First-Class Functions)

일급 함수는 다음과 같은 특징을 가집니다:

1. 함수는 변수에 할당될 수 있습니다.
2. 함수는 다른 함수의 인자로 전달될 수 있습니다.
3. 함수는 다른 함수의 반환 값이 될 수 있습니다.

## 고차 함수 (Higher-Order Functions)

고차 함수는 힘수를 인자로 받아 유연한 코드 작성을 할 수 있도록 도와줍니다.
결과가 상황에 따라 바뀌고 그것을 처리하는 곳에 유용합니다. ex) httpHandler 과 같은 handler 류

1. 함수를 인자로 받을 수 있습니다.
2. 함수를 반환할 수 있습니다.

## 커링

커링은 여러 개의 인자를 받는 함수를 일련의 함수로 변환하여 각 함수가 하나의 인자만 받도록 하는 기법입니다. Go는 순수 함수형 언어가 아니므로 커링을 직접적으로 지원하지는 않지만, 클로저와 고차 함수를 사용하여 비슷한 효과를 얻을 수 있습니다.

```go
// 두 개의 인자를 받아 더하는 함수
func add(a, b int) int {
	return a + b
}

// 하나의 인자를 받아 부분적으로 적용된 함수 생성
func curriedAdd(a int) func(int) int {
	return func(b int) int {
		return add(a, b)
	}
}

func main() {
	// 5를 인자로 받아 새로운 함수를 생성
	addFive := curriedAdd(5)

	// 새로운 함수는 하나의 인자만 받음
	result := addFive(3) // 5 + 3
	fmt.Println("Result:", result) // 출력: Result: 8
}
```

## Defer

Go의 defer 키워드는 매우 유용한 기능으로, 특정 함수 호출을 현재 함수가 반환되기 직전에 자동으로 실행되도록 예약합니다. defer 키워드는 주로 리소스를 정리하거나, 파일 핸들러나 데이터베이스 연결을 닫는 등의 작업에 사용됩니다.

**defer 키워드의 특징**

1. 지연 실행: defer 키워드를 사용한 함수 호출은 그 함수의 실행이 끝날 때까지 지연됩니다.
2. 인수 평가: defer를 사용할 때 함수의 인수는 즉시 평가되지만, 함수 자체는 지연 실행됩니다.
3. 여러 개의 defer: 여러 개의 defer 호출이 있을 경우, 후입선출(LIFO) 순서로 실행됩니다.

```go

func main() {
	// 파일 열기
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	// main 함수가 반환되기 직전에 파일 닫기
	defer file.Close()

	// 파일 작업 수행
	// ...
}
```

## 익명 함수

익명 함수(Anonymous Functions)는 이름이 없는 함수로, 한 번만 사용하거나 빠르게 클로저(Closure)를 생성할 때 유용합니다. Go 언어에서는 익명 함수를 정의하고 즉시 실행하거나 변수에 할당할 수 있습니다. 익명 함수는 특히 고차 함수나 콜백 함수로 사용될 때 유용합니다.

1. 즉발

```go
func main() {
    // 익명 함수 정의와 동시에 실행
    func(message string) {
        fmt.Println(message)
    }("Hello, World!")
}
```

2. 변수에 할당

```go
    // 익명 함수를 변수에 할당
    greet := func(name string) string {
        return fmt.Sprintf("Hello, %s!", name)
    }

    // 할당된 익명 함수 호출
    message := greet("Alice")
    fmt.Println(message) // 출력: Hello, Alice!
```

3. 클로저(Closure) 생성

```go
  // 클로저 생성
    counter := func() func() int {
        count := 0
        return func() int {
            count++
            return count
        }
    }()

    // 클로저 호출
    fmt.Println(counter()) // 출력: 1
    fmt.Println(counter()) // 출력: 2
    fmt.Println(counter()) // 출력: 3
```

# _ Pointer _

포인터는 다른 변수의 메모리 주소를 저장하는 변수입니다. 포인터는 데이터 자체가 아닌 데이터가 저장된 위치를 가리킵니다.

## 포인터 선언 및 사용

1. 포인터 선언:
   포인터를 선언할 때는 \* 문법을 사용합니다.

```go
var p *int
```

2. 포인터 할당:
   포인터를 할당할 때는 & 연산자를 사용하여 변수의 주소를 가져옵니다.

```go
myString := "hello"
myStringPtr := &myString //  myString 변수의 메모리 주소
```

3. 포인터 역참조:
   포인터가 가리키는 값을 가져오려면 \* 연산자를 사용하여 역참조(dereference)합니다.

```go
fmt.Println(*myStringPtr) // 출력: hello
```

4. 포인터 값 변경

```go
myString := "hello"

// 포인터 선언 및 주소 할당
myStringPointer = &myString

// 포인터를 통해 값 변경
*myStringPorinter = "helloPointer" // myString은 "helloPointer"가 된다.
```

## 리시버(receiver)의 종류

이 메서드가 어떤 타입에 속해 있는지를 나타냅니다. 리시버는 함수가 특정 타입의 메서드임을 명시합니다.

### 값 리시버(Value Receiver)

값 리시버는 타입의 값을 복사하여 메서드에 전달합니다.
리시버 타입을 포인터로 지정하지 않으면 값 리시버가 됩니다.
값 리시버를 사용하는 메서드는 원본 값을 변경하지 않습니다.

```go
type Rectangle struct {
	width, height int
}

// 값 리시버 메서드
func (r Rectangle) Area() int {
	return r.width * r.height
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Area:", rect.Area()) // 출력: Area: 50
}
```

### 포인터 리시버(Pointer Receiver)

포인터 리시버는 타입의 포인터를 메서드에 전달합니다.
리시버 타입을 \*Type으로 지정하면 포인터 리시버가 됩니다.
포인터 리시버를 사용하는 메서드는 원본 값을 변경할 수 있습니다.

```go
type Rectangle struct {
	width, height int
}

// 포인터 리시버 메서드
func (r *Rectangle) Scale(factor int) {
	r.width *= factor
	r.height *= factor
}

func main() {
	rect := Rectangle{width: 10, height: 5}
	fmt.Println("Before scaling:", rect) // 출력: Before scaling: {10 5}
	rect.Scale(2)
	fmt.Println("After scaling:", rect)  // 출력: After scaling: {20 10}
}
```

## 포인터의 유용성

1. 효율성:
   포인터를 사용하면 **데이터를 복사하지 않고도 직접 메모리에서 데이터를 조작**할 수 있으므로 프로그램의 성능을 향상시킬 수 있습니다.

2. 공유된 데이터:
   여러 함수가 **동일한 데이터를 직접 수정**할 수 있습니다. 이를 통해 함수 간에 데이터를 효율적으로 전달하고 수정할 수 있습니다.

</br>

# Concurrency & Parallelism

## Concurrency(동시성)

동시성은 여러 작업을 동시에 처리할 수 있는 능력입니다. 동시성은 단일 프로세서에서 이루어질 수도 있고, 여러 프로세서에서 이루어질 수도 있습니다. 동시성의 목표는 시스템이 여러 작업을 ‘동시에’ 처리하는 것처럼 보이게 하는 것입니다. 이는 멀티태스킹, 즉 작업을 빠르게 전환함으로써 이루어집니다.

```go
package main

import (
    "fmt"
    "time"
)

func say(s string) {
    for i := 0; i < 5; i++ {
        time.Sleep(100 * time.Millisecond)
        fmt.Println(s)
    }
}

func main() {
    go say("world")
    say("hello")
}
```

## Parallelism(병렬성)

병렬성은 실제로 여러 작업이 동시에 실행되는 것입니다. 이는 다중 코어 프로세서에서 여러 작업을 병렬로 실행하여 성능을 향상시킵니다. 병렬성은 동시성을 달성하기 위한 방법 중 하나입니다.

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func say(s string, wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 5; i++ {
        fmt.Println(s)
    }
}

func main() {
    runtime.GOMAXPROCS(2) // 사용할 코어 수를 2로 설정
    var wg sync.WaitGroup

    wg.Add(2)
    go say("world", &wg)
    go say("hello", &wg)

    wg.Wait()
}
```

## Channel

채널은 고루틴 간에 데이터를 주고받는 통로입니다. 채널을 사용하면 여러 고루틴이 안전하게 데이터를 주고받을 수 있습니다.
채널은 다음과 같은 방식으로 작동합니다

1. 채널 생성: make 함수로 채널을 생성합니다.

```go
c := make(chan int)
```

2. 데이터 전송: 채널을 통해 데이터를 전송(<- 연산자 사용)합니다.

```go
c <- value
```

3. 데이터 수신: 채널을 통해 데이터를 수신(<- 연산자 사용)합니다.

```go
value := <-c
```

### **데드락**

Go에서는 채널을 잘못 사용하거나, 고루틴이 서로를 기다리는 상황을 잘못 설정하면 데드락이 발생할 수 있습니다.

```go
package main

func main() {
    c := make(chan int)
    c <- 1 // 여기서 고루틴이 멈춥니다.
    <-c
}
```

메인 고루틴이 채널에 데이터를 보내지만, 동일한 고루틴에서 데이터를 받으려 하기 때문에 채널의 **수신자가 없어 무한히 대기**하게 됩니다.

데드락을 피하기 위해선 시용 후 채널을 닫거나 확실히 발송을 시작한다면 리시버가 있어야합니다.
