package trainingFirst

import "fmt"

func TrainingFirst() {
	return
	//Q.1
	f := 1.11
	fmt.Println(int(f))

	//Q.2
	/*
		出力結果
		[5,6]
	*/

	//Q.3
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Milk": 30,
	}
	fmt.Printf("%T %v", m, m)
}

func TrainingFirstAnswer() {
	return
	//Q.1
	f := 1.11
	i := int(f)
	fmt.Printf("\n%T %v", i, i)
}
