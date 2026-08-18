// Package sum provides functionality for adding
// two numbesr together.
package sum

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Add adds two numbers (integers or floats) together.
// See [addition].
//
// [addition]: https://www.mathsisfun.com/numbers/addition.html
func Add[T Number](a, b T) T {
	return a + b
}
