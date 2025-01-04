package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	// // カウントが+1される。
	// wg.Add(1)
	// go func() {
	// 	// カウントが-1される。
	// 	defer wg.Done()
	// 	fmt.Println("- Goroutine involed.")
	// }()
	// // カウントが0になるまで待機。
	// wg.Wait()
	// // 起動しているgoroutineの数を取得。
	// fmt.Printf("- Num of working goroutines: %d\n", runtime.NumGoroutine())
	// fmt.Println("- Main func finish.")

	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalln("Error:", err)
	}
	defer func() {
		// if文の判定に使うerrへの代入を、if文内で同時に行うことができる。
		// ここではファイルクローズした返り値をerrに代入している。
		if err := f.Close(); err != nil {
			log.Fatalln("Error:", err)
		}
	}()
	if err := trace.Start(f); err != nil {
		log.Fatalln("Error:", err)
	}
	defer trace.Stop()
	// "main"という名前のタスクを作成。
	ctx, t := trace.NewTask(context.Background(), "main")
	defer t.End()
	fmt.Println("The number of logical CPU Cores:", runtime.NumCPU())

	// task(ctx, "Task1")
	// task(ctx, "Task2")
	// task(ctx, "Task3")

	var wg sync.WaitGroup
	wg.Add(3)
	go cTask(ctx, &wg, "Task1")
	go cTask(ctx, &wg, "Task2")
	go cTask(ctx, &wg, "Task3")
	wg.Wait()

	s := []int{1, 2, 3}
	for _, i := range s {
		wg.Add(1)
		go func(i int) {
			defer wg.Done() // 関数終了時にDoneする。
			fmt.Println(i)
		}(i) // 無名関数に引数を渡す。
	}
	wg.Wait()

	fmt.Println("Main func finish.")

	// main関数を実行後、rootに「trace.out」というファイルが作成されたら以下を実行。
	// `$ go tool trace trace.out`
}
func task(ctx context.Context, name string) {
	defer trace.StartRegion(ctx, name).End()
	time.Sleep(time.Second) // 1秒間待つ
	fmt.Println(name)
}
func cTask(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer trace.StartRegion(ctx, name).End()
	// WaitGroupのカウントが-1される。
	defer wg.Done()
	time.Sleep(time.Second) // 1秒間待つ
	fmt.Println(name)
}
