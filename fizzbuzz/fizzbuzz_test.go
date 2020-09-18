package fizzbuzz

import "testing"

func TestInput1ShouldBeDisplay1(t *testing.T) {
	v := fizzbuzz(1)
	if "1" != v {
		t.Error("fizzbuzz of 1 should be '1' but have", v)
	}
}

func TestInput1ShouldBeDisplay2(t *testing.T) {
	v := fizzbuzz(2)
	if "2" != v {
		t.Error("fizzbuzz of 2 should be '2' but have", v)
	}
}