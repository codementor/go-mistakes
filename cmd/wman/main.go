package main

import (
	"fmt"
	"os"

	"github.com/codementor/go-mistakes/pkg/string"
)

func main() {
	name := os.Args[1]
	fmt.Printf("hello %s, your name backward is %q", name, string.Reverse(name))
}
