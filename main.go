// Lagrange provides algorithms for working with multivariate Lagrange
// polynomials.
//
// https://en.wikipedia.org/wiki/Lagrange_polynomial
package lagrange

import (
	"errors"
)

// Interpolant is a Lagrange interpolant.
type Interpolant struct {
	nd uint
	xs []float64
	ys []float64
}

// New creates an interpolant for a given set of multidimensional points and the
// corresponding scalar values of the target function.
func New(dimensions uint, points, values []float64) (*Interpolant, error) {
	if uint(len(points)) != uint(len(values))*dimensions {
		return nil, errors.New("the numbers of points and values should be equal")
	}
	return &Interpolant{
		nd: dimensions,
		xs: points,
		ys: values,
	}, nil
}

// Evaluate computes the values of the interpolant at set of multidimensional
// points.
func (self *Interpolant) Evaluate(_ []float64) []float64 {
	return nil
}
