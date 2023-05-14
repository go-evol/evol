package main

import (
	"fmt"
	"testing"
)

type foo struct{}
type bar struct{}

func main() {
	var x interface{}
	fmt.Println(testing.AllocsPerRun(100000, func() {
		x = foo{}
		x = bar{}
	}))
	fmt.Printf("%T\n", x)
}
