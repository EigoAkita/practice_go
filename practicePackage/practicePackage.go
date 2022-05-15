package practicePackage

import (
	"context"
	"fmt"
	"time"

	"github.com/go-ini/ini"
	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
	"golang.org/x/sync/semaphore"
)

func PracticePackage() {
	PracticeStockPrice()
	return
	// PracticeIni()
	// PracticeSemaphore()
}

//----------Semaphore----------

func PracticeSemaphore() {
	ctx := context.TODO()
	go longProcess(ctx)
	go longProcess(ctx)
	go longProcess(ctx)
	time.Sleep(5 * time.Second)
}

//semaphore設定 => goroutineの実行数決定可能
//NewWeighted(1) => goroutineの実行数１　=> goroutine1実行の度処理が終了
var s *semaphore.Weighted = semaphore.NewWeighted(1)

func longProcess(ctx context.Context) {
	//goroutineの同時進行処理 => 一つに限定 => 他の処理は実行不可にする
	//goroutine処理3つ
	/*
		実行結果
		Could not get lock
		Could not get lock
		Wait ..
		Done
	*/
	//TryAcquireの使用例
	//必要な処理実行後、次の処理から該当の処理が無必要な場合
	isAcquire := s.TryAcquire(1)
	if !isAcquire {
		fmt.Println("Could not get lock")
		return
	}
	//s.Acquireで処理実行ロック => s.Release実行しない限り、次の処理に進行不可
	// if err := s.Acquire(ctx,1); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	defer s.Release(1)
	fmt.Println("Wait..")
	time.Sleep(1 * time.Second)
	fmt.Println("Done")
}

//----------ini----------

type ConfingList struct {
	Port      int
	DBname    string
	SQLDriver string
}

var Confing ConfingList

func init() {
	cfg, _ := ini.Load("config.ini")
	Confing = ConfingList{
		Port: cfg.Section("web").Key("port").MustInt(),
		//MustString() <= valueが空の場合の文字列を指定
		DBname: cfg.Section("db").Key("name").MustString("example.sql"),
		//String() <= valueが空の場合、空の文字列が代入
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func PracticeIni() {
	fmt.Printf("%T %v\n", Confing.Port, Confing.Port)
	fmt.Printf("%T %v\n", Confing.DBname, Confing.DBname)
	fmt.Printf("%T %v\n", Confing.SQLDriver, Confing.SQLDriver)
}

//----------株価分析----------

//原因不明エラーコード by official talib example code
// func PracticeStockPrice() {
// 	spy, _ := quote.NewQuoteFromYahoo("spy", "2019-04-01", "2019-04-04", quote.Daily, true)
// 	fmt.Print(spy.CSV())
// 	rsi2 := talib.Rsi(spy.Close, 2)
// 	fmt.Println(rsi2)
// }

//和哉様ご回答コード by Udemy

func PracticeStockPrice() {
    spy, _ := quote.NewQuoteFromCoinbase("BTC-USD", "2021-04-01", "2021-05-01", quote.Daily)
    fmt.Print(spy.CSV())
    rsi2 := talib.Rsi(spy.Close, 2)
    fmt.Println(rsi2)
	mva := talib.Ema(spy.Close,14)
	fmt.Println(mva)
}
