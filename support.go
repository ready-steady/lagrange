package lagrange

type product struct {
	index  uint
	stride uint
	repeat uint
}

type weightProduct struct {
	product
	values []float64
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

func newWeightProduct(count uint) weightProduct {
	values := make([]float64, count)
	for i := range values {
		values[i] = 1.0
	}
	return weightProduct{
		product: newProduct(count),
		values:  values,
	}
}

func (self *weightProduct) next(values []float64) {
	self.product.next(uint(len(values)), func(_, order, position uint) {
		self.values[position] *= values[order]
	})
}
