package main

import (
	"fmt"
	_ "time"
)

var ( // global variable
	num int = square(4)
)

// init is a special function, for every package
// can write multiple init in the same package as well
func init() {
	println("square:", num)
}

func init() {
	println("calling init once again")
}

func init() {
	println("calling init yet again")
}

var (
	days = 30
)

const (
	DAY          = 1
	WEEK         = DAY * 7
	SECONDS      = 60
	MINUITS      = 60
	SECONDS_HOUR = MINUITS * SECONDS
	SECONDS_DAY  = 25 * SECONDS_HOUR // constantes are evaluated and kept in the binary which gets loaded into the code segment when the process is started.
	//SECONDS_MONTH = SECONDS_DAY * days // this does not work as days is avariable and can be changed during runtime, so this cannot be evaluated at compiletime
	OK    = (true && false) || (true || false) || (false && true) && !true
	VAL   = 10
	SQARE = VAL * VAL
	// SQUARE1 = squareconst() // reason is sqareconst is returnign a value which is evaluated at runtime
)

func main() {

	var r1, i1 float32 = 10.23, 13.45

	c1 := complex(r1, i1) // complex64

	c2 := complex(10.23, 13.45) // complex128

	c3 := complex128(c1) + c2

	c4 := c1 - complex64(c2)

	c5 := 10.23 + 13.45i // complex number

	c6 := c2 * c5

	c7 := c2 / c5

	fmt.Println(c1, c2, c3, c4, c5, c6, c7)

	r2 := real(c1)
	i2 := imag(c1)

	fmt.Println(r2, i2)

	//c = a+bi

}

func square(n int) int {
	return n * n
}

func squareconst() int {
	return 10 * 10
}
