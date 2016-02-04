// Lagrange provides algorithms for working with multivariate Lagrange
// polynomials.
//
// https://en.wikipedia.org/wiki/Lagrange_polynomial
package lagrange

// Interpolant is a Lagrange interpolant.
type Interpolant struct {
	nd     uint
	nn     uint
	grids  [][]float64
	values []float64
}

// New creates an interpolant given a set of one-dimensional grids and a set of
// values of the target function obtained at the nodes of the corresponding
// tensor-product grid.
func New(grids [][]float64, values []float64) *Interpolant {
	return &Interpolant{
		nd:     uint(len(grids)),
		nn:     uint(len(values)),
		grids:  grids,
		values: values,
	}
}

// Evaluate computes the values of the interpolant at set of multidimensional
// points.
func (self *Interpolant) Evaluate(points []float64) []float64 {
	nd, nn := self.nd, self.nn
	grids, values := self.grids, self.values
	np := uint(len(points)) / nd
	result := make([]float64, np)
	for i := uint(0); i < np; i++ {
		product := newWeight(nn)
		for j := uint(0); j < nd; j++ {
			product.next(lagrange(grids[j], points[i*nd+j]))
		}
		result[i] = dot(product.values, values)
	}
	return result
}

// Tensor constructs a tensor-product grid given a set of one-dimentional grids.
func Tensor(grids [][]float64) []float64 {
	dimensions := uint(len(grids))
	count := uint(1)
	for i := uint(0); i < dimensions; i++ {
		count *= uint(len(grids[i]))
	}
	product := newGrid(dimensions, count)
	for i := uint(0); i < dimensions; i++ {
		product.next(grids[i])
	}
	return product.values
}

func lagrange(nodes []float64, point float64) []float64 {
	nn := uint(len(nodes))
	values := make([]float64, nn)
	for i := uint(0); i < nn; i++ {
		values[i] = 1.0
		for j := uint(0); j < nn; j++ {
			if i != j {
				values[i] *= (point - nodes[j]) / (nodes[i] - nodes[j])
			}
		}
	}
	return values
}

func dot(vector1, vector2 []float64) float64 {
	nn := uint(len(vector1))
	value := 0.0
	for i := uint(0); i < nn; i++ {
		value += vector1[i] * vector2[i]
	}
	return value
}
