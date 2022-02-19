package money

// Franc holds money in the CHF currency
type Franc struct {
	amount int
}

// Times multiply the value of Franc by a given number and returns a new Franc
func (f *Franc) Times(multiplier int) *Franc {
	return &Franc{f.amount * multiplier}
}

// Equals compare Franc with everything else
func (f *Franc) Equals(i interface{}) bool {
	franc, ok := i.(Franc)
	return ok && f.amount == franc.amount
}
