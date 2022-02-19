package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Sum(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d', got '%d'", expected, sum)
	}
}

func TestExampleSum(t *testing.T) {
	sum := ExampleAdd()
	expected := 3

	if sum != expected {
		t.Errorf("expected '%d', got '%d'", expected, sum)
	}
}

func TestSubtration(t *testing.T) {
	difference := Subtraction(10, 5)
	expected := 5

	if difference != expected {
		t.Errorf("expected '%d', got '%d'", expected, difference)
	}
}

func TestExampleSubtration(t *testing.T) {
	diff := ExampleSubtraction()
	expected := 2

	if diff != expected {
		t.Errorf("expected '%d', got '%d'", expected, diff)
	}
}

func TestMultiplication(t *testing.T) {
	prod := Multipliation(2, 2)
	expected := 4
	if prod != expected {
		t.Errorf("expected '%d', got '%d'", expected, prod)
	}
}
func TestExampleMultiplication(t *testing.T) {
	prod := ExampleMultiplication()
	expected := 20

	if prod != expected {
		t.Errorf("expected '%d', got '%d'", expected, prod)
	}
}

func TestDivision(t *testing.T) {
	div := Division(10, 5)
	expected := 2
	if div != expected {
		t.Errorf("expected '%d', got '%d'", expected, div)
	}
}

func TestExampleDivision(t *testing.T) {
	div := ExampleDivision()
	expected := 5
	if div != expected {
		t.Errorf("expected '%d', got '%d'", expected, div)
	}
}
