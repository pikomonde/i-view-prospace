package resource

// AddResourcePrice is used to add resource's price to database
func (s *ServiceResource) AddResourcePrice(resourceName string, resourceNum, credits int) error {
	if resourceNum == 0 {
		return ErrInvalidZeroResourceUnit
	}
	s.Dict[resourceName] = float64(credits) / float64(resourceNum)
	return nil
}

// GetResourcePrice is used to get resource's price to database
func (s *ServiceResource) GetResourcePrice(resourceNum int, resourceName string) (float64, error) {
	resourcePrice, ok := s.Dict[resourceName]
	if !ok {
		return 0, ErrNoResourceFound
	}
	return resourcePrice * float64(resourceNum), nil
}
