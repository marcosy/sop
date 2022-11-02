package cli

import (
	"flag"
	"fmt"

	"github.com/marcosy/sop/internal/calculator"
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

func (c *Cli) Run() int {
	separator := flag.String("s", "\n", "String used as element separator")
	flag.Usage = c.showHelp
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		c.showHelp()
		return 0
	}

	if len(args) != 3 {
		c.showHelp()
		return 1
	}

	calc, err := c.newCalculator(args[1], args[2], *separator)
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
	c.printf("sop - A command line tool to perform set operations with files\n\n")
	c.printf("Usage:\tsop [options] <operation> <filepath A> <filepath B>\n\n")

	c.printf("operation:\n")
	c.printf("  union\n")
	c.printf("\tPrint elements that are in file A or file B\n")
	c.printf("  intersection\n")
	c.printf("\tPrint elements that are in file A and file B\n")
	c.printf("  difference\n")
	c.printf("\tPrint elements of file A that do not exist in file B\n")

	c.printf("\n")

	c.printf("options:\n")
	flag.PrintDefaults()

	c.printf("\n")

	c.printf("Examples:\n")
	c.printf("  sop union fileA.txt fileB.txt\n")
	c.printf("  sop -s \",\" union fileA.csv fileB.csv\n")
}
