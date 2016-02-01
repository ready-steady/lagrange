package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestProduct(t *testing.T) {
	weights := make([]float64, 24)
	for i := range weights {
		weights[i] = 42.0
	}

	product := &product{}
	product.next(0, 4, []float64{2.0, 3.0, 5.0, 7.0}, weights)
	product.next(1, 3, []float64{11.0, 13.0, 17.0}, weights)
	product.next(2, 2, []float64{19.0, 21.0}, weights)

	assert.Equal(weights, []float64{
		418.0, 627.0, 1045.0, 1463.0, 494.0, 741.0, 1235.0, 1729.0, 646.0, 969.0, 1615.0, 2261.0,
		462.0, 693.0, 1155.0, 1617.0, 546.0, 819.0, 1365.0, 1911.0, 714.0, 1071.0, 1785.0, 2499.0,
	}, t)
}
