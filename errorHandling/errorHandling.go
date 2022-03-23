package errorHandling

import (
	"fmt"
	"log"
	"os"
)

func ErrorHandling() {
	practiceError()
}

func practiceError() {
	/*
	:= (short variable declaration)
	以下の処理に
	file,err :=
	count,err :=
	とerrという変数を二度使用してるがエラーにならない
	それはなぜか :=　の左の変数を定義する度に初期化してる為である
	（fileとcount）
	*/
	file, err := os.Open("./main.go")
	if err != nil {
		log.Fatalln("ERROR")
	}

	defer file.Close()
	data := make([]byte, 100)

	//ファイル読込
	count, err := file.Read(data)
	if err != nil {
		log.Fatalln("ERROR SECOND")
	}
	/*
	byte使用理由 => 文字列に限らずデータを扱う為 
	読み込むファイルが画像、音楽ファイルと言ったものまで様々
	それらが決して文字列で読込可能とは限らない
	その為、読み込んだ生データである、byte配列を返す
	*/
	fmt.Println(data)
	fmt.Println(count, "\n", string(data))

	//上記の処理で定義したerrをオーバーライド（err := とすると初期化不可の為エラー）
	/*
	err = os.Chdir("chocolate")
	if err != nil {
		log.Fatalln("ERROR")
	}
	*/

	//上記の処理を簡略化
	if err = os.Chdir("chocolate"); err != nil {
		log.Fatalln("ERROR")
	}

}
