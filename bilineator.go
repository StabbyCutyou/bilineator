package bilineator

// Bilineator represents a set of values and the result of a function F with those
// inputs, such that it can derive the result of F over any inputs.
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

// Bilineate will let you determine the result of a function F over inputs X and Y
// based on the given inputs in the Bilineator
func (b *Bilineator) Bilineate(x float64, y float64) float64 {
	// I ran some benchmarks - I can cut out the allocations and shave 1ns per op it seems.
	// Doesn't seem worth it at the time.
	fxy1 := (((b.X2 - x) / (b.X2 - b.X1)) * b.Q11) + ((x - b.X1) / (b.X2 - b.X1) * b.Q21)
	fxy2 := (((b.X2 - x) / (b.X2 - b.X1)) * b.Q12) + ((x - b.X1) / (b.X2 - b.X1) * b.Q22)

	fxy := (((b.Y2 - y) / (b.Y2 - b.Y1)) * fxy1) + ((y - b.Y1) / (b.Y2 - b.Y1) * fxy2)

	return fxy
}
