package functions

import "fmt"

func printReports(messages []string) {
	// messages를 받아 message를 사용
	for _, message := range messages {
		// 함수를 인자로 받는 함수 사용 이떄 함수는 이름이 존재하지 않음
		// 인자로 사용할 함수를 작성
		printCostReport(func(msg string) int {
			return len(msg) * 2
		}, message)
	}
}


func test(messages []string) {
	defer fmt.Println("====================================")
	printReports(messages)
}

func main() {
	test([]string{
		"Here's Johnny!",
		"Go ahead, make my day",
		"You had me at hello",
		"There's no place like home",
	})

	test([]string{
		"Hello, my name is Inigo Montoya. You killed my father. Prepare to die.",
		"May the Force be with you.",
		"Show me the money!",
		"Go ahead, make my day.",
	})
}

// 아래 함수를 인자로 받음
// func(msg string) int {
// 	return len(msg) * 2
// }
func printCostReport(costCalculator func(string) int, message string) {
	// cost는 인자로 받은 함수를 사용
	// 메시지 길이의 두배가 cost
	cost := costCalculator(message)
	fmt.Printf(`Message: "%s" Cost: %v cents`, message, cost)
	fmt.Println()
}