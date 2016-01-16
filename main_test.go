package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestEvaluate3D(t *testing.T) {
	interpolant, _ := New(3, TRAIN_POINTS, TRAIN_VALUES)
	values := interpolant.Evaluate(TEST_POINTS)
	assert.EqualWithin(values, TEST_VALUES, 1e-15, t)
}
