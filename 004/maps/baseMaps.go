package maps

import (
	"errors"
	"fmt"
)


func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	userMap := make(map[string]user)
	if len(names) != len(phoneNumbers){
		return nil, errors.New("invalid sizes");
	}

	for i:=0; i < len(names); i++{
		name := names[i]
		phoneNumber := phoneNumbers[i]
		userMap[name] = user{
			name: name,
			phoneNumber: phoneNumber,
		}
	}
	return userMap, nil
}


type user struct {
	name        string
	phoneNumber int
}

func baseMaptest(names []string, phoneNumbers []int) {
	fmt.Println("Creating map...")
	defer fmt.Println("====================================")
	users, err := getUserMap(names, phoneNumbers)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, name := range names {
		fmt.Printf("key: %v, value:\n", name)
		fmt.Println(" - name:", users[name].name)
		fmt.Println(" - number:", users[name].phoneNumber)
	}
}

func baseMaps(){
	baseMaptest(
		[]string{"John", "Bob", "Jill"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	baseMaptest(
		[]string{"John", "Bob"},
		[]int{14355550987, 98765550987, 18265554567},
	)
	baseMaptest(
		[]string{"George", "Sally", "Rich", "Sue"},
		[]int{20955559812, 38385550982, 48265554567, 16045559873},
	)
}