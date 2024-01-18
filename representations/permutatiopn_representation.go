package representations

import (
	"fmt"
)

type PermutationRepresentation struct {
	Genes []int
}

func (p *PermutationRepresentation) Generate(gf Generation, payload GeneratePayload) Representation {
	return gf(p, payload)
}

func (p *PermutationRepresentation) Mutate(mf Mutation) Representation {
	return mf(p)
}

func (p *PermutationRepresentation) Show() string {
	return fmt.Sprintf("%d", p.Genes)
}

func (p *PermutationRepresentation) Get() map[string]interface{} {
	return map[string]interface{}{
		"value": p.Genes,
	}
}