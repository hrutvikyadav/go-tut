package main

// Fibonacci is a function that returns
// a function that returns an int.
func Fibonacci() func() int {
	type LastTwo struct {
		s_last, last int
	}
	var lastTwo LastTwo = LastTwo{s_last: 0, last: 1}

	// Closures are functions that refer to variables from their "enclosing" environment, and may read or modify them, and are bound to them
	// Fibonacci() returns a closure that is bound to lastTwo which is declared in its enclosing environment
	return func() int {
		sl := lastTwo.s_last
		l := lastTwo.last

		lastTwo.s_last = l
		lastTwo.last = l + sl

		return sl
	}
}
