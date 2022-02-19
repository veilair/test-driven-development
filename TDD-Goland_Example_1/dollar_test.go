package money

import "testing"

func TestMultiplication(t *testing.T) {
	five := Dollar{5}
	if !five.Times(2).Equals(Dollar{10}) {
		t.Errorf("Expected to be 10!")
	}
	if !five.Times(3).Equals(Dollar{15}) {
		t.Errorf("Expected to be 15!")
	}
}

func TestEquality(t *testing.T) {
	five := Dollar{5}
	if !five.Equals(Dollar{5}) {
		t.Error("Expected to be equal!")
	}
	if five.Equals(Dollar{6}) {
		t.Error("Expected to be different!")
	}
}
