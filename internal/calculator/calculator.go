package calculator

import (
	"fmt"
	"os"
	"strings"
)

func New(filepathA, filepathB, separator string) (I, error) {
	t := &T{
		separator: separator,
	}

	setA, err := os.ReadFile(filepathA)
	if err != nil {
		return nil, fmt.Errorf("unable to read first file: %w", err)
	}

	setB, err := os.ReadFile(filepathB)
	if err != nil {
		return nil, fmt.Errorf("unable to read second file: %w", err)
	}

	t.setA = string(setA)
	t.setB = string(setB)

	return t, nil
}

type I interface {
	Union() string
	Intersection() string
	Difference() string
}

type T struct {
	setA      string
	setB      string
	separator string
}

func (t *T) Union() string {
	elementsA := strings.Split(t.setA, t.separator)
	elementsB := strings.Split(t.setB, t.separator)

	union := make(map[string]struct{})

	for _, e := range elementsA {
		union[e] = struct{}{}
	}

	for _, e := range elementsB {
		union[e] = struct{}{}
	}

	return t.format(union)
}

func (t *T) Intersection() string {
	elementsA := strings.Split(t.setA, t.separator)
	elementsB := strings.Split(t.setB, t.separator)

	lookupA := make(map[string]struct{})
	for _, e := range elementsA {
		lookupA[e] = struct{}{}
	}

	intersection := make(map[string]struct{})
	for _, e := range elementsB {
		if _, ok := lookupA[e]; ok {
			intersection[e] = struct{}{}
		}
	}

	return t.format(intersection)
}

func (t *T) Difference() string {
	elementsA := strings.Split(t.setA, t.separator)
	elementsB := strings.Split(t.setB, t.separator)

	difference := make(map[string]struct{})

	for _, e := range elementsA {
		difference[e] = struct{}{}
	}

	for _, e := range elementsB {
		delete(difference, e)
	}

	return t.format(difference)
}

func (t *T) GetSetA() string {
	return t.setA
}

func (t *T) GetSetB() string {
	return t.setB
}

func (t *T) GetSeparator() string {
	return t.separator
}

func (t *T) format(set map[string]struct{}) string {
	var str string
	for k := range set {
		if k != "" {
			str += k + t.separator
		}
	}

	return strings.TrimSuffix(str, t.separator)
}
