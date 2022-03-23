package practiceSwitch

import (
	"fmt"
	"time"
)

func PracticeSwitch() {
	items := []int{
		0,
		1,
		2,
	}

	//通常のswitch文のパターン
	for i := 0; i < len(items); i++ {
		switch i {
		case 0:
			fmt.Println("iOS")
		case 1:
			fmt.Println("Android")
		case 2:
			fmt.Println("Linux")
		}
	}

	//関数呼出でswitch文を形成するパターン
	switch food := getFood(); food {
	case "焼肉":
		fmt.Println("GOOD TASTE")
	case "デスソース":
		fmt.Println("BAD TASTE")
	default:
		fmt.Println("HUNGRY")
	}

	food := getFood()
	switch {
	case food == "焼肉":
		fmt.Println("GOOD TASTE !!!")
	case food == "デスソース":
		fmt.Println("BAD TASTE !!!")
	default:
		fmt.Println("HUNGRY")
	}

	switch {

	}

	//switch文の記述不要パターン
	t := time.Now()
	fmt.Println(t.Hour())

	switch {
	case t.Hour() < 12:
		fmt.Println("Morning")
	case t.Hour() < 17:
		fmt.Println("Afternoon")
	}

}

func getFood() string {
	return "焼肉"
}
