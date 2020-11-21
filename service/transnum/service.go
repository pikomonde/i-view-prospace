package transnum

type serviceTransnum struct {
	Dict dictionary
}

// New returns serviceTransnum service
func New() *serviceTransnum {
	s := serviceTransnum{}
	s.Dict = make(dictionary)
	return &s
}
