package scoring

// gramSchmidtOrthogonalization takes a slice of vectors (factors) and returns an orthogonalized set of vectors.
func gramSchmidtOrthogonalization(vectors [][]float64) ([][]float64, error) {
	if len(vectors) == 0 {
		return nil, nil
	}

	orthogonalized := make([][]float64, len(vectors))

	for i, v := range vectors {
		u := make([]float64, len(v))
		copy(u, v)

		for j := 0; j < i; j++ {
			proj := projection(u, orthogonalized[j])
			for k := range u {
				u[k] -= proj[k]
			}
		}
		orthogonalized[i] = u
	}

	return orthogonalized, nil
}

// projection calculates the projection of vector v onto vector u.
func projection(v, u []float64) []float64 {
	dotProduct := 0.0
	uSquared := 0.0

	for i := range v {
		dotProduct += v[i] * u[i]
		uSquared += u[i] * u[i]
	}

	if uSquared == 0 {
		return make([]float64, len(v))
	}

	scalar := dotProduct / uSquared
	proj := make([]float64, len(u))
	for i := range u {
		proj[i] = scalar * u[i]
	}

	return proj
}
