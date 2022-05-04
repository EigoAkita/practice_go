package practiceGoroutine

import (
	"fmt"
	"sync"
	"time"
)

// func goroutine1(s []int, c chan int) {
// 	sum := 0
// 	for _, v := range s {
// 		sum += v
// 		c <- sum
// 	}
// 	close(c)
// }

// func PracticeGoroutine() {
// 	s := []int{1, 2, 3, 4, 5}
// 	c := make(chan int, len(s))
// 	go goroutine1(s, c)
// 	for i := range c {
// 		fmt.Println(i)
// 	}
// }

// func producer(ch chan int, i int) {
// 	// Something
// 	ch <- i * 2
// }

// func consumer(ch chan int, wg *sync.WaitGroup) {
// 	for i := range ch {
// 		func () {
// 			//インナーファンクション終了後 → defer wg.Done()実行
// 			//コードの書き方としては上記の方が綺麗
// 			defer wg.Done()
// 			fmt.Println("process", i*1000)
// 		}()
// 		wg.Done()
// 	}
// }

// func PracticeGoroutine2() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)

// 	// Producer
// 	for i := 0; i < 10; i++ {
// 		wg.Add(1)
// 		go producer(ch, i)
// 	}

// 	// Consumer
// 	go consumer(ch, &wg)
// 	wg.Wait()
// 	close(ch)
// }

// func producer(first chan int) {
// 	defer close(first)
// 	for i := 0; i < 10; i++ {
// 		first <- i
// 	}
// }

// func multi2(first chan int, second chan int) {
// 	defer close(second)
// 	for i := range first {
// 		second <- i * 2
// 	}
// }

// func multi4(second chan int, third chan int) {
// 	defer close(third)
// 	for i := range second {
// 		third <- i * 4
// 	}
// }

// func PracticeGoroutine3() {
// 	first := make(chan int)
// 	second := make(chan int)
// 	third := make(chan int)

// 	go producer(first)
// 	go multi2(first, second)
// 	go multi4(second, third)
// 	for result := range third {
// 		fmt.Println(result)
// 	}
// }

// func goroutine1(ch chan string) {
// 	for {
// 		ch <- "packet form 1"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func goroutine2(ch chan string) {
// 	for {
// 		ch <- "packet form 2"
// 		time.Sleep(1 * time.Second)
// 	}
// }

// func PracticeGoroutine4() {
// 	c1 := make(chan string)
// 	c2 := make(chan string)

// 	go goroutine1(c1)
// 	go goroutine2(c2)

// 	for {
// 		select {
// 		case msg1 := <-c1:
// 			fmt.Println(msg1)
// 		case msg2 := <-c2:
// 			fmt.Println(msg2)
// 		}
// 	}
// }

// func PracticeGoroutine5() {
// 	tick := time.Tick(100 * time.Millisecond)
// 	boom := time.After(500 * time.Millisecond)
// //下記の様にOuterLoopとfor文の先頭に記述する事で、
// //break時にOuterLoopと記載すれば、for文を抜けるロジックが記述可能
// OuterLoop:
// 	for {
// 		select {
// 		case t := <-tick:
// 			fmt.Println("TICK.", t)
// 		case <-boom:
// 			fmt.Println("BOOM.")
// 			break OuterLoop
// 		default:
// 			fmt.Println("    .")
// 			time.Sleep(50 * time.Millisecond)
// 		}
// 	}
// 	fmt.Println("###########")
// }

type Counter struct {
	v   map[string]int
	mux sync.Mutex
}

func (c *Counter) Inc(key string) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.v[key]++
}

func (c *Counter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func PracticeGoroutine6() {
	c := Counter{v: make(map[string]int)}
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	go func() {
		for i := 0; i < 10; i++ {
			c.Inc("Key")
		}
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(c.Value("Key"))
}
