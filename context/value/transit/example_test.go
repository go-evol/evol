package transit

import (
	"context"
	"fmt"
)

type key string

var (
	pkey = key("p")
	ckey = key("c")
)

func ExampleWithValue() {
	pctx := context.WithValue(context.Background(), pkey, "pval")
	cctx := context.WithValue(pctx, ckey, "cval")

	fmt.Println(pctx.Value(pkey), pctx.Value(ckey))
	fmt.Println(cctx.Value(pkey), cctx.Value(ckey))

	// Output:
	// pval <nil>
	// pval cval
}
