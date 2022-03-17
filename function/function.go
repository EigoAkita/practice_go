package practiceFunction
import "fmt"

func CommonFunction() {
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
	// practiceClosure()
	practiceVariadicArgument()

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
func call(price, item int) (result int) {
	result = price * item
	///resultを戻り値にする事は決まっているので、returnする際にresultは省略可能
	// return result
	return
}

func convert(price int) float64 {
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

func incrementGenerator() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

/*
返り値func (radius float64) float64
初期化時、3.14がpiの引数として代入(初期化時に代入した変数は、以後変化せず値が保持)
次に、インナーファンクションのradiusの引数は未代入
2がradiusの引数として代入
参考変数【1】参照
以下のクロージャーを使用する理由として、
引数として入ってくる値が変化したとしても対応できる様にする為です。
*/
func circleArea(pi float64) func(radius float64) float64 {
	return func(radius float64) float64 {
		fmt.Println("pi", pi)
		fmt.Println("radius", radius)
		return pi * radius * radius
	}
}

func practiceClosure() {
	counter := incrementGenerator()
	fmt.Println(counter())
	fmt.Println(counter())
	fmt.Println(counter())

	///参考変数【1】
	c1 := circleArea(3.14)
	fmt.Println(c1(2))
	fmt.Println(c1(4))
	fmt.Println(c1(5))

	c2 := circleArea(3)
	fmt.Println(c2(2))
}

///可変長引数
func practiceVariadicArgument() {
	foo()
	foo(20, 30, 40)

	s := []int{1, 2, 3}
	//foo(1,2,3)の様に展開 ⇒ s...
	foo(s...)
}

func foo(params ...int) {
	//len = length
	fmt.Println(len(params), params)
	/*
		以下for文
		引数の数に応じてfor文が回る
		for index,param ⇒ [0赤,1黄2,青]
		for _,param ⇒ [赤,黄,青]
	*/
	for _, param := range params {
		fmt.Println(param)
	}

}
