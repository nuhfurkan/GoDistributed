package representations

import (
	"fmt"
)

type FloatingPointRepresentation struct {
	Genes []float64
}

func (fp *FloatingPointRepresentation) Generate(gf Generation, payload GeneratePayload) Representation {
	return gf(fp, payload)
}

func (fp *FloatingPointRepresentation) Mutate(mf Mutation) Representation {
	return mf(fp)
}

func (fp *FloatingPointRepresentation) Show() string {
	return fmt.Sprintf("%f", fp.Genes)
}

func (fp *FloatingPointRepresentation) Get() map[string]interface{} {
	return map[string]interface{}{
		"value": fp.Genes,
	}
}