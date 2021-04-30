package main

import (
	"fmt"
	"math"
)

// complex number with real and imaginary parts
type cmplx struct {
	real float64
	img  float64
}

// function that removes appearing 0.0000000000001's etc. in float64 after operations
func RoundFloat(val float64) float64 {
	val = val * 100000
	val = math.Round(val)
	val = val / 100000
	return val
}

// printing complex number and operation it was obtained after
func Print(op string, r cmplx) string {
	return fmt.Sprintf("%v: real: %.4f, imaginary: %.4f\n", op, r.real, r.img)
}

// sum of 2 complex numbers
func (c *cmplx) Sum(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real + s.real), img: RoundFloat(f.img + s.img)}
}

// substraction of 2 complex numbers
func (c *cmplx) Sub(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real - s.real), img: RoundFloat(f.img - s.img)}
}

// multiplication of 2 complex numbers
func (c *cmplx) Mult(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real*s.real - f.img*s.img), img: RoundFloat(f.real*s.img + f.img*s.real)}
}

// division of 2 complex numbers
func (c *cmplx) Div(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat((f.real*s.real + f.img*s.img) / (s.real*s.real + s.img*s.img)), img: RoundFloat((s.real*f.img - s.img*f.real) / (s.real*s.real + s.img*s.img))}

}

// a few operations
func main() {
	first := cmplx{real: 5.1, img: 3}
	second := cmplx{real: 7.4, img: 4}
	third := cmplx{real: 4.5, img: 2.0}

	var result cmplx

	result = result.Sum(first, third)
	fmt.Println(Print("Sum", result))

	result = result.Mult(second, third)
	fmt.Println(Print("Multiplication", result))

}
