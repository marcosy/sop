package cli

import (
	"fmt"

	"github.com/marcosy/setop/internal/calculator"
)

const (
	opUnion        = "union"
	opIntersection = "intersection"
	opDifference   = "difference"
)

func New(opts ...Opt) *Cli {
	defaultCLI := &Cli{
		printf: func(s string, i ...interface{}) {
			_, _ = fmt.Printf(s, i...)
		},
		newCalculator: calculator.New,
	}

	for _, opt := range opts {
		opt(defaultCLI)
	}

	return defaultCLI
}

type Cli struct {
	printf        printer
	newCalculator calculatorConstructor
}

func (c *Cli) Run(args []string) int {
	if len(args) == 0 {
		c.showHelp()
		return 0
	}

	if len(args) != 3 {
		c.showHelp()
		return 1
	}

	calc, err := c.newCalculator(args[1], args[2])
	if err != nil {
		c.printf("Unable to perform operation: %v", err)
		return 2
	}

	operation := args[0]
	switch operation {
	case opUnion:
		c.printf(calc.Union())
	case opIntersection:
		c.printf(calc.Intersection())
	case opDifference:
		c.printf(calc.Difference())
	default:
		c.showHelp()
		return 3
	}

	c.printf("\n")
	return 0
}

func (c *Cli) showHelp() {
	c.printf("Usage:\n\tsetop <operation> <filepath 1> <filepath 2>\n\n")
	c.printf("<operation> must be one of: union, intersection, difference\n\n")
	c.printf("Example: setop union file1.txt file2.txt\n")
}
