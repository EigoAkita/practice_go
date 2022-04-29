package practiceStruct

import "fmt"

type Vertex struct {
	//struct内の変数を小文字にするとプライベート変数
	X int
	Y int
	S string
}

func changeVertex(v Vertex) {
	v.X = 1000
}

func changeVertex2(v *Vertex) {
	//vの実体に行くには*を付ける必要がある
	// (*v).X = 1000

	//structの場合は、自動的以下のvのポインタの実体を
	//指してくれる為、(*v).Xと書かなくて良い
	v.X = 1000
}

func PracticeStruct() {
	practiceStructSecond()
	return
	v := Vertex{X: 1, Y: 2}
	fmt.Println(v)

	fmt.Println(v.X, v.Y)

	v.X = 100
	fmt.Println(v.X, v.Y)

	v2 := Vertex{X: 1}
	fmt.Println(v2)
	v3 := Vertex{X: 1, Y: 2, S: "テスト"}
	fmt.Println(v3)
	v4 := Vertex{}
	fmt.Printf("%T %v \n", v4, v4)
	//結果 : practiceStruct.Vertex {0 0 }

	var v5 Vertex
	fmt.Printf("%T %v \n", v5, v5)
	//結果 : practiceStruct.Vertex {0 0 }

	v6 := new(Vertex)
	fmt.Printf("%T %v \n", v6, v6)
	//結果 : *practiceStruct.Vertex &{0 0 }

	v7 := &Vertex{}
	fmt.Printf("%T %v \n", v7, v7)
	//結果 : *practiceStruct.Vertex &{0 0 }
	/*
		実際には、ポインタが返り値として返ってくるのを
		明示的に分かりやすくする為、&Vertex{}と変数に
		代入する事が多い
	*/

	//スライスやmapの表現は以下のmakeを使用して表現
	//する事が好ましい
	s := make([]int, 0)
	// s := []int{}
	fmt.Println(s)

}

func practiceStructSecond() {
	v := Vertex{1, 2, "テスト"}
	changeVertex(v)
	fmt.Println(v)

	v2 := &Vertex{1, 2, "テスト2"}
	changeVertex2(v2)

	fmt.Println(v2)
	fmt.Println(*v2)
}
