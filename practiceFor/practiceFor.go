package practiceFor

import "fmt"

func PracticeFor() {
	for i := 0; i < 10; i++ {
		if i == 3 {
			fmt.Println("contine")
			//i==3の場合continue
			//処理がスキップ
			continue
		}
		if i > 5 {
			///i==5の場合、forループ文を抜け処理が停止
			fmt.Println("break")
			break
		}
		fmt.Println("番号", i)
	}

	sum := 1
	//for ; sum < 10 ; ⇐ セミコロン省略可能
	for  sum < 10 {
		sum += sum
		fmt.Println(sum)
	}

	fmt.Println(sum)
}
