package generations

import (
	representations "go-distributed/representations"
	"math/rand"
)

// A certain number of ordered integers
func RandomPermutationWithNumbers(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	var res []int

	for i := 0; i < payload.Length; i+=1 {
		res = append(res, rand.Intn(payload.IntegerMax-payload.IntegerMin+1) + payload.IntegerMin)
	}

	return &representations.PermutationRepresentation{
		Genes:	res,
	}
}

// PartiallyMappedCrossover
func PartiallyMappedCrossover(rep representations.Representation) representations.Representation {
	return nil
}

// EdgeCrossover
func EdgeCrossover(rep representations.Representation) representations.Representation {
	return nil
}

// OrderCrossover
func OrderCrossover(rep representations.Representation) representations.Representation {
	return nil
}

// CycleCrossover
func CycleCrossover(rep representations.Representation) representations.Representation {
	return nil
}