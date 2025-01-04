// channel：複数のgoroutine間でデータを送受信できるようにするもの。
// channelの受信操作が開始されるまで、書き込み操作はブロックされる。

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// ch := make(chan int)
	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	ch <- 10
	// 	time.Sleep(500 * time.Millisecond)
	// }()
	// fmt.Println(<-ch)
	// wg.Wait()

	// goroutineリーク: goroutineが稼働し続けてメモリが解放されないこと。
	ch1 := make(chan int)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	fmt.Printf("Num of working goroutines: %d\n", runtime.NumGoroutine())

	// バッファ付きのchannel
	ch2 := make(chan int, 1)
	ch2 <- 2
	ch2 <- 3 // バッファが1なのでdeadlockが発生してしまう。バッファを2以上にすると回避できた。
	fmt.Println(<-ch2)
	// ch2 <- 2　// deadlock
}
