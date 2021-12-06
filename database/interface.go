package database

type DbObject interface {
	Create() error
	Update() error
	Delete() error
}

type Sampler interface {
	Sample() DbObject
}
