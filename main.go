package main

import (
	"fmt"
	"math"
)

type cmplx struct {
	real float64
	img  float64
}

func RoundFloat(val float64) float64 {
	val = val * 100000
	val = math.Round(val)
	val = val / 100000
	return val
}

func (c *cmplx) Print(op string, r cmplx) {
	fmt.Printf("%v: real: %.4f, imaginary: %.4f\n", op, r.real, r.img)
}

func (c *cmplx) Sum(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real + s.real), img: RoundFloat(f.img + s.img)}
}

func (c *cmplx) Sub(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real - s.real), img: RoundFloat(f.img - s.img)}
}

func (c *cmplx) Mult(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat(f.real*s.real - f.img*s.img), img: RoundFloat(f.real*s.img + f.img*s.real)}
}

func (c *cmplx) Div(f cmplx, s cmplx) cmplx {
	return cmplx{real: RoundFloat((f.real*s.real + f.img*s.img) / (s.real*s.real + s.img*s.img)), img: RoundFloat((s.real*f.img - s.img*f.real) / (s.real*s.real + s.img*s.img))}

}

func main() {
	first := cmplx{real: 5.1, img: 3}
	second := cmplx{real: 7.4, img: 4}
	third := cmplx{real: 4.5, img: 2.0}

	var result cmplx

	result = result.Sum(first, third)
	result.Print("Sum", result)

	result = result.Mult(second, third)
	result.Print("Multiplication", result)

}
