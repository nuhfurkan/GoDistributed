package mutations

import (
	"math/rand"

	representations "go-distributed/representations"
)

func BinaryBitFlip(rep representations.Representation) representations.Representation {	
	if genes, ok := rep.Get()["value"].([]int); ok {
		mutatedGenes := make([]int, len(genes))
		copy(mutatedGenes, genes)
		index := rand.Intn(len(mutatedGenes))
		mutatedGenes[index] = 1 - mutatedGenes[index]

		return &representations.BinaryRepresentation{
			Genes: mutatedGenes,
		}
	} else {
		return nil	
	}
}