package representations

type GeneratePayload struct {
	Father				Representation	`json:"father,omitempty"`
	Mother				Representation	`json:"mother,omitempty"`

	// Binary Representation
	Length				int				`json:"length,omitempty"`
	NPoint				int				`json:"npoint,omitempty"`
	MPoint				int				`json:"mpoint,omitempty"`

	// Floating Point Representation
	FloatingPointMin	float64			`json:"floatingpointmin,omitempty"`
	FloatingPointMax	float64			`json:"floatingpointmax,omitempty"`

	// Integer Representation
	IntegerMax			int				`json:"integermax,omitempty"`
	IntegerMin			int				`json:"integermin,omitempty"`

	GenerationSize		int				`json:"generationsize,omitempty"`
	DesiredScore		float64			`json:"desired_score,omitempty"`
}

type Mutation func(Representation) Representation
type Generation func(Representation, GeneratePayload) Representation

type Representation interface {
	Generate(gf Generation, payload GeneratePayload) Representation
	Mutate(mf Mutation) Representation
	Show() string
	Get() map[string]interface{}
}

var Representations = map[string]Representation {
	"binary":			&BinaryRepresentation{},
	"integer":			&IntegerRepresentation{},
	"floating_point":	&FloatingPointRepresentation{},
	"permutation":		&PermutationRepresentation{},
}

func GetRepresentation(rep string) []string {
	var rep_mutations []string
	for k := range Representations {
		rep_mutations = append(rep_mutations, k)
	}
	return rep_mutations
}

func SelectRepresentation(rep string) Representation {
	if selected_rep, ok := Representations[rep]; ok {
		return selected_rep
	} else {
		return nil
	}
}