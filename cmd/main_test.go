package main

import "testing"

func TestAddNumber(t *testing.T) {
	testTable := []struct {
		a      int
		b      int
		expect int
	}{
		{
			a:      50,
			b:      5,
			expect: 55,
		},
		{
			a:      -3,
			b:      1,
			expect: -2,
		},
		{
			a:      0,
			b:      0,
			expect: 0,
		},
	}

	for _, testCase := range testTable {
		result := AddNumber(testCase.a, testCase.b)

		if result != testCase.expect {
			t.Errorf("Unexpected result. Expect: %d; Result: %d\n", testCase.expect, result)
		}
	}
}
