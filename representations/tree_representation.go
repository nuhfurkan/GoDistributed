package representations

import (
	"fmt"
)

type TreeFunction func(TreeRepresentation) interface{}

type TreeRepresentation struct {
	RightBranch		*TreeRepresentation
	LeftBranch		*TreeRepresentation
	NodeFunction	TreeFunction
}

func (t *TreeRepresentation) Generate(gf Generation, payload GeneratePayload) Representation {
	return gf(t, payload)
}

func (t *TreeRepresentation) Mutate(mf Mutation) Representation {
	return mf(t)
}

func (t *TreeRepresentation) Show() string {
	return fmt.Sprintf("%v", t)
}

// value is always the content of the representation
func (t *TreeRepresentation) Get() map[string]interface{} {
	return map[string]interface{}{
		"right":	t.RightBranch,
		"left":		t.LeftBranch,
		"func":		t.NodeFunction,
	}
}