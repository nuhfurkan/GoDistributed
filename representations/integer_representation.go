package representations

import (
	"fmt"
)

type IntegerRepresentation struct {
	Genes []int
}

func (i *IntegerRepresentation) Generate(gf Generation, payload GeneratePayload) Representation {
	return gf(i, payload)
}

func (i *IntegerRepresentation) Mutate(mf Mutation) Representation {
	return mf(i)
}

func (i *IntegerRepresentation) Show() string {
	return fmt.Sprintf("%d", i.Genes)
}

func (i *IntegerRepresentation) Get() map[string]interface{} {
	return map[string]interface{}{
		"value": i.Genes,
	}
}