package mutations

import (
	representations "go-distributed/representations"
)

// Two positions randomly selected and swapped
func SwapMutation(rep representations.Representation) representations.Representation {	
	return nil
}

// Two alleles are selected at random and the second moved next to the first, shuffling along the others to make room.
func InsertMutation(rep representations.Representation) representations.Representation {	
	return nil
}

// within a randomly selected region all the aleles scrambled
func ScrambleMutation(rep representations.Representation) representations.Representation {	
	return nil
}

// Within a randomly selected region order of aleles inversed AATTGATCAAGCATGCA
//	Selecte region											   ATTGAT
// Results													   TAGTTA
func InsertionMutation(rep representations.Representation) representations.Representation {	
	return nil
}