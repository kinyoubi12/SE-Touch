package main

import
	"testing"

func Test1(t *testing.T){
	fizzbuzz(-5)
}

func Test2(t *testing.T) {
	var tests = []struct{
		input   int
		expected string
		}{
			{1, "1"},
			{3, "Fizz"},
			{5, "Buzz"},
			{15, "Fizzbuzz"},
			{-5, "Must be Positive Number"},
		}
	for _, test := range tests {
			output := fizzbuzz(test.input)
			if output != test.expected {
				t.Errorf("Fizzbuzz(%d): expected %q, output %q", test.input, test.expected, output)
			}
		}
	
	}