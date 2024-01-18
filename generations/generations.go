package generations

import representations "go-distributed/representations"


var GenerationFunctions = map[string] interface{} {
	"binary": map[string]interface{}{
		"random": 		RandomNBitGeneration,
		"one-point": 	OnePointCrossover,
		"n-point":		NPointCrossover,
		"uniform":		UniformCrossover,
	},
	"integer": map[string]interface{}{
		"random": 	RandomIntegerInInteval,
		
	},
	"floating_point": map[string]interface{}{
		"random": 	RandomFloatGeneration,
		
	},
	"permutation": map[string]interface{}{
		"random":	RandomPermutationWithNumbers,
	},
}

func GetGenerations(rep string) []string {
	var rep_generations []string
	if generation_list, ok := GenerationFunctions[rep].(map[string]interface{}); ok {
		for key := range generation_list {
			rep_generations = append(rep_generations, key)
		}
	}
	return rep_generations
}

func SelectGeneration(rep string, name string) func(rep representations.Representation, payload representations.GeneratePayload) representations.Representation {
	if generation_list, ok := GenerationFunctions[rep].(map[string]interface{}); ok {
		if generation_func, ok := generation_list[name].(func(rep representations.Representation, payload representations.GeneratePayload) representations.Representation); ok {
			return generation_func
		} else {
			return nil
		}
	} else {
		return nil
	}
}