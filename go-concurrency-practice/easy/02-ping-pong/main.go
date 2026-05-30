package main

import "fmt"

// PingPong runs n volleys between two goroutines and returns the ordered
// sequence of labels, e.g. n=2 -> ["ping","pong","ping","pong"].
func PingPong(n int) []string {
	// TODO: implement using two goroutines and unbuffered channels.
	_ = n
	return nil
}

func main() {
	fmt.Println(PingPong(3))
}
