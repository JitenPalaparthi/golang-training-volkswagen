package main

import "fmt"

func main() {
	c1 := new(Calc)
	sum := c1.Add(12, 123, 134, 124, 12, 4, 43, 543, 2, 3).Add(12, -12, -123, -123).Add(100, 1000).Mul(10).Mul(2).Get()
	fmt.Println("Sum:", sum)
}

type ICalc interface {
	Add(num ...int) ICalc
	Mul(num int) ICalc
	Get() int
}

type Calc struct {
	sum int
}

func (c *Calc) Add(nums ...int) ICalc {
	for _, v := range nums {
		c.sum += v
	}
	return c
}
func (c *Calc) Mul(num int) ICalc {
	c.sum = c.sum * num
	return c
}
func (c *Calc) Get() int {
	return c.sum
}
