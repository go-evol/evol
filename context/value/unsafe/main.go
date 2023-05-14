package main

import (
	"context"
	"sync"
)

type key struct{}

var userKey key

func main() {
	userVal := map[int]bool{}
	ctx := context.WithValue(context.Background(), userKey, userVal)
	wg := &sync.WaitGroup{}
	cnt := 2
	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutateVal(ctx)
		}()
	}
	wg.Wait()
}

func mutateVal(ctx context.Context) {
	for i := 0; i < 100; i++ {
		userVal := ctx.Value(userKey).(map[int]bool)
		userVal[1] = true
	}
}
