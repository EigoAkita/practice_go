package practiceCommon

import (
	"fmt"
)

// type MyInt int

// ///無名関数
// func (i MyInt) Double() int {
// 	fmt.Printf("%T %v\n", i, i)
// 	fmt.Printf("%T %v\n", 1, 1)
// 	return int(i * 2)
// }

// type Human interface {
// 	Say() string
// }

// type Person struct {
// 	Name string
// }

// func (p *Person) Say() string {
// 	p.Name = "Mr." + p.Name
// 	fmt.Println(p.Name)
// 	return p.Name
// }

// type Dog struct {
// 	Name string
// }

// func DriveCar(human Human) {
// 	if human.Say() == "Mr.Mike" {
// 		fmt.Println("Run")
// 	} else {
// 		fmt.Println("Get out")
// 	}
// }

func PracticeCommon() {
	// myInt := MyInt(10)
	// fmt.Println(myInt.Double())
	// var mike Human = &Person{"Mike"}
	// mike.Say()
	// DriveCar(mike)
	// doInt(10)
	// var i interface{} = 10
	// doInt(i)
	// doString("Hello World")
	// do("Mike")
	// do(true)
	// doCommon(10)
	// doCommon("Mike")
	// doCommon(true)

	// var i int = 10
	// ii := float64(10)
	// fmt.Println(i, ii)

	// mike := Person{"Mike", 22}
	// fmt.Println(mike)

	if err := myFunc(); err != nil {
		fmt.Println(err)
	}

}

type UserNotFound struct {
	UserName string
}

func (e *UserNotFound) Error() string {
	return fmt.Sprintf("User Not Found : %v", e.UserName)
}

func myFunc() error {
	ok := false
	if ok {
		return nil
	}
	return &UserNotFound{UserName: "Mike"}
}

// func (p Person) String() string {
// 	return fmt.Sprintf("My name is %v", p.Name)
// }

// type Person struct {
// 	Name string
// 	Age  int
// }

// func doCommon(c interface{}) {
// 	//switch type 文
// 	switch v := c.(type) {
// 	case int:
// 		fmt.Println(v * 2)
// 		break
// 	case string:
// 		fmt.Println(v + "!")
// 		break
// 	default:
// 		fmt.Printf("I don't know %T\n", v)
// 	}

// }

// func doInt(i interface{}) {
// 	//typeアサーション
// 	ii := i.(int)
// 	ii *= 2
// 	fmt.Println(ii)
// }

// func doString(s interface{}) {
// 	ss := s.(string)
// 	fmt.Println(ss)
// }
