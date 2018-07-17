package types

type ServiceInterface interface {
	AvgTime() float64
	Online24() float32
	ToService() *Service
	SmallText() string
	GraphData() string
	AvgUptime() string
	LimitedFailures() []*Failure
	TotalFailures() (uint64, error)
	TotalFailures24Hours() (uint64, error)
}

type FailureInterface interface {
	ToFailure() *Failure
	Ago() string
	ParseError() string
}
