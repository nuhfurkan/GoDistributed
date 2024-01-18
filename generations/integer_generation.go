package generations

import (
	representations "go-distributed/representations"
	"math/rand"
)

// Random integer within a range created
func RandomIntegerInInteval(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	genes := make([]int, payload.Length)
	for i := range genes {
		genes[i] = rand.Intn(payload.IntegerMax-payload.IntegerMin+1) + payload.IntegerMin 
	}

	return &representations.IntegerRepresentation{
		Genes:	genes,
	}
}


func SimpleArithmaticRecombination(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	return nil
}

func SingleArithmaticRecombination(rep representations.Representation) representations.Representation {
	return nil
}

func WholeArithmaticRecombination(rep representations.Representation) representations.Representation {
	return nil
}