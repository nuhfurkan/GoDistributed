package serve

import representation "go-distributed/representations"

type Experiment struct {
	Representation	string							`json:"representation"`
	Mutation		string							`json:"mutation"`
	Generation		string							`json:"generation"`
	Payload			representation.GeneratePayload	`json:"payload"`
}

type SetupConfig struct {
	Filename		string	`json:"filename"`
}