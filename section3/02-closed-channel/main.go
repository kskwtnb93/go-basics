package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	ch1 := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(<-ch1)
	}()
	ch1 <- 10
	close(ch1)
	// ch1が開いているか
	v, ok := <-ch1
	fmt.Printf("%v %v\n", v, ok)
	wg.Wait()

	fmt.Println("========================")

	ch2 := make(chan int, 2)
	ch2 <- 1
	ch2 <- 2
	close(ch2)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)
	v, ok = <-ch2
	fmt.Printf("%v %v\n", v, ok)

	fmt.Println("========================")

	ch3 := generateCountStream()
	for v := range ch3 {
		fmt.Println(v)
		time.Sleep(2 * time.Second)
	}

	fmt.Println("========================")

	// データの値を持たない通知専用のchannel
	nCh := make(chan struct{})
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("goroutine %v started\n", i)
			<-nCh
			fmt.Println(i)
		}(i)
	}
	time.Sleep(2 * time.Second)
	close(nCh) // 上記の <-nCh が実行される。
	fmt.Println("Unblocked by manual close.")
	wg.Wait()
	fmt.Println("Finish.")
}

// channelの生成、書き込み、close操作をカプセル化。関数の利用者は読み取り専用のchannelのみにアクセスできる。
func generateCountStream() <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i <= 5; i++ {
			ch <- i
			fmt.Println("write")
		}
	}()
	return ch
}
