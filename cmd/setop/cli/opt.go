package cli

import "github.com/marcosy/setop/internal/operator"

func WithPrinter(p printer) opt {
	return func(c *cli) {
		c.printf = p
	}
}

func WithUnionConstructor(constructor unionConstructor) opt {
	return func(c *cli) {
		c.newUnion = constructor
	}
}

type opt func(*cli)
type printer func(string, ...interface{}) (int, error)
type unionConstructor func(string, string) (operator.I, error)
