package generations

import (
	representations "go-distributed/representations"
	"math/rand"
)

func RandomNBitGeneration(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	genes := make([]int, payload.Length)
	for i := range genes {
		genes[i] = rand.Intn(2)
	}
	return &representations.BinaryRepresentation{Genes: genes}
}

func OnePointCrossover(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	if genesfather, ok := payload.Father.Get()["value"].([]int); ok {
		if genesmother, ok := payload.Mother.Get()["value"].([]int); ok {
			genes := append(genesfather[:payload.NPoint], genesmother[payload.NPoint:]...)
	
			return &representations.BinaryRepresentation{
				Genes: genes,
			}
		}
	}
	return nil
}

func NPointCrossover(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	if genesfather, ok := payload.Father.Get()["value"].([]int); ok {
		if genesmother, ok := payload.Mother.Get()["value"].([]int); ok {
			genes := append(
				append(genesfather[:payload.NPoint], genesmother[payload.NPoint:payload.MPoint]...), genesfather[payload.MPoint:]...
			)
	
			return &representations.BinaryRepresentation{
				Genes: genes,
			}
		}
	}
	return nil
}
 
func UniformCrossover(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	if genesfather, ok := payload.Father.Get()["value"].([]int); ok {
		if genesmother, ok := payload.Mother.Get()["value"].([]int); ok {
			genes := make([]int, len(genesfather))
	
			for i := 0; i < len(genesfather); i++ {
				if rand.Float64() < 0.5 {
					genes[i] = genesfather[i]
				} else {
					genes[i] = genesmother[i]
				}
			}

			return &representations.BinaryRepresentation{
				Genes: genes,
			}
		}
	}
	return nil
}