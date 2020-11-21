package resource

type serviceResource struct {
	Dict dictionary
}

// New returns ServiceResource service
func New() *serviceResource {
	s := serviceResource{}
	s.Dict = make(dictionary)
	return &s
}
