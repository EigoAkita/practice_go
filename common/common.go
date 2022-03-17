package practiceCommon

import (
	"fmt"
	"strconv"
)

var isAction bool = false

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
	// fmt.Println("bool値返還")
	// practiceBool()
	// typeConversion()
	// practiceArray()
	// practiceMap()
	practiceByte()

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

func practiceBool() {
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
		Atoiの返り値はint,error

		変数定義時
		※error

		i, err := strconv.Atoi(s)
		(err未使用の場合に限る)

		※!error

		i, _ := strconv.Atoi(s)

	*/

	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("ERROR")
	}
	fmt.Printf("%T %v", i, i)

	// h := "Hello World"
	// fmt.Printf(string(h[0]))
}

func practiceArray() {
	/*
		var a [2]int
		a[0] = 200
		a[1] = 300
		fmt.Println(a)

		var b [2]int = [2]int{100, 200}
		fmt.Println(b)

		var b []int = []int {100,200}
		b = append(b,300)
		fmt.Println(b)
	*/

	/*
					n := []int{1, 2, 3, 4, 5, 6}
					fmt.Println(n)

					fmt.Println(n[2])
					fmt.Println(n[2:4])
					fmt.Println(n[:2])

					var board = [][]int{
						{1, 2, 3},
						{4, 5, 6},
						{7, 8, 9},
					}

					fmt.Println(board)
					n = append(n, 100,200,300,400,500,600)
					fmt.Println(n)


				///int配列　長さ3　キャパシティ5　初期化
				n := make([]int, 3, 5)
				fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
				n = append(n, 0, 0)
				fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

				n = append(n, 1, 2, 3, 4, 5, 6)
				fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)


			a := make([]int, 3)
			fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

			///0のスライス　メモリ確保有
			b := make([]int, 0)
			fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
			///初期値nil 　 メモリ確保無
			var c []int
			fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)



		// c := make([]int, 5)
		c := make([]int, 0, 5)
		fmt.Println(c)
		for i := 0; i < 5; i++ {
			c = append(c, i)
			fmt.Println(c)
		}
		fmt.Println(c)
	*/

}

func practiceMap() {

	m := map[string]int{"カフェオレ": 100, "バナナ": 100}
	fmt.Println(m)
	fmt.Println(m["カフェオレ"])

	m["カフェオレ"] = 300
	fmt.Println(m)
	m["あんぱん"] = 150
	fmt.Println(m)

	///map内key無　出力0
	fmt.Println(m["nil"])

	///スライス内のkey存在確認　有 = true 無 = false
	v, ok := m["カフェオレ"]
	fmt.Println(v, ok)
	_, ok2 := m["nothing"]
	if !ok2 {
		fmt.Println(ok2)
	}

	///メモリ上 空map作成
	m2 := make(map[string]int)
	m2["pc"] = 200
	fmt.Println(m2)

	///宣言しているが、メモリ上に代入するmapがnilの為ERRORになる
	//varの時は初期値はnil
	/*
		var m3 map[string]int
		m3["pc"] = 200
		fmt.Println(m3)
	*/

}

///ネットワーク系
///ファイルの読込などで使用
func practiceByte() {
	b := []byte{72, 73}
	fmt.Println(b)
	fmt.Println(string(b))

	c := []byte("HI")
	fmt.Println(c)
	fmt.Println(string(c))
}
