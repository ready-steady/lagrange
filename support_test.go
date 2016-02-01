package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestProduct(t *testing.T) {
	product := newProduct(24)

	product.next([]float64{2, 3, 5, 7})
	product.next([]float64{11, 13, 17})
	product.next([]float64{19, 23})

	assert.Equal(product.weights, []float64{
		2 * 11 * 19,
		3 * 11 * 19,
		5 * 11 * 19,
		7 * 11 * 19,
		2 * 13 * 19,
		3 * 13 * 19,
		5 * 13 * 19,
		7 * 13 * 19,
		2 * 17 * 19,
		3 * 17 * 19,
		5 * 17 * 19,
		7 * 17 * 19,
		2 * 11 * 23,
		3 * 11 * 23,
		5 * 11 * 23,
		7 * 11 * 23,
		2 * 13 * 23,
		3 * 13 * 23,
		5 * 13 * 23,
		7 * 13 * 23,
		2 * 17 * 23,
		3 * 17 * 23,
		5 * 17 * 23,
		7 * 17 * 23,
	}, t)
}
