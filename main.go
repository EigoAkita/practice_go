package main

import (
	"fmt"
	// "os/user"
	// "time"
)

///変数宣言　=> var
var (
	n int    = 100
	s string = "test"
	// var t bool = true
	// var f bool = false
	t, f bool = true, false
)

func main() {
	fmt.Print(n, s, t, f)
	/// := ショート変数 => 関数内限定定義
	xn := "\n"
	xi := 1
	xf64 := 1.2
	xs := "2回目テスト"
	xt := true
	xf := false
	fmt.Println(xn, xi, xs, xt, xf)
	fmt.Printf("\n%T", xf64)
	constFunc()
	// fmt.Println(user.Current())
	// fmt.Print(itemCount())
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
