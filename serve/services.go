package serve

import (
	"fmt"
	generations "go-distributed/generations"
	mpiserver "go-distributed/mpi_server"
	mutations "go-distributed/mutations"
	representation "go-distributed/representations"
	"log"
)

func GetMutations() string {
	return ""
}

func repeatExperiment(exp Experiment, reps []representation.Representation) {
	selected_rep := representation.SelectRepresentation(exp.Representation)
	if selected_rep == nil {
		log.Println("err")
		return
	} 

	selected_generation := generations.SelectGeneration(exp.Representation, exp.Generation)
	if selected_generation == nil {
		log.Println("err")
		return
	}
	generated := selected_rep.Generate(selected_generation, exp.Payload) 
	
	mut_func := mutations.SelectMutation(exp.Representation, exp.Mutation)
	if mut_func == nil {
		log.Println("err")
	}
	mutated := generated.Mutate(mut_func)
	log.Printf("Mutated: %s\n", mutated.Show())

	jobs := mpiserver.SafeStack{}

	for i:=0; i < exp.Payload.GenerationSize; i++ {
		fmt.Println("Job no", i, "created")
		jobs.Push(
			selected_rep.Generate(selected_generation, exp.Payload),
		)
	}

	best_ones := mpi.StartMaster(jobs)
	if best_ones[1].Score > exp.Payload.DesiredScore {
		fmt.Println("The best found representation is", "\n", best_ones[0].Data.Show())
		fmt.Println("The second found representation is", "\n", best_ones[1].Data.Show())
		return
	} else {
		repeatExperiment(exp, []representation.Representation{
			best_ones[0].Data,
			best_ones[1].Data,
		})
	}

}

func createExperiment(exp Experiment) {
	selected_rep := representation.SelectRepresentation(exp.Representation)
	if selected_rep == nil {
		log.Println("err")
		return
	} 

	selected_generation := generations.SelectGeneration(exp.Representation, "random")
	if selected_generation == nil {
		log.Println("err")
		return
	}
	generated := selected_rep.Generate(selected_generation, exp.Payload) 
	
	mut_func := mutations.SelectMutation(exp.Representation, exp.Mutation)
	if mut_func == nil {
		log.Println("err")
	}
	mutated := generated.Mutate(mut_func)
	log.Printf("Mutated: %s\n", mutated.Show())

	jobs := mpiserver.SafeStack{}

	for i:=0; i < exp.Payload.GenerationSize; i++ {
		fmt.Println("Job no", i, "created")
		jobs.Push(
			selected_rep.Generate(selected_generation, exp.Payload),
		)
	}

	best_ones := mpi.StartMaster(jobs)
	if best_ones[1].Score > exp.Payload.DesiredScore {
		fmt.Println("The best found representation is", "\n", best_ones[0].Data.Show())
		fmt.Println("The second found representation is", "\n", best_ones[1].Data.Show())
	} else {
		repeatExperiment(exp, []representation.Representation{
			best_ones[0].Data,
			best_ones[1].Data,
		})
	}
}