package main

import (
	"os"

	"github.com/codementor/go-mistakes/pkg/string/cmd"
)

func main() {
	if err := cmd.NewWmanCmd().Execute(); err != nil {
		os.Exit(-1)
	}
}
