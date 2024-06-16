package ainterface

import (
	"fmt"
	"time"
)

// message interface를 인자로 받는 함수
// interface 안에 있는 메서드에 접근
func sendMessage(msg message){
	fmt.Println(msg.getMessage())
}

// interface를 만듬
// getMessage라는 메서드를 가지고 있음
type message interface {
	getMessage() string
}

// 새로운 구조체 생성
type birthdayMessage struct {
	birthdayTime time.Time
	recipientName string
}

// 구조체에 getMessage라는 메서드 정의
func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}


// test는 message라는 인터페이스를 인자로 받음
func testInter(m message){
	sendMessage(m)
	fmt.Println("=========================")
}

func baseInter(){
	testInter(sendingReport{
		reportName:"reporter name",
		numberOfSends:10,
	})

	testInter(birthdayMessage{
		recipientName: "John",
		birthdayTime: time.Date(2011,01,31,0,0,0,0,time.UTC),
	})

}


