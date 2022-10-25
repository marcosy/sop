package cli

import "github.com/marcosy/setop/internal/calculator"

func WithPrinter(p printer) opt {
	return func(c *cli) {
		c.printf = p
	}
}

func WithCalcConstructor(constructor calculatorConstructor) opt {
	return func(c *cli) {
		c.newCalculator = constructor
	}
}

type opt func(*cli)
type printer func(string, ...interface{}) (int, error)
type calculatorConstructor func(string, string) (calculator.I, error)
