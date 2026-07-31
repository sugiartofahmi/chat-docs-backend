package interfaces

type JobInterface interface {
	Name() string
	Handle(payload []byte) error
}
