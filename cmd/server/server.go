package main

import (
	"fmt"
	"github.com/soulmonk/cuppa-workers-authentication/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.RunServer(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
