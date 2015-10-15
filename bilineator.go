package bilineator

type Bilineator struct {
	X1 float64
	Y1 float64

	X2 float64
	Y2 float64

	Q11 float64
	Q12 float64
	Q21 float64
	Q22 float64
}

func (b *Bilineator) Bilineate(x float64, y float64) float64 {
	fxy1 := ((b.X2 - x/b.X2 - b.X1) * b.Q11) + ((x - b.X1/b.X2 - b.X1) * b.Q21)
	fxy2 := ((b.X2 - x/b.X2 - b.X1) * b.Q12) + ((x - b.X1/b.X2 - b.X1) * b.Q22)

	fxy := ((b.Y2 - y/b.Y2 - b.Y1) * fxy1) + ((y - b.Y1/b.Y2 - b.Y1) * fxy2)

	return fxy
}
