/*
Common denominators

You will have a list of rationals in the form
{ {numer_1, denom_1} , ... {numer_n, denom_n} }
where all numbers are positive ints.

You have to produce a result in the form
(N_1, D) ... (N_n, D)
depending on the language (See Example tests)
in which D is as small as possible and

N_1/D == numer_1/denom_1 ... N_n/D == numer_n,/denom_n.
Example:

convertFracs [(1, 2), (1, 3), (1, 4)] `shouldBe` [(6, 12), (4, 12), (3, 12)]
Note:
Due to the fact that first translations were written long ago - more than 4
years - these translations have only irreducible fractions. Newer translations
have some reducible fractions. To be on the safe side it is better to do a bit
more work by simplifying fractions even if they don't have to be.
*/
package main

import "fmt"

func ConvertFracts(a [][]int) string {
	den_lcm := 1
	for _, val := range a {
		f := gcd(val[0], val[1])
		val[0], val[1] = val[0]/f, val[1]/f
		den_lcm = lcm(val[1], den_lcm)
	}
	out := ""
	for _, val := range a {
		out += fmt.Sprintf("(%v,%v)", den_lcm/val[1]*val[0], den_lcm)
	}
	return out
}

func gcd(a, b int) int {
	if b != 0 {
		return gcd(b, a%b)
	}
	return a
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func main() {
	fmt.Println(ConvertFracts([][]int{{69, 130}, {87, 1310}, {30, 40}}))
}
