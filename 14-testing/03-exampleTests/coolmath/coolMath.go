//Package coolMath helps us to proof that you know how to sum numbers
package coolmath

//Sum sums an unlimited ammount of integer parameters
func Sum(xi ...int) int {
	var s int
	for _, v := range xi {
		s += v
	}
	return s
}
