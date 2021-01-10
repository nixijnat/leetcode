package main

import (
	"reflect"
	"testing"
)

func TestFunc(t *testing.T) {
	for _, c := range []struct {
		input   [][]string
		values  []float64
		queries [][]string
		expect  []float64
	}{
		{
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
			},
			[]float64{2.0, 3.0},
			[][]string{
				[]string{"a", "c"},
				[]string{"b", "a"},
				[]string{"a", "e"},
				[]string{"a", "a"},
				[]string{"x", "x"},
			},
			[]float64{6.00000, 0.50000, -1.00000, 1.00000, -1.00000},
		},
		{
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "c"},
				[]string{"bc", "cd"},
			},
			[]float64{1.5, 2.5, 5.0},
			[][]string{
				[]string{"a", "c"},
				[]string{"c", "b"},
				[]string{"bc", "cd"},
				[]string{"cd", "bc"},
			},
			[]float64{3.75000, 0.40000, 5.00000, 0.20000},
		},
		{
			[][]string{
				[]string{"a", "b"},
			},
			[]float64{0.5},
			[][]string{
				[]string{"a", "b"},
				[]string{"b", "a"},
				[]string{"a", "c"},
				[]string{"x", "y"},
			},
			[]float64{0.50000, 2.00000, -1.00000, -1.00000},
		},
		{
			[][]string{
				[]string{"x1", "x2"},
				[]string{"x2", "x3"},
				[]string{"x3", "x4"},
				[]string{"x4", "x5"},
			},
			[]float64{3.0, 4.0, 5.0, 6.0},
			[][]string{
				[]string{"x1", "x5"},
				[]string{"x5", "x2"},
				[]string{"x2", "x4"},
				[]string{"x2", "x2"},
				[]string{"x2", "x9"},
				[]string{"x9", "x9"},
			},
			[]float64{360.0, 0.00833, 20.0, 1.0, -1.0, -1.0},
		},
	} {
		v := calcEquation(c.input, c.values, c.queries)
		t.Log(c, v, reflect.DeepEqual(v, c.expect))
	}
}
