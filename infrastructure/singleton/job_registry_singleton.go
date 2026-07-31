package singleton

import "go-service/infrastructure/job/interfaces"

func JobRegistrySingleton() interfaces.JobRegistryInterface {
	return jobRegistry
}
