package practicePanicAndRecover

import "fmt"

func PracticePanicAndRecover() {
	//セーブ失敗で自身のアプリが強制終了
	save()
	fmt.Println("OK")
}

//Third Party の ConnectDBを使用する想定
func thirdPartyConnectDB() {
	//DB処理でエラーが発生して強制終了の想定
	//Goではpanicの記述を推奨していない
	//なぜならpanicは自分ではどうする事も不可のエラーだから
	//なのでGoはエラーとしてしっかり認識し、エラーとして返してあげる事を推奨している
	panic("Unable to connect database")
}

//自身のアプリでセーブする処理があると想定
func save() {
	//panicの例外をキャッチして処理を実行
	/*
	panicの処理実行後、以下のメソッドが発火
	panicで発生したエラーをrecoverでキャッチして、
	その中身の例外文言をキャッチしてfmt.Printlnで結果を出力
	recoverを使用する事で、システムをrecoverし強制終了不可にする
	*/
	// recoverの処理はdeferを付与してpanic処理の最初に記述しないとエラー
	defer func ()  {
		s := recover()
		fmt.Println(s)
	}()
	thirdPartyConnectDB()
}