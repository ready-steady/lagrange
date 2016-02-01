package lagrange

type product struct {
	jump    uint
	skip    uint
	repeat  uint
	weights []float64
}

func newProduct(count uint) *product {
	weights := make([]float64, count)
	for i := range weights {
		weights[i] = 1.0
	}
	return &product{
		jump:    1,
		skip:    1,
		repeat:  count,
		weights: weights,
	}
}

func (self *product) next(values []float64) {
	order := uint(len(values))
	self.repeat /= order
	self.skip *= order
	for i := uint(0); i < order; i++ {
		start := i * self.jump
		for j := uint(0); j < self.repeat; j++ {
			for k := start; k < start+self.jump; k++ {
				self.weights[k] *= values[i]
			}
			start += self.skip
		}
	}
	self.jump *= order
}
