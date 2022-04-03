package trainingSecond

import (
	"fmt"
	"sort"
)

func minInt(a []int) int {
	result := sort.IntSlice(a)
	sort.Sort(result)
    return result[0]
}

func TrainingSecond() {
	/*
		Q1 . 以下のスライスから一番小さい数を探して出力するコードを書いてください。
	*/

	l := []int{100, 300, 23, 11, 23, 2, 4, 6, 4};

	lResult := minInt(l)

	fmt.Println(lResult)

	/*
		Q2. 以下の果物の価格の合計を出力するコードを書いてください。
	*/

	m := map[string]int{
		"apple":  200,
		"banana": 300,
		"grapes": 150,
		"orange": 80,
		"papaya": 500,
		"kiwi":   90,
	}

	mResult := 0

	for _, value2 := range m {
		mResult += value2
	}

	fmt.Println(mResult)
}