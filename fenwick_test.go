package fenwick_test

import (
	"math/rand"
	"testing"

	"github.com/volkhin/go-fenwick"
)

func TestBasic(t *testing.T) {
	f := fenwick.New(10)
	if err := f.Add(15, 1); err == nil {
		t.Errorf("Error! %v", err)
	}
	value := f.Sum(0, 9)
	if value != 0 {
		t.Errorf("Error: got %v; expected 0, nil", value)
	}
	if err := f.Add(3, 42); err != nil {
		t.Errorf("Error! %v", err)
	}
	value = f.Sum(0, 9)
	if value != 42 {
		t.Errorf("Error: got %v; expected 42", value)
	}
}

func TestCorrectness(t *testing.T) {
	rawValues := make([]int, 100)
	f := fenwick.New(100)
	for iteration := 0; iteration < 1000; iteration++ {
		op := rand.Intn(2)
		switch op {
		case 0: // add
			pos := rand.Intn(len(rawValues))
			value := rand.Intn(1000)
			rawValues[pos] += value
			f.Add(pos, value)
		case 1: // sum
			begin := rand.Intn(len(rawValues))
			end := rand.Intn(len(rawValues))
			if begin > end {
				begin, end = end, begin
			}
			expected := f.Sum(begin, end)
			actual := 0
			for i := begin; i <= end; i++ {
				actual += rawValues[i]
			}
			if actual != expected {
				t.Errorf("Actual (%d) != expected (%d)\n", actual, expected)
			}
		}
	}
}
