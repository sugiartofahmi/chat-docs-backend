package interfaces

type JobRegistryInterface interface {
	Register(job JobInterface)
	Get(name string) (JobInterface, bool)
}
