package lagrange

type product struct {
	jump   uint
	repeat uint
	skip   uint
}

func (self *product) next(index, order uint, factors []float64, weights []float64) {
	if index == 0 {
		self.jump = 1
		self.skip = 1
		self.repeat = uint(len(weights))
		for i := range weights {
			weights[i] = 1.0
		}
	}

	self.repeat /= order
	self.skip *= order

	for j := uint(0); j < order; j++ {
		start := 0 + j*self.jump
		for k := uint(1); k <= self.repeat; k++ {
			for i := start; i < start+self.jump; i++ {
				weights[i] *= factors[j]
			}
			start += self.skip
		}
	}

	self.jump *= order
}
