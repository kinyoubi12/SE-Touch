package main

import(
	"testing"
)

func TestCal(t *testing.T){
	if cal(2) != 4 {
		t.Error("Expected 2 + 2 to equal 4")
	}
}

func TestTableCal(t *testing.T){
	var tests = []struct{
		input int
		expected int
	} {
		{2,4},
		{-1,1},
		{0,2},
		{99999,100001},
	}

	for _, test := range tests {
		if output := cal(test.input); output != test.expected {
			t.Error("Test failed: {} inputed, {} expected, received: {}", test.input, test.expected, output)
		}
	}
}