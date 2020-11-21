package resource

// ServiceResource contains galactic, roman, and decimal numerals related services
type ServiceResource struct {
	Dict dictionary
}

// New returns ServiceResource service
func New() *ServiceResource {
	s := ServiceResource{}
	s.Dict = make(dictionary)
	return &s
}
