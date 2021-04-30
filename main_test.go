package main

import "testing"

func TestRoundFloat(t *testing.T) {
	val := 4.200000011234
	got := RoundFloat(val)
	want := 4.2

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestSum(t *testing.T) {
	number1 := cmplx{real: 2.1, img: 3.4}
	number2 := cmplx{real: 4.6, img: 5.2}

	var got cmplx
	got = got.Sum(number1, number2)
	want := cmplx{real: 6.7000, img: 8.6000}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestSubstraction(t *testing.T) {
	number1 := cmplx{real: 5.3, img: 6.1}
	number2 := cmplx{real: 4.4, img: 5.2}

	var got cmplx
	got = got.Sub(number1, number2)
	want := cmplx{real: 0.9, img: 0.9}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestMultiplication(t *testing.T) {
	number1 := cmplx{real: 3.6, img: 2.2}
	number2 := cmplx{real: 8.5, img: 7.6}

	var got cmplx
	got = got.Mult(number1, number2)
	want := cmplx{real: 13.88, img: 46.06}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestDivision(t *testing.T) {
	number1 := cmplx{real: 6.5, img: 5.8}
	number2 := cmplx{real: 4.2, img: 3.7}

	var got cmplx
	got = got.Div(number1, number2)
	want := cmplx{real: 1.55634, img: 0.00989}

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
