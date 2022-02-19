package integers

func Sum(x, y int) int {
	return x + y
}

func ExampleAdd() int {
	sum := Sum(1, 2)
	return sum
	//output 3
}

func Subtraction(x, y int) int {
	return x - y
}

func ExampleSubtraction() int {
	difference := Subtraction(7, 5)
	return difference
}

func Multipliation(x, y int) int {
	return x * y
}

func ExampleMultiplication() int {
	product := Multipliation(5, 4)
	return product
}

func Division(x, y int) int {
	return x / y
}

func ExampleDivision() int {
	quotient := Division(10, 2)
	return quotient
}
