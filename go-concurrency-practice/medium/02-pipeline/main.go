package main

import (
	"context"
	"fmt"
)

func Generate(ctx context.Context, n int) <-chan int {
	out := make(chan int)
	// TODO: implement; close out when done.
	close(out)
	_ = ctx
	_ = n
	return out
}

func Square(ctx context.Context, in <-chan int) <-chan int {
	out := make(chan int)
	// TODO: implement; close out when done.
	close(out)
	_ = ctx
	_ = in
	return out
}

func Sum(ctx context.Context, in <-chan int) int64 {
	// TODO: implement
	_ = ctx
	_ = in
	return 0
}

func SumOfSquares(ctx context.Context, n int) int64 {
	return Sum(ctx, Square(ctx, Generate(ctx, n)))
}

func main() {
	fmt.Println(SumOfSquares(context.Background(), 10)) // expect 385
}
