package representations

import (
	"fmt"
)

type BinaryRepresentation struct {
	Genes	[]int
	Length	int
}

func (b *BinaryRepresentation) Generate(gf Generation, payload GeneratePayload) Representation {
	return gf(b, payload)
}

func (b *BinaryRepresentation) Mutate(mf Mutation) Representation {
	return mf(b)
}

func (b *BinaryRepresentation) Show() string {
	return fmt.Sprintf("%v", b.Genes)
}

// value is always the content of the representation
func (b *BinaryRepresentation) Get() map[string]interface{} {
	return map[string]interface{}{
		"value": b.Genes,
	}
}