package main

import "fmt"

func main() {

	a, b := 10, 20

	c := (a * a) + (b * b) + 2*a*b

	fmt.Println("a+b square:", c)

	d := c >= 100

	fmt.Println("d:", d)

	e := a > 100 && d
	fmt.Println("e:", e)

	f := a / b

	g := a % b

	h := float32((a + b) / 3)

	fmt.Println(f, g, h)
}

// ()
// * ,/ ,%
// + , -
//

// expression vs statement
// arthimetic
// logical
// comparision

/*

Precedence Level	Operators	Associativity
1 (highest)	() (parentheses), [ ] (indexing), . (selector), ->, ++, --	Left-to-right
2	*, /, %, <<, >>, &, &^	Left-to-right
3	+, -, `	, ^`
4	==, !=, <, <=, >, >=	Left-to-right
5	&&	|| Left-to-right
6 (lowest)	`
*/
