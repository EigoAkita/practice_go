package main

import (
	"main/errorHandling"
	"main/practiceDefer"
	"main/practiceFor"
	"main/practiceIf"
	"main/practiceLogging"
	"main/practicePanicAndRecover"
	"main/practiceRange"
	"main/practiceSwitch"
	"main/trainingFirst"
)

///変数宣言　=> var
var (
	// n int    = 100
	// s string = "test"
	// var t bool = true
	// var f bool = false
	// t, f bool = true, false
)

func main() {
	// fmt.Print(n, s, t, f)
	/// := ショート変数 => 関数内限定定義
	// xn := "\n"
	// xi := 1
	// xf64 := 1.2
	// xs := "2回目テスト"
	// xt := true
	// xf := false
	// fmt.Println(xn, xi, xs, xt, xf)
	// fmt.Printf("\n%T", xf64)
	// fmt.Println(user.Current())
	// fmt.Print(itemCount())
	// constFunc()
	// commonFunction()
	practicePanicAndRecover.PracticePanicAndRecover()
	return
	errorHandling.ErrorHandling()
	practiceLogging.PracticeLogging()
	practiceDefer.PracticeDefer()
	practiceSwitch.PracticeSwitch()
	practiceRange.PracticeRange()
	practiceFor.PracticeFor()
	trainingFirst.TrainingFirst()
	trainingFirst.TrainingFirstAnswer()
	practiceIf.PracticeIf()
}

/*
func itemCount () []int {
	var sum int
	var arr[] int
	fmt.Print(sum)
	for i := 0; i <= 10; i++ {
		arr = append(arr, i)
	}
	return arr;
}
*/
