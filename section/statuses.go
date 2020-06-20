package section

// Status ... A status for a process
type Status string

const (
	// NotStarted ... A process has not been executed yet
	NotStarted Status = "Not Started"
	// Started ... A process currently executing
	Started Status = "Running"
	// Failed ... A process that finished with a exit code other than 0
	Failed Status = "Failed :("
	// Success ... A process that finished with a exit code of 0
	Success Status = "Finished"
)
