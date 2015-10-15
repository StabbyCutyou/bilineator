package bilineator

import (
	"log"
	"testing"
)

func TestBilineator(t *testing.T) {
	b := &Bilineator{
		X1: 5.0,
		X2: 10.0,

		Y1: 6.0,
		Y2: 11.0,

		Q11: 20.0,
		Q12: 50.0,
		Q21: 2.0,
		Q22: 40.0,
	}

	log.Println(b.Bilineate(5.0, 6.0))
	log.Println(b.Bilineate(5.0, 11.0))
	log.Println(b.Bilineate(10.0, 6.0))
	log.Println(b.Bilineate(10.0, 11.0))
}
