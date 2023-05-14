package timeout

import (
	"context"
	"fmt"
	"time"
)

func ExampleWithDeadline() {
	ctx := context.Background()
	dlCtx, cancel := context.WithDeadline(ctx, time.Now().Add(time.Hour))
	type key string
	var k key = "abc"
	childCtx := context.WithValue(dlCtx, k, 123)
	cancel()
	err := childCtx.Err()
	fmt.Println(err)

	// output:
	// context canceled
}
