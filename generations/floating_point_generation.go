package generations

import (
	representations "go-distributed/representations"
	"math/rand"
)

// Float within a range created
func RandomFloatGeneration(rep representations.FloatingPointRepresentation, payload representations.GeneratePayload) representations.FloatingPointRepresentation {
	genes := make([]float64, payload.Length)
	for i := range genes {
		genes[i] = payload.FloatingPointMin + rand.Float64()*float64(payload.FloatingPointMax)
	}
	
	return representations.FloatingPointRepresentation{
		Genes: genes,
	}
}

