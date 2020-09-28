//Package mymath provides a sum solution
package mymath

//Sum is cool
func Sum(xi ...int) int {
	s := 0
	for _, v := range xi {
		s += v
	}
	return s
}
