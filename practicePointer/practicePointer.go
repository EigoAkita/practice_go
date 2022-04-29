package practicepointer

import (
	"fmt"
)

func PracticePointer() {
	differenceBetweenNewAndMake()
	return

	// practicePointerFirst()
	/*
		1.nの変数を宣言
		2.関数（one）に100のアドレスを代入
		3.oneの引数に100のアドレスが代入
		4.100のアドレスの中に1を代入(100のアドレス : 1 ←こんな感じ)
		5.fmt.Println(n)でnを出力
		6.出力結果「1」
		7.元々定義していたvar n int = 100の中身も書き換えてしまう（デリファレンス）
	*/
	var n int = 100
	one(&n)
	fmt.Println(n)
}

func one(x *int) {
	*x = 1
}

func practicePointerFirst() {
	var n int = 100
	fmt.Println(n)

	//100を代入したメモリのアドレス
	fmt.Println(&n)

	//integerのpointer型
	//& = アンパサンド
	var p *int = &n
	fmt.Println(p)

	//アドレスが差してるメモリの中身が表示
	fmt.Println(*p)
}

//Lesson「newとmakeの違い」
func differenceBetweenNewAndMake() {
	//返り値がポインタ     => new
	//返り値がポインタ以外 => make
	s := make([]int, 0)
	fmt.Printf("%T\n", s)

	m := make(map[string]int)
	fmt.Printf("%T\n", m)

	ch := make(chan int)
	fmt.Printf("%T\n", ch)

	var q *int = new(int)
	fmt.Printf("%T\n", q)

	var st = new(struct{})
	fmt.Printf("%T\n", st)

	return

	/*
		var n int = 100
		var p *int = &n
	*/

	//値無しの状態で、ポインタが入る領域を確保したい場合↓↓↓↓↓
	var p *int = new(int)
	//0のアドレス
	fmt.Println(p)
	//pのアドレスに入ってる値
	fmt.Println(*p)
	*p++
	fmt.Println(*p)

}
