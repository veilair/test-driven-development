package money

import "testing"

func TestFrancMultiplication(t *testing.T) {
	five := Franc{5}
	if !five.Times(2).Equals(Franc{10}) {
		t.Errorf("Expected to be 10!")
	}
	if !five.Times(3).Equals(Franc{15}) {
		t.Errorf("Expected to be 15!")
	}
}
