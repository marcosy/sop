package calculator_test

import (
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/marcosy/sop/internal/calculator"
	"github.com/stretchr/testify/require"
)

const separator string = "\n"

func TestNew(t *testing.T) {
	fA := makeTempFile(t, "1\n2\n3\n")
	fB := makeTempFile(t, "3\n4\n5\n")

	tests := []struct {
		name         string
		filepathA    string
		filepathB    string
		expContentA  string
		expContentB  string
		expSeparator string
		expError     string
	}{
		{
			name:         "When both files can be read, it succeeds",
			filepathA:    fA.Name(),
			filepathB:    fB.Name(),
			expContentA:  "1\n2\n3\n",
			expContentB:  "3\n4\n5\n",
			expSeparator: separator,
		},
		{
			name:      "When 1st file cannot be read, show proper error",
			filepathA: "wrong-file.txt",
			filepathB: fB.Name(),
			expError:  "unable to read first file: open wrong-file.txt: no such file or directory",
		},
		{
			name:      "When 2nd file cannot be read, show proper error",
			filepathA: fA.Name(),
			filepathB: "wrong-file.txt",
			expError:  "unable to read second file: open wrong-file.txt: no such file or directory",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := newCalculator(t, test.filepathA, test.filepathB)
			if test.expError != "" {
				require.EqualError(t, err, test.expError)
				return
			}

			require.NoError(t, err, "unable to create new union")
			require.Equal(t, test.expContentA, c.GetSetA())
			require.Equal(t, test.expContentB, c.GetSetB())
			require.Equal(t, test.expSeparator, c.GetSeparator())
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name   string
		setA   string
		setB   string
		expSet string
	}{
		{
			name:   "Union of all different elements",
			setA:   "1\n2\n3",
			setB:   "4\n5\n6",
			expSet: "1\n2\n3\n4\n5\n6",
		},
		{
			name:   "Union with some repeated elements",
			setA:   "1\n2\n3",
			setB:   "3\n4\n5",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union with trailing separator",
			setA:   "1\n2\n3\n",
			setB:   "3\n4\n5\n",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union with trailing separators",
			setA:   "1\n2\n3\n\n",
			setB:   "3\n4\n5\n\n",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union of setA with empty set",
			setA:   "1\n2\n3",
			setB:   "",
			expSet: "1\n2\n3",
		},
		{
			name:   "Union of empty set with setB",
			setA:   "",
			setB:   "1\n2\n3",
			expSet: "1\n2\n3",
		},
		{
			name:   "Union of empty sets",
			setA:   "",
			setB:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fA := makeTempFile(t, test.setA)
			fB := makeTempFile(t, test.setB)
			c, err := newCalculator(t, fA.Name(), fB.Name())
			require.NoError(t, err)

			actResult := c.Union()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		setA   string
		setB   string
		expSet string
	}{
		{
			name:   "Intersection of all different elements",
			setA:   "1\n2\n3",
			setB:   "4\n5\n6",
			expSet: "",
		},
		{
			name:   "Intersection with some repeated elements",
			setA:   "1\n2\n3\n4",
			setB:   "3\n4\n5",
			expSet: "3\n4",
		},
		{
			name:   "Intersection with trailing separator",
			setA:   "1\n2\n3\n4",
			setB:   "3\n4\n5\n",
			expSet: "3\n4",
		},
		{
			name:   "Intersection with trailing separators",
			setA:   "1\n2\n3\n4\n\n",
			setB:   "3\n4\n5\n\n",
			expSet: "3\n4",
		},
		{
			name:   "Intersection of setA with empty set",
			setA:   "1\n2\n3",
			setB:   "",
			expSet: "",
		},
		{
			name:   "Intersection of empty set with setB",
			setA:   "",
			setB:   "1\n2\n3",
			expSet: "",
		},
		{
			name:   "Intersection of empty sets",
			setA:   "",
			setB:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fA := makeTempFile(t, test.setA)
			fB := makeTempFile(t, test.setB)
			c, err := newCalculator(t, fA.Name(), fB.Name())
			require.NoError(t, err)

			actResult := c.Intersection()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		setA   string
		setB   string
		expSet string
	}{
		{
			name:   "Difference of all different elements",
			setA:   "1\n2\n3",
			setB:   "4\n5\n6",
			expSet: "1\n2\n3",
		},
		{
			name:   "Difference with some repeated elements",
			setA:   "1\n2\n3\n4",
			setB:   "3\n4\n5",
			expSet: "1\n2",
		},
		{
			name:   "Difference with trailing separator",
			setA:   "1\n2\n3\n4",
			setB:   "3\n4\n5\n",
			expSet: "1\n2",
		},
		{
			name:   "Difference with trailing separators",
			setA:   "1\n2\n3\n4\n\n",
			setB:   "3\n4\n5\n\n",
			expSet: "1\n2",
		},
		{
			name:   "Difference of setA with empty set",
			setA:   "1\n2\n3",
			setB:   "",
			expSet: "1\n2\n3",
		},
		{
			name:   "Difference of empty set with setB",
			setA:   "",
			setB:   "1\n2\n3",
			expSet: "",
		},
		{
			name:   "Difference of empty sets",
			setA:   "",
			setB:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			fA := makeTempFile(t, test.setA)
			fB := makeTempFile(t, test.setB)
			c, err := newCalculator(t, fA.Name(), fB.Name())
			require.NoError(t, err)

			actResult := c.Difference()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func newCalculator(t *testing.T, fpath1, fpath2 string) (*calculator.T, error) {
	op, err := calculator.New(fpath1, fpath2, separator)
	if err != nil {
		return nil, err
	}

	u, ok := op.(*calculator.T)
	require.True(t, ok, "unable to cast operator to union")
	return u, err
}

func makeTempFile(t *testing.T, content string) *os.File {
	f, err := ioutil.TempFile("", "sop-test-file")
	require.NoError(t, err, "unable to create temp file")

	_, err = f.WriteString(content)
	require.NoError(t, err, "unable to write to temp file")

	require.NoError(t, f.Close(), "unable to close temp file")
	return f
}

func requireEqualSets(t *testing.T, setA, setB, sep string) {
	slice1 := strings.Split(setA, sep)
	slice2 := strings.Split(setB, sep)

	sort.Strings(slice1)
	sort.Strings(slice2)

	require.ElementsMatch(t, slice1, slice2)
}
