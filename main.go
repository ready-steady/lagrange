// Lagrange provides algorithms for working with multivariate Lagrange
// polynomials.
//
// https://en.wikipedia.org/wiki/Lagrange_polynomial
package lagrange

// Interpolant is a Lagrange interpolant.
type Interpolant struct {
	nd     uint
	nn     uint
	grid   [][]float64
	values []float64
}

// New creates an interpolant for a given set of multidimensional grid and the
// corresponding scalar values of the target function.
func New(grid [][]float64, values []float64) *Interpolant {
	return &Interpolant{
		nd:     uint(len(grid)),
		nn:     uint(len(values)),
		grid:   grid,
		values: values,
	}
}

// Evaluate computes the values of the interpolant at set of multidimensional
// points.
func (self *Interpolant) Evaluate(points []float64) []float64 {
	nd, nn := self.nd, self.nn
	grid, values := self.grid, self.values
	np := uint(len(points)) / nd
	result := make([]float64, np)
	for i := uint(0); i < np; i++ {
		product := newProduct(nn)
		for j := uint(0); j < nd; j++ {
			product.next(lagrange(grid[j], points[i*nd+j]))
		}
		result[i] = dot(product.weights, values)
	}
	return result
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
