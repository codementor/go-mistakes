package main

import (
	"fmt"
)

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		i := i
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
}