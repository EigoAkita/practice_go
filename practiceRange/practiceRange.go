package practiceRange

import "fmt"

func PracticeRange() {
	l := []string{"python", "java", "dart"}

	//本来の配列の値取得
	for i := 0; i < len(l); i++ {
		fmt.Println(i, l[i])
	}
	//Goの場合、記述簡略 + 配列の値を取得可能
	for i, v := range l {
		fmt.Println(i, v)
	}
	//index取得無ver
	for _, v := range l {
		fmt.Println(v)
	}

	for i := range l{
		fmt.Println(i)
	}

	m := map[string]int{
		"Coffee": 120,
		"Milk":   100,
		"Candy":  30,
	}

	//本来のmapの値取得
	for key, value := range m {
		fmt.Println(key, value)
	}

	//簡略化後の値取得
	//ケース1：　keyのみ取得
	for key := range m {
		fmt.Println(key)
	}

	//ケース2：valueのみ取得
	for _,value := range m {
		fmt.Println(value)
	}
}
