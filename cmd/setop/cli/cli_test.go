package cli_test

import (
	"fmt"
	"testing"

	"github.com/marcosy/setop/cmd/setop/cli"
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
			name:        "When ALL parameters are missing, shows help message and exit code equal to 1",
			args:        []string{},
			expMessage:  helpMessage,
			expExitCode: 1,
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
			name:        "When union is invoked, union operator is called",
			args:        []string{"union", "file1.txt", "file2.txt"},
			expMessage:  "union was called\n",
			expExitCode: 0,
		},
		{
			name:          "When union operator fails, exit code is non zero",
			args:          []string{"union", "file1.txt", "file2.txt"},
			failOperation: true,
			expMessage:    "Unable to compute union: something went wrong",
			expExitCode:   3,
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

func (r *recorder) printf(format string, a ...interface{}) (int, error) {
	r.message += fmt.Sprintf(format, a...)
	return 0, nil
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
