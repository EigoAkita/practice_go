package main

import (
	"fmt"
	"strconv"
	// "strings"
)

const Pi = 3.14
const (
	UserName = "test_user\n"
	Password = "test_password\n"
)

func constFunc() {
	// println(UserName, Password, Pi)
	/*
		x := 1 + 1
		fmt.Println("\n", x)
		fmt.Println("\n1 + 1 = ", 1+1)
	*/

	/*
		x := 0
		fmt.Println("\n", x)
		x++
		fmt.Println("\n", x)
	*/
	fmt.Println("bool値返還")
	practicBool()
	typeConversion()
	///文字列

	///"Hello World"の3番目を取得
	// fmt.Println("\n", string("Hello World"[3]))

	///文字列置換
	// var s string = "Chocolate"
	///s変数、"C"文字列、"c"文字列に変換、最初の1文字
	// fmt.Println("\n", strings.Replace(s, "C", "c", 1))

	// s = strings.Replace(s, "o", "A", 2)
	// fmt.Println("\n", s)
	// fmt.Println(strings.Contains(s, "A"))
}

func practicBool() {
	t, f := true, false
	fmt.Printf("%T %v %t\n", t, 0, t)
	fmt.Printf("%T %v %t\n", f, 0, f)
}

func typeConversion() {
	var x int = 1
	xx := float64(x)

	/*
	%T = 変数の型
	%v = 変数の値
	%f = 変数の値をfloat型に変換
	*/
	fmt.Printf("%T %v %f\n", xx, xx, xx)

	var s string = "14"

	/*
	string型をint型に変換
	Atoiの引数はstring
	Atoiの返り値はint,errorの為、変数を定義する時は i , _ := となる
	*/
	
	i, err := strconv.Atoi(s)
	if err != nil{
		fmt.Printf("ERROR")
	}
	fmt.Printf("%T %v", i, i)
}
