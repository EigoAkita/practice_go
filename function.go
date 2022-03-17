package main

import "fmt"

func commonFunction() {
	/*
	add(20, 30)
	r := returnAdd(50, 50)
	fmt.Println(r)
	fmt.Println(returnBool())
	fmt.Println(returnMap())
	r3 := call(20,30)
	fmt.Println(r3)
	fmt.Println(convert(20))

	f := func () int {
		return 25
	}
	fmt.Println(f())
  func (x int) {
		fmt.Println(x)
	}(2022)
	*/
	practiceClosure()
}

func add(x int, y int) {
	fmt.Println(x + y)
}

func returnAdd(x int, y int) int {
	return x + y
}

func returnBool() bool {
	return false
}

func returnMap() map[string]int {
	return map[string]int{
		"ミルク": 200,
	}
}

///引数にint型のprice,itemを代入
///int型のresultの変数を戻り値にする様に定義
func call (price,item int)(result int){
	result = price * item
	///resultを戻り値にする事は決まっているので、returnする際にresultは省略可能
	// return result
	return
}

func convert (price int) float64 {
	return float64(price)
}


/*
【クロージャ―】
practiceClosure関数に直接xの変数を定義
⇒ 他のコードがxを書き換える恐れ有
⇒ 処理を分離する事で影響範囲が及ばない様にする
*/

/*
デバッグ使用方法
ステップイン ⇒ 処理の中に入る
ステップオーバー ⇒　次の別の処理を実行（処理の中に入らない）
二つの使い分けが重要
*/

func incrementGenerator () (func() int) {
	x := 0
    return func()int{
		x++
		return x
	}
}
func practiceClosure () {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())
}