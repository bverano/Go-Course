//Package say lets us say hi
package say

import "fmt"

//SayHi lets us say hi and welcome someone
func SayHi(n string) string {
	return fmt.Sprint("Hi, ", n, ". Welcome!")
}
