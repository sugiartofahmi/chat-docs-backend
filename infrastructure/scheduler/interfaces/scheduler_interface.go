package interfaces

type SchedulerInterface interface {
	Schedule(name string, spec string, callback func()) error
	Start()
}
