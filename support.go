package lagrange

type product struct {
	index  uint
	stride uint
	repeat uint
}

func newProduct(count uint) product {
	return product{
		index:  0,
		stride: 1,
		repeat: count,
	}
}

func (self *product) next(order uint, process func(uint, uint, uint)) {
	self.repeat /= order
	stride := self.stride * order
	for i := uint(0); i < order; i++ {
		start := i * self.stride
		for j := uint(0); j < self.repeat; j++ {
			for k := start; k < start+self.stride; k++ {
				process(self.index, i, k)
			}
			start += stride
		}
	}
	self.stride = stride
	self.index++
}

type grid struct {
	product
	dimensions uint
	values     []float64
}

func newGrid(dimensions uint, count uint) grid {
	return grid{
		product:    newProduct(count),
		dimensions: dimensions,
		values:     make([]float64, dimensions*count),
	}
}

func (self *grid) next(values []float64) {
	self.product.next(uint(len(values)), func(index, order, position uint) {
		self.values[position*self.dimensions+index] = values[order]
	})
}

type weight struct {
	product
	values []float64
}

func newWeight(count uint) weight {
	values := make([]float64, count)
	for i := range values {
		values[i] = 1.0
	}
	return weight{
		product: newProduct(count),
		values:  values,
	}
}

func (self *weight) next(values []float64) {
	self.product.next(uint(len(values)), func(_, order, position uint) {
		self.values[position] *= values[order]
	})
}
