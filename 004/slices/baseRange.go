package slices

import "fmt"


func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badWord := range badWords {
			if word == badWord {
				return i
			}
		} 
	}

	return -1
}


func rangeTest(msg []string, badWords []string) {
	i := indexOfFirstBadWord(msg, badWords)
	fmt.Printf("Scanning message: %v for bad words:\n", msg)
	for _, badWord := range badWords {
		fmt.Println(` -`, badWord)
	}
	fmt.Printf("Index: %v\n", i)
	fmt.Println("====================================")
}

func baseRange(){

	badWords := []string{"crap", "shoot", "dang", "frick"}
	message := []string{"hey", "there", "john"}
	rangeTest(message, badWords)

	message = []string{"ugh", "oh", "my", "frick"}
	rangeTest(message, badWords)

	message = []string{"what", "the", "shoot", "I", "hate", "that", "crap"}
	rangeTest(message, badWords)
}