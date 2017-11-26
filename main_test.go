package lagrange

import (
	"testing"

	"github.com/ready-steady/assert"
)

func TestEvaluate3D(t *testing.T) {
	interpolant := New(trainPoints, trainValues)
	values := interpolant.Evaluate(testPoints)
	assert.Close(values, testValues, 1e-15, t)
}
