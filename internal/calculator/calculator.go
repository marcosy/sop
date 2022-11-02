package calculator

import (
	"fmt"
	"os"
	"strings"
)

func New(filepath1, filepath2, separator string) (I, error) {
	t := &T{
		separator: separator,
	}

	set1, err := os.ReadFile(filepath1)
	if err != nil {
		return nil, fmt.Errorf("unable to read first file: %w", err)
	}

	set2, err := os.ReadFile(filepath2)
	if err != nil {
		return nil, fmt.Errorf("unable to read second file: %w", err)
	}

	t.set1 = string(set1)
	t.set2 = string(set2)

	return t, nil
}

type I interface {
	Union() string
	Intersection() string
	Difference() string
}

type T struct {
	set1      string
	set2      string
	separator string
}

func (t *T) Union() string {
	slice1 := strings.Split(t.set1, t.separator)
	slice2 := strings.Split(t.set2, t.separator)

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
			resultSet += k + t.separator
		}
	}

	resultSet = strings.TrimSuffix(resultSet, t.separator)

	return resultSet
}

func (t *T) Intersection() string {
	slice1 := strings.Split(t.set1, t.separator)
	slice2 := strings.Split(t.set2, t.separator)

	map1 := make(map[string]struct{})
	for _, e := range slice1 {
		map1[e] = struct{}{}
	}

	var resultSet string
	for _, e := range slice2 {
		if _, ok := map1[e]; !ok {
			continue
		}

		if e == "" {
			continue
		}

		resultSet += e + t.separator
	}

	resultSet = strings.TrimSuffix(resultSet, t.separator)

	return resultSet
}

func (t *T) Difference() string {
	slice1 := strings.Split(t.set1, t.separator)
	slice2 := strings.Split(t.set2, t.separator)

	resultMap := make(map[string]struct{})
	for _, e := range slice1 {
		resultMap[e] = struct{}{}
	}

	for _, e := range slice2 {
		delete(resultMap, e)
	}

	var resultSet string
	for k := range resultMap {
		if k != "" {
			resultSet += k + t.separator
		}
	}

	resultSet = strings.TrimSuffix(resultSet, t.separator)

	return resultSet
}

func (t *T) GetSet1() string {
	return t.set1
}

func (t *T) GetSet2() string {
	return t.set2
}

func (t *T) GetSeparator() string {
	return t.separator
}
