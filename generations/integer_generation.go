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
	genes := make([]int, payload.Length)
	for i := range genes {
		if father_genes, ok := payload.Father.Get()["value"].([]int); ok {
			if mother_genes, ok := payload.Mother.Get()["value"].([]int); ok {
				genes[i] = (father_genes[i] + mother_genes[i]) / 2
			}
		}
	}

	return &representations.IntegerRepresentation{
		Genes:	genes,
	}
}

/*
func SingleArithmaticRecombination(rep representations.Representation) representations.Representation {
	
	return nil
}

func WholeArithmaticRecombination(rep representations.Representation) representations.Representation {
	return nil
}
*/