package bilineator

import (
	"log"
	"testing"
)

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

	log.Println(b.Q11)
	log.Println(b.Q12)
	log.Println(b.Q21)
	log.Println(b.Q22)

	log.Println(b.Bilineate(5.0, 6.0))
	log.Println(b.Bilineate(5.0, 11.0))
	log.Println(b.Bilineate(10.0, 6.0))
	log.Println(b.Bilineate(10.0, 11.0))
}
