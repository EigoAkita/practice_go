package practiceDefer

import (
	"fmt"
	"os"
)

// Defer = 遅延実行
func PracticeDefer() {

	openFile()
	printHelloWorldDefer()
	indexRunDefer()
}

func openFile () {
	//ローカルのファイル読込
	file, _ := os.Open("practiceSwitch/practiceSwitch.go")
	/*
	最後にファイルをクローズする事が必要
	deferでクローズする宣言を最初に記載する事で、
	最後にファイルを閉じてくれるので、ファイルの閉じ忘れが無い
	*/
	defer file.Close()
	//ファイル読込時、バイト配列を記述する必要有
	data := make([]byte, 200)
	//ファイル読込
	file.Read(data)
	//バイト配列をstring型にキャスト
	fmt.Println(string(data))
}

func printHelloWorldDefer() {
	/*
		実行結果
		Hello
		World
	*/
	defer fmt.Println("World")
	fmt.Println("Hello")
}

func indexRunDefer() {
	/*
	staking defer
	deferを以下の様に列挙
	実行順 => 最後のdeferから順に実行
	run
	success
	3
	2
	1
	*/

	fmt.Println("run")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("success")
}
