package main

import (
	"os"

	"github.com/marcosy/setop/cmd/setop/cli"
)

func main() {
	os.Exit(cli.New().Run(os.Args[1:]))
}
