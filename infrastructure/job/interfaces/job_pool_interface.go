package interfaces

type JobPoolInterface interface {
	Start(numWorkers int)
	Delegate(jobName string, payload any)
}
