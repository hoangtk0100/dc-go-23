package main

import "testing"

func TestNumDifferentIntegers(t *testing.T) {
	testCases := []struct {
		name   string
		input  string
		expect int
	}{
		{
			name:   "a123bc34d8ef34",
			input:  "a123bc34d8ef34",
			expect: 3,
		},
		{
			name:   "A1b01c001",
			input:  "A1b01c001",
			expect: 1,
		},
		{
			name:   "93T0zmuh97Uqd9v",
			input:  "93T0zmuh97Uqd9v",
			expect: 4,
		},
		{
			name:   "23Wt91",
			input:  "23Wt91",
			expect: 2,
		},
		{
			name:   "Gzp97l7Rcnk",
			input:  "Gzp97l7Rcnk",
			expect: 2,
		},
		{
			name:   "Sw17LBthDge657EY7F",
			input:  "Sw17LBthDge657EY7F",
			expect: 3,
		},
		{
			name:   "jUngthdfC168m8Z",
			input:  "jUngthdfC168m8Z",
			expect: 2,
		},
		{
			name:   "P36kB7b796F1P",
			input:  "P36kB7b796F1P",
			expect: 4,
		},
		{
			name:   "2h24xXhX8LVnT2o3i",
			input:  "2h24xXhX8LVnT2o3i",
			expect: 4,
		},
		{
			name:   "69W0u969n2bUuZ036l4",
			input:  "69W0u969n2bUuZ036l4",
			expect: 6,
		},
		{
			name:   "918na87mow90ASBd4madi2nmac10",
			input:  "918na87mow90ASBd4ma90di2nmac10",
			expect: 6,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := numDifferentIntegers(tc.input)
			if got != tc.expect {
				t.Errorf("Input: %s, expect %d, but got %d", tc.input, tc.expect, got)
			}
		})
	}
}
