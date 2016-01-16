package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestEvaluate3D(t *testing.T) {
	interpolant, _ := New(3, trainPoints, trainValues)
	values := interpolant.Evaluate(testPoints)
	assert.EqualWithin(values, testValues, 1e-15, t)
}
