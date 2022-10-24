package cli

import (
	"fmt"

	"github.com/marcosy/setop/internal/operator/union"
)

const (
	opUnion        = "union"
	opIntersection = "intersection"
	opDifference   = "difference"
)

func New(opts ...opt) *cli {
	defaultCLI := &cli{
		printf:   fmt.Printf,
		newUnion: union.New,
	}

	for _, opt := range opts {
		opt(defaultCLI)
	}

	return defaultCLI
}

type cli struct {
	printf   printer
	newUnion unionConstructor
}

func (c *cli) Run(args []string) int {
	if len(args) != 3 {
		c.showHelp()
		return 1
	}

	operation := args[0]
	switch operation {
	case opUnion:
		u, err := c.newUnion(args[1], args[2])
		if err != nil {
			c.printf("Unable to compute union: %v", err)
			return 3
		}

		c.printf(u.Do())
		c.printf("\n")

	case opIntersection:
		c.printf("Set intersection is not implemented yet\n")
		return 2
	case opDifference:
		c.printf("Set difference is not implemented yet\n")
		return 2
	default:
		c.showHelp()
		return 1
	}

	return 0
}

func (c *cli) showHelp() {
	c.printf("Usage:\n\tsetop <operation> <filepath 1> <filepath 2>\n\n")
	c.printf("<operation> must be one of: union, intersection, difference\n\n")
	c.printf("Example: setop union file1.txt file2.txt\n")
}
