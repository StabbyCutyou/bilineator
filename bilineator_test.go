package bilineator

import (
	"fmt"
	"math/rand"
	"testing"
)

type TestCase struct {
	X float64
	Y float64
	Q float64
}

func TestBilineator(t *testing.T) {
	f := func(x float64, y float64) float64 {
		return x + y*2
	}
	b := &Bilineator{
		X1: 5.0,
		X2: 10.0,

		Y1: 6.0,
		Y2: 11.0,

		Q11: f(5.0, 6.0),
		Q12: f(5.0, 11.0),
		Q21: f(10.0, 6.0),
		Q22: f(10.0, 11.0),
	}

	testCases := []TestCase{
		TestCase{5.0, 6.0, b.Q11},
		TestCase{5.0, 11.0, b.Q12},
		TestCase{10.0, 6.0, b.Q21},
		TestCase{10.0, 11.0, b.Q22},
	}

	for _, tc := range testCases {
		if b.Bilineate(tc.X, tc.Y) != f(tc.X, tc.Y) {
			t.Error("Did not match")
			t.Fail()
		}
	}
}

func TestRandomBilineator(t *testing.T) {
	f := func(x float64, y float64) float64 {
		return x + y*2
	}
	b := &Bilineator{
		X1: 5.0,
		X2: 10.0,

		Y1: 6.0,
		Y2: 11.0,

		Q11: f(5.0, 6.0),
		Q12: f(5.0, 11.0),
		Q21: f(10.0, 6.0),
		Q22: f(10.0, 11.0),
	}

	for i := 0; i < 100; i++ {
		x := rand.Float64()
		y := rand.Float64()
		q := f(x, y)
		p := b.Bilineate(x, y)
		if fmt.Sprintf("%.12f", p) != fmt.Sprintf("%.12f", q) {
			t.Error("Did not match")
			t.Fail()
		}
	}
}

func BenchmarkBilineator(b *testing.B) {
	f := func(x float64, y float64) float64 {
		return x + y*2
	}
	bi := &Bilineator{
		X1: 5.0,
		X2: 10.0,

		Y1: 6.0,
		Y2: 11.0,

		Q11: f(5.0, 6.0),
		Q12: f(5.0, 11.0),
		Q21: f(10.0, 6.0),
		Q22: f(10.0, 11.0),
	}

	for i := 0; i < b.N; i++ {
		bi.Bilineate(5.0, float64(i))
	}
}
