package union

import (
	"fmt"
	"os"
	"strings"

	"github.com/marcosy/setop/internal/operator"
)

func New(filepath1, filepath2 string) (operator.I, error) {
	u := &Union{
		separator: "\n",
	}

	set1, err := os.ReadFile(filepath1)
	if err != nil {
		return nil, fmt.Errorf("unable to read first file: %w", err)
	}

	set2, err := os.ReadFile(filepath2)
	if err != nil {
		return nil, fmt.Errorf("unable to read second file: %w", err)
	}

	u.set1 = string(set1)
	u.set2 = string(set2)

	return u, nil
}

type Union struct {
	set1      string
	set2      string
	separator string
}

func (u *Union) Do() string {
	slice1 := strings.Split(u.set1, u.separator)
	slice2 := strings.Split(u.set2, u.separator)

	resultMap := make(map[string]struct{})
	for _, e := range slice1 {
		resultMap[e] = struct{}{}
	}

	for _, e := range slice2 {
		resultMap[e] = struct{}{}
	}

	var resultSet string
	for k := range resultMap {
		if k != "" {
			resultSet += k + u.separator
		}
	}

	resultSet = strings.TrimSuffix(resultSet, u.separator)

	return resultSet
}

func (u *Union) GetSet1() string {
	return u.set1
}

func (u *Union) GetSet2() string {
	return u.set2
}

func (u *Union) GetSeparator() string {
	return u.separator
}
