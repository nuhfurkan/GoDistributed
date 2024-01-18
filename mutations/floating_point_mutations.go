package mutations

import (
	representations "go-distributed/representations"
)

// To-Do
func getSmallNumber() float64 {
	return 0
}

// To-Do fix here
func UniformMutation(rep representations.Representation) representations.Representation {
	if defaultValue, ok := rep.Get()["value"].(float64); ok {
		return &representations.FloatingPointRepresentation{
			Genes: []float64{ defaultValue + getSmallNumber()},
		}
	} else {
		return nil
	}
}

func NonUniformMutation(rep representations.Representation) representations.Representation {
	// Implement non uniform mutation for real valued
	return nil
}