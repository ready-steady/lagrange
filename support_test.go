package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestGrid(t *testing.T) {
	product := newGrid(3, 4*3*2)

	product.next([]float64{1, 2, 3, 4})
	product.next([]float64{10, 20, 30})
	product.next([]float64{100, 200})

	assert.Equal(product.values, []float64{
		1, 10, 100,
		2, 10, 100,
		3, 10, 100,
		4, 10, 100,
		1, 20, 100,
		2, 20, 100,
		3, 20, 100,
		4, 20, 100,
		1, 30, 100,
		2, 30, 100,
		3, 30, 100,
		4, 30, 100,
		1, 10, 200,
		2, 10, 200,
		3, 10, 200,
		4, 10, 200,
		1, 20, 200,
		2, 20, 200,
		3, 20, 200,
		4, 20, 200,
		1, 30, 200,
		2, 30, 200,
		3, 30, 200,
		4, 30, 200,
	}, t)
}

func TestWeight(t *testing.T) {
	product := newWeight(4 * 3 * 2)

	product.next([]float64{1, 2, 3, 4})
	product.next([]float64{10, 20, 30})
	product.next([]float64{100, 200})

	assert.Equal(product.values, []float64{
		1 * 10 * 100,
		2 * 10 * 100,
		3 * 10 * 100,
		4 * 10 * 100,
		1 * 20 * 100,
		2 * 20 * 100,
		3 * 20 * 100,
		4 * 20 * 100,
		1 * 30 * 100,
		2 * 30 * 100,
		3 * 30 * 100,
		4 * 30 * 100,
		1 * 10 * 200,
		2 * 10 * 200,
		3 * 10 * 200,
		4 * 10 * 200,
		1 * 20 * 200,
		2 * 20 * 200,
		3 * 20 * 200,
		4 * 20 * 200,
		1 * 30 * 200,
		2 * 30 * 200,
		3 * 30 * 200,
		4 * 30 * 200,
	}, t)
}
