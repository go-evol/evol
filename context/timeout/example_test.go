package timeout

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithTimeout() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()
	done := make(chan struct{})
	go slowBusiness(done)
	select {
	case <-ctx.Done():
		fmt.Println(ctx.Err())
	case <-done:
		fmt.Println("business end")
	}

	// output:
	// context deadline exceeded
}

func ExampleAfterFunc() {
	bsChan := make(chan struct{})
	go func() {
		slowBusiness(bsChan)
	}()

	timer := time.AfterFunc(time.Millisecond, func() {
		fmt.Println("timer timeout")
	})
	<-bsChan
	fmt.Println("business end")
	timer.Stop()

	// output:
	// timer timeout
	// business end
}

func slowBusiness(done chan<- struct{}) {
	time.Sleep(2 * time.Millisecond)
	done <- struct{}{}
}
