package transnum

// ServiceTransnum contains galactic, roman, and decimal numerals related services
type ServiceTransnum struct {
	Dict dictionary
}

// New returns ServiceTransnum service
func New() *ServiceTransnum {
	s := ServiceTransnum{}
	s.Dict = make(dictionary)
	return &s
}
