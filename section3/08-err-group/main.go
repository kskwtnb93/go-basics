package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 800*time.Millisecond)
	defer cancel()

	// eg := new(errgroup.Group)
	// eg, ctx := errgroup.WithContext(context.Background())
	eg, ctx := errgroup.WithContext(ctx)
	// s := []string{"task1", "fake1", "task2", "fake2", "task3"}
	s := []string{"task1", "task2", "task3", "task4"}
	for _, v := range s {
		task := v
		eg.Go(func() error {
			return doTask(ctx, task)
		})
	}
	if err := eg.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println("Finish!")
}

func doTask(ctx context.Context, task string) error {
	// if task == "fake1" || task == "fake2" {
	// 	return fmt.Errorf("%v failed", task)
	// }
	// fmt.Printf("Task %v completed\n", task)

	var t *time.Ticker

	// switch task {
	// case "fake1":
	// 	t = time.NewTicker(500 * time.Millisecond)
	// case "fake2":
	// 	t = time.NewTicker(700 * time.Millisecond)
	// default:
	// 	t = time.NewTicker(1000 * time.Millisecond)
	// }
	switch task {
	case "task1":
		t = time.NewTicker(500 * time.Millisecond)
	case "task2":
		t = time.NewTicker(700 * time.Millisecond)
	default:
		t = time.NewTicker(1000 * time.Millisecond)
	}

	select {
	case <-ctx.Done():
		fmt.Printf("%v cancelled: %v\n", task, ctx.Err())
		return ctx.Err()
	case <-t.C:
		t.Stop()
		// if task == "fake1" || task == "fake2" {
		// 	return fmt.Errorf("%v process failed", task)
		// }
		fmt.Printf("Task %v completed\n", task)
	}

	return nil
}
