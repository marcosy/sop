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

func TestNew(t *testing.T) {
	f1 := makeTempFile(t, "1\n2\n3\n")
	f2 := makeTempFile(t, "3\n4\n5\n")

	tests := []struct {
		name         string
		filepath1    string
		filepath2    string
		expContent1  string
		expContent2  string
		expSeparator string
		expError     string
	}{
		{
			name:         "When both files can be read, it succeeds",
			filepath1:    f1.Name(),
			filepath2:    f2.Name(),
			expContent1:  "1\n2\n3\n",
			expContent2:  "3\n4\n5\n",
			expSeparator: "\n",
		},
		{
			name:      "When 1st file cannot be read, show proper error",
			filepath1: "wrong-file.txt",
			filepath2: f2.Name(),
			expError:  "unable to read first file: open wrong-file.txt: no such file or directory",
		},
		{
			name:      "When 2nd file cannot be read, show proper error",
			filepath1: f1.Name(),
			filepath2: "wrong-file.txt",
			expError:  "unable to read second file: open wrong-file.txt: no such file or directory",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c, err := newCalculator(t, test.filepath1, test.filepath2)
			if test.expError != "" {
				require.EqualError(t, err, test.expError)
				return
			}

			require.NoError(t, err, "unable to create new union")
			require.Equal(t, test.expContent1, c.GetSet1())
			require.Equal(t, test.expContent2, c.GetSet2())
			require.Equal(t, test.expSeparator, c.GetSeparator())
		})
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		name   string
		set1   string
		set2   string
		expSet string
	}{
		{
			name:   "Union of all different elements",
			set1:   "1\n2\n3",
			set2:   "4\n5\n6",
			expSet: "1\n2\n3\n4\n5\n6",
		},
		{
			name:   "Union with some repeated elements",
			set1:   "1\n2\n3",
			set2:   "3\n4\n5",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union with trailing separator",
			set1:   "1\n2\n3\n",
			set2:   "3\n4\n5\n",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union with trailing separators",
			set1:   "1\n2\n3\n\n",
			set2:   "3\n4\n5\n\n",
			expSet: "1\n2\n3\n4\n5",
		},
		{
			name:   "Union of set1 with empty set",
			set1:   "1\n2\n3",
			set2:   "",
			expSet: "1\n2\n3",
		},
		{
			name:   "Union of empty set with set2",
			set1:   "",
			set2:   "1\n2\n3",
			expSet: "1\n2\n3",
		},
		{
			name:   "Union of empty sets",
			set1:   "",
			set2:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f1 := makeTempFile(t, test.set1)
			f2 := makeTempFile(t, test.set2)
			c, err := newCalculator(t, f1.Name(), f2.Name())
			require.NoError(t, err)

			actResult := c.Union()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		set1   string
		set2   string
		expSet string
	}{
		{
			name:   "Intersection of all different elements",
			set1:   "1\n2\n3",
			set2:   "4\n5\n6",
			expSet: "",
		},
		{
			name:   "Intersection with some repeated elements",
			set1:   "1\n2\n3\n4",
			set2:   "3\n4\n5",
			expSet: "3\n4",
		},
		{
			name:   "Intersection with trailing separator",
			set1:   "1\n2\n3\n4",
			set2:   "3\n4\n5\n",
			expSet: "3\n4",
		},
		{
			name:   "Intersection with trailing separators",
			set1:   "1\n2\n3\n4\n\n",
			set2:   "3\n4\n5\n\n",
			expSet: "3\n4",
		},
		{
			name:   "Intersection of set1 with empty set",
			set1:   "1\n2\n3",
			set2:   "",
			expSet: "",
		},
		{
			name:   "Intersection of empty set with set2",
			set1:   "",
			set2:   "1\n2\n3",
			expSet: "",
		},
		{
			name:   "Intersection of empty sets",
			set1:   "",
			set2:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f1 := makeTempFile(t, test.set1)
			f2 := makeTempFile(t, test.set2)
			c, err := newCalculator(t, f1.Name(), f2.Name())
			require.NoError(t, err)

			actResult := c.Intersection()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func TestDifference(t *testing.T) {
	tests := []struct {
		name   string
		set1   string
		set2   string
		expSet string
	}{
		{
			name:   "Difference of all different elements",
			set1:   "1\n2\n3",
			set2:   "4\n5\n6",
			expSet: "1\n2\n3",
		},
		{
			name:   "Difference with some repeated elements",
			set1:   "1\n2\n3\n4",
			set2:   "3\n4\n5",
			expSet: "1\n2",
		},
		{
			name:   "Difference with trailing separator",
			set1:   "1\n2\n3\n4",
			set2:   "3\n4\n5\n",
			expSet: "1\n2",
		},
		{
			name:   "Difference with trailing separators",
			set1:   "1\n2\n3\n4\n\n",
			set2:   "3\n4\n5\n\n",
			expSet: "1\n2",
		},
		{
			name:   "Difference of set1 with empty set",
			set1:   "1\n2\n3",
			set2:   "",
			expSet: "1\n2\n3",
		},
		{
			name:   "Difference of empty set with set2",
			set1:   "",
			set2:   "1\n2\n3",
			expSet: "",
		},
		{
			name:   "Difference of empty sets",
			set1:   "",
			set2:   "",
			expSet: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f1 := makeTempFile(t, test.set1)
			f2 := makeTempFile(t, test.set2)
			c, err := newCalculator(t, f1.Name(), f2.Name())
			require.NoError(t, err)

			actResult := c.Difference()

			requireEqualSets(t, test.expSet, actResult, c.GetSeparator())
		})
	}
}

func newCalculator(t *testing.T, fpath1, fpath2 string) (*calculator.T, error) {
	op, err := calculator.New(fpath1, fpath2)
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

func requireEqualSets(t *testing.T, set1, set2, sep string) {
	slice1 := strings.Split(set1, sep)
	slice2 := strings.Split(set2, sep)

	sort.Strings(slice1)
	sort.Strings(slice2)

	require.ElementsMatch(t, slice1, slice2)
}
