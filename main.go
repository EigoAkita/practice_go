package main

import (
	"main/errorHandling"
	"main/makeWebApplications"
	"main/practiceCommon"
	"main/practiceDefer"
	"main/practiceEmbedded"
	"main/practiceFor"
	"main/practiceGoroutine"
	"main/practiceHttp"
	"main/practiceIf"
	"main/practiceJson"
	"main/practiceLogging"
	"main/practiceMethod"
	"main/practicePackage"
	"main/practicePanicAndRecover"
	practicepointer "main/practicePointer"
	"main/practiceRange"
	"main/practiceSqlite"
	"main/practiceStruct"
	"main/practiceSwitch"
	"main/trainingFirst"
	"main/trainingSecond"
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
	makeWebApplications.MakeWebApplications()
	return
	practiceSqlite.PracticeSqlite()
	practicePackage.PracticePackage()
	practiceJson.PracticeHmac()
	practiceJson.PracticeJson()
	practiceHttp.PracticeHttp()
	practiceGoroutine.PracticeGoroutine6()
	practiceCommon.PracticeCommon()
	practiceEmbedded.PracticeEmbedded()
	practiceMethod.PracticeMethod()
	practiceStruct.PracticeStruct()
	practicepointer.PracticePointer()
	trainingSecond.TrainingSecond()
	practicePanicAndRecover.PracticePanicAndRecover()
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
