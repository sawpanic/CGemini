package scoring

import (
	"testing"
)

func TestGramSchmidtOrthogonalization(t *testing.T) {
	vectors := [][]float64{
		{1, 1},
		{1, 2},
	}

	orthogonalized, err := gramSchmidtOrthogonalization(vectors)
	if err != nil {
		t.Fatalf("gramSchmidtOrthogonalization failed: %v", err)
	}

	if len(orthogonalized) != 2 {
		t.Fatalf("expected 2 orthogonalized vectors, got %d", len(orthogonalized))
	}

	// Check for orthogonality (dot product should be close to 0)
	dotProduct := orthogonalized[0][0]*orthogonalized[1][0] + orthogonalized[0][1]*orthogonalized[1][1]
	if dotProduct > 1e-9 {
		t.Fatalf("vectors are not orthogonal, dot product is %f", dotProduct)
	}
}
