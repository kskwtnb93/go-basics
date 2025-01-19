package main

import (
	"context"
	"fmt"
	"time"
)

// context: メインgoルーチンからサブgoルーチンを一斉にキャンセル

func main() {
	// var wg sync.WaitGroup
	// ctx, cancel := context.WithTimeout(context.Background(), 400*time.Millisecond)
	// defer cancel()
	// wg.Add(3)
	// go subTask(ctx, &wg, "a")
	// go subTask(ctx, &wg, "b")
	// go subTask(ctx, &wg, "c")
	// wg.Wait()

	// var wg sync.WaitGroup
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := criticalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("Critical task cancelled due to: %v\n", err)
	// 		cancel()
	// 		return
	// 	}
	// 	fmt.Println("Success", v)
	// }()
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()
	// 	v, err := normalTask(ctx)
	// 	if err != nil {
	// 		fmt.Printf("Normal task cancelled due to: %v\n", err)
	// 		return
	// 	}
	// 	fmt.Println("Success", v)
	// }()
	// wg.Wait()

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(40*time.Millisecond))
	defer cancel()
	ch := subTask2(ctx)
	v, ok := <-ch
	if ok {
		fmt.Println(v)
	}
	fmt.Println("Finish!")
}

// func subTask(ctx context.Context, wg *sync.WaitGroup, id string) {
// 	defer wg.Done()
// 	t := time.NewTicker(500 * time.Millisecond)
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println(ctx.Err())
// 			return
// 		case <-t.C:
// 			t.Stop()
// 			fmt.Println(id)
// 			return
// 		}
// 	}
// }

// func criticalTask(ctx context.Context) (string, error) {
// 	ctx, cancel := context.WithTimeout(ctx, 1100*time.Millisecond)
// 	defer cancel()
// 	t := time.NewTicker(1000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "A", nil
// }

// func normalTask(ctx context.Context) (string, error) {
// 	t := time.NewTicker(3000 * time.Millisecond)
// 	select {
// 	case <-ctx.Done():
// 		return "", ctx.Err()
// 	case <-t.C:
// 		t.Stop()
// 	}
// 	return "B", nil
// }

func subTask2(ctx context.Context) <-chan string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		deadline, ok := ctx.Deadline()
		if ok {
			if deadline.Sub(time.Now().Add(50*time.Millisecond)) < 0 {
				fmt.Println("Impossible to meet deadline.")
				return
			}
		}
		time.Sleep(30 * time.Millisecond)
		ch <- "hello"
	}()
	return ch
}
