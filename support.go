package lagrange

type product struct {
	stride  uint
	repeat  uint
	weights []float64
}

func newProduct(count uint) *product {
	weights := make([]float64, count)
	for i := range weights {
		weights[i] = 1.0
	}
	return &product{
		stride:  1,
		repeat:  count,
		weights: weights,
	}
}

func (self *product) next(values []float64) {
	order := uint(len(values))
	self.repeat /= order
	stride := self.stride * order
	for i := uint(0); i < order; i++ {
		start := i * self.stride
		for j := uint(0); j < self.repeat; j++ {
			for k := start; k < start+self.stride; k++ {
				self.weights[k] *= values[i]
			}
			start += stride
		}
	}
	self.stride = stride
}
