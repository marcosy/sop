package main

import (
	"os"

	"github.com/marcosy/sop/internal/cli"
)

func main() {
	os.Exit(cli.New().Run(os.Args[1:]))
}
