package main

import "testing"

func TestCountRectangles(t *testing.T) {
	testCases := []struct {
		name   string
		input  [][]int
		expect int
	}{
		{
			name:   "Case 01",
			input:  [][]int{},
			expect: 0,
		},
		{
			name: "Case 02",
			input: [][]int{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
			expect: 1,
		},
		{
			name: "Case 03",
			input: [][]int{
				{1, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0},
				{1, 0, 0, 1, 1, 1, 0},
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
				{0, 0, 0, 0, 0, 0, 1},
			},
			expect: 6,
		},
		{
			name: "Case 04",
			input: [][]int{
				{0, 1, 0, 1, 1, 1, 0},
				{0, 1, 0, 0, 0, 0, 0},
				{0, 1, 0, 1, 1, 0, 0},
				{0, 0, 0, 1, 1, 0, 0},
			},
			expect: 3,
		},
		{
			name: "Case 05",
			input: [][]int{
				{1, 0, 0, 0},
				{0, 0, 0, 0},
				{1, 0, 0, 1},
				{0, 1, 0, 1},
				{0, 1, 0, 0},
				{0, 1, 0, 1},
			},
			expect: 5,
		},
		{
			name: "Case 06",
			input: [][]int{
				{0, 0, 1, 0, 0},
				{0, 0, 1, 0, 1},
				{1, 0, 0, 1, 0},
				{0, 1, 0, 0, 0},
				{0, 0, 1, 1, 0},
				{0, 0, 1, 1, 0},
			},
			expect: 6,
		},
		{
			name: "Case 07",
			input: [][]int{
				{0, 0, 1, 0, 0, 1, 0, 0, 0, 0},
				{0, 0, 1, 0, 0, 0, 1, 0, 1, 1},
				{1, 0, 0, 1, 0, 0, 0, 0, 1, 1},
				{0, 1, 0, 0, 1, 0, 0, 1, 0, 0},
				{1, 0, 1, 1, 0, 1, 0, 0, 0, 0},
				{1, 0, 0, 0, 0, 1, 0, 1, 1, 0},
				{1, 0, 1, 1, 0, 1, 0, 1, 1, 0},
				{0, 0, 1, 1, 0, 0, 1, 0, 0, 0},
			},
			expect: 15,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := countRectangles(tc.input)
			if got != tc.expect {
				t.Errorf("Case: %s, expect %d, but got %d", tc.name, tc.expect, got)
			}
		})
	}
}
