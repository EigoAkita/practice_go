package practiceIf

import "fmt"

func by2(num int) string {
	if num%2 == 0 {
		return "YES"
	} else {
		return "NO"
	}
}

func PracticeIf() {
	result := by2(10)
	fmt.Println(result)

	if result == "YES" {
		fmt.Println("great")
	}

	//14~19の処理を結合した処理
	//処理は記述可能、否result2を以降変数として出力は不可
	if result2 := by2(10); result2 == "YES" {
		fmt.Println("great 2")
	}
	/*
		num := 9
		if num%2 == 0 {
			fmt.Println("by 2")
		} else if num%3 == 0 {
			fmt.Println("by 3")
		} else {
			fmt.Println("failed")
		}

		x, y := 10,10
		if x== 10 && y == 10 {
			fmt.Println("&&")
		}

		if x == 10 || y == 10 {
			fmt.Println("||")
		}
	*/
}
