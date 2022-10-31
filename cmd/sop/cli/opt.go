package cli

import "github.com/marcosy/sop/internal/calculator"

func WithPrinter(p printer) Opt {
	return func(c *Cli) {
		c.printf = p
	}
}

func WithCalcConstructor(constructor calculatorConstructor) Opt {
	return func(c *Cli) {
		c.newCalculator = constructor
	}
}

type Opt func(*Cli)
type printer func(string, ...interface{})
type calculatorConstructor func(string, string) (calculator.I, error)
