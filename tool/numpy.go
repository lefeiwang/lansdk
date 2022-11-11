package tool

import (
	"gonum.org/v1/gonum/mat"
)

func ConsineV(matrix [][]float64, vector []float64) (consineV []float64) {
	consineV = make([]float64, 0)
	for _, tarVector := range matrix {
		tarVec := mat.NewVecDense(512, tarVector)
		vec := mat.NewVecDense(512, vector)
		consineV = append(consineV, mat.Dot(tarVec, vec)/(tarVec.Norm(2)*vec.Norm(2)))
	}
	return
}
