package practiceLogging

import (
	"fmt"
	"io"
	"log"
	"os"
)

func PracticeLogging() {
	LoggingSettings("milk.log")
	return
	firstLogging()
}

func LoggingSettings(logFile string) {
	/*
		OpenFile使用 => どのモードでファイルを開くか指定可能
		0666はパーミッション（自分,グループ,他人の合計値で算出（今回は666 => 全ユーザーがファイルREAD WRITE可能））
	*/
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//os.Stdout = 画面上の出力エラー　logfileに書込
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//フラグの設定 ログの日付、時間、ファイルパス
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	//multiLogFlagにフラグ設定追加
	log.SetOutput(multiLogFile)
}

func firstLogging() {
	_, err := os.Open("coffee")
	if err != nil {
		log.Fatalln("ERROR", err)
	}
	log.Println("logging")
	log.Printf("%T %v", "天上天下唯我独尊", "天上天下唯我独尊")
	/*
		log.Fatalln実行→処理が終了
		log.Fatalf実行→処理が終了
		使用箇所重要
	*/
	log.Fatalf("%T %v", "Z李", "Z李")
	log.Fatalln("ERROR")
	fmt.Println("処理実行中")

}
