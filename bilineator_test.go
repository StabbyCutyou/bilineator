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

	log.Println(b.Bilineate(14.0, 20.0))
	log.Println(b.Bilineate(15.0, 20.0))
	log.Println(b.Bilineate(14.0, 21.0))
	log.Println(b.Bilineate(15.0, 21.0))
}
