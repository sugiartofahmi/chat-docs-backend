package job

import "go-service/infrastructure/job/interfaces"

type JobRegistry struct {
	jobs map[string]interfaces.JobInterface
}

func NewJobRegistry() interfaces.JobRegistryInterface {
	return &JobRegistry{jobs: make(map[string]interfaces.JobInterface)}
}

func (r *JobRegistry) Register(job interfaces.JobInterface) {
	r.jobs[job.Name()] = job
}

func (r *JobRegistry) Get(name string) (interfaces.JobInterface, bool) {
	job, ok := r.jobs[name]
	return job, ok
}
