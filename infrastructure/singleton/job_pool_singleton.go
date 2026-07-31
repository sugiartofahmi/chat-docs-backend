package singleton

import "go-service/infrastructure/job/interfaces"

func JobPoolSingleton() interfaces.JobPoolInterface {
	return jobPool
}
