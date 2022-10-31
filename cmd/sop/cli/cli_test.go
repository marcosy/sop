package cli_test

import (
	"fmt"
	"testing"

	"github.com/marcosy/setop/cmd/sop/cli"
	"github.com/marcosy/setop/internal/calculator"
	"github.com/stretchr/testify/assert"
)

var helpMessage = `Usage:
	setop <operation> <filepath 1> <filepath 2>

<operation> must be one of: union, intersection, difference

Example: setop union file1.txt file2.txt
`

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		args          []string
		failOperation bool
		expMessage    string
		expExitCode   int
	}{
		{
			name:        "When ALL parameters are missing, shows help message and exit code equal to 0",
			args:        []string{},
			expMessage:  helpMessage,
			expExitCode: 0,
		},
		{
			name:        "When ONE parameter is missing, shows help message and exit code equal to 1",
			args:        []string{"union", "file1.txt"},
			expMessage:  helpMessage,
			expExitCode: 1,
		},
		{
			name:        "When 1st parameter is unknown, shows help message and exit code equal to 1",
			args:        []string{"unknown", "file1.txt", "file2.txt"},
			expMessage:  helpMessage,
			expExitCode: 1,
		},
		{
			name:          "When calc creation fails, exit code is non zero",
			args:          []string{"union", "file1.txt", "file2.txt"},
			failOperation: true,
			expMessage:    "Unable to perform operation: something went wrong",
			expExitCode:   3,
		},
		{
			name:        "When union is invoked, union operator is called",
			args:        []string{"union", "file1.txt", "file2.txt"},
			expMessage:  "union was called\n",
			expExitCode: 0,
		},
		{
			name:        "When intersection is invoked, intersection operator is called",
			args:        []string{"intersection", "file1.txt", "file2.txt"},
			expMessage:  "intersection was called\n",
			expExitCode: 0,
		},
		{
			name:        "When difference is invoked, difference operator is called",
			args:        []string{"difference", "file1.txt", "file2.txt"},
			expMessage:  "difference was called\n",
			expExitCode: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			r := new(recorder)
			calcConstructor := newFakeCalcConstructor
			if test.failOperation {
				calcConstructor = newFakeCalcConstructorFaulty
			}

			c := cli.New(
				cli.WithPrinter(r.printf),
				cli.WithCalcConstructor(calcConstructor),
			)
			actExitCode := c.Run(test.args)

			assert.Equal(t, test.expExitCode, actExitCode)
			assert.Equal(t, test.expMessage, r.message)
		})
	}
}

type recorder struct {
	message string
}

func (r *recorder) printf(format string, a ...interface{}) {
	r.message += fmt.Sprintf(format, a...)
}

func newFakeCalcConstructor(a, b string) (calculator.I, error) {
	return &fakeCalc{}, nil
}

func newFakeCalcConstructorFaulty(a, b string) (calculator.I, error) {
	return nil, fmt.Errorf("something went wrong")
}

type fakeCalc struct{}

func (*fakeCalc) Union() string {
	return "union was called"
}

func (*fakeCalc) Intersection() string {
	return "intersection was called"
}

func (*fakeCalc) Difference() string {
	return "difference was called"
}
