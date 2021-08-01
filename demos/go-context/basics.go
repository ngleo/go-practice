package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	// any standard input will cancel the task
	go func() {
		s := bufio.NewScanner(os.Stdin)
		s.Scan()
		fmt.Println("Cancelling task")
		cancel()
	}()

	completeTask(ctx)
}

func completeTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed")
	case <-ctx.Done():
		fmt.Print("Task cancelled")
	}
}
