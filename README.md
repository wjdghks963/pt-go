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
2. build : 빌드
3. install : 환경 변수에 맞춰 빌드된 파일을 해당 위치에 떨굼 => GOBIN

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
