package singleton

import "go-service/infrastructure/scheduler/interfaces"

func SchedulerSingleton() interfaces.SchedulerInterface {
	return scheduler
}
