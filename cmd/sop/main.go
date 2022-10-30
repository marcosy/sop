package main

import (
	"os"

	"github.com/marcosy/setop/cmd/sop/cli"
)

func main() {
	os.Exit(cli.New().Run(os.Args[1:]))
}
