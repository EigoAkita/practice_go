package practiceMethod

import "fmt"

type Vertex struct {
	X, Y int
}

func (v Vertex) Area() int {
	return v.X * v.Y
}

//値レシーバー
func Area(v Vertex) int {
	return v.X * v.Y
}

//ポインタレシーバー
func (v *Vertex) Scale(i int) {
	v.X = v.X * i
	v.Y = v.Y * i
}

func PracticeMethod() {
	v := Vertex{3, 4}
	fmt.Println(Area(v))
	fmt.Println(v.Area())

	v.Scale(10)
	fmt.Println(v.Area())
}
