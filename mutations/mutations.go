package mutations

import (
	representations "go-distributed/representations"
)

var MutationFunctions = map[string] interface{} {
	"binary": map[string]interface{}{
		"bit_flip": 	BinaryBitFlip,
	},
	"integer": map[string]interface{}{
		"creep": CreepMutation,
		"random_reset": RandomResettingMutation,
	},
	"floating_point": map[string]interface{}{
		"uniform": 		UniformMutation,
		"non_uniform": 	NonUniformMutation,
	},
	"permutation": map[string]interface{}{
		"insert":		InsertMutation,
		"swap":			SwapMutation,
		"scramble":		ScrambleMutation,
		"insertation":	InsertionMutation,
	},
}

func GetMutations(rep string) []string {
	var rep_mutations []string
	if mutation_list, ok := MutationFunctions[rep].(map[string]interface{}); ok {
		for key := range mutation_list {
			rep_mutations = append(rep_mutations, key)
		}
	}
	return rep_mutations
}

func SelectMutation(rep string, name string) func(rep representations.Representation) representations.Representation {
	if mutation_list, ok := MutationFunctions[rep].(map[string]interface{}); ok {
		if mutation_func, ok := mutation_list[name].(func(rep representations.Representation) representations.Representation); ok {
			return mutation_func
		} else {
			return nil
		}
	} else {
		return nil
	}
}