package parser

import (
	servRsrc "github.com/pikomonde/i-view-prospace/service/resource"
	servTnum "github.com/pikomonde/i-view-prospace/service/transnum"
)

// ServiceParser contains input parser related services
type ServiceParser struct {
	ServiceTransnum *servTnum.ServiceTransnum
	ServiceResource *servRsrc.ServiceResource
}

// New returns ServiceParser service
func New() *ServiceParser {
	s := ServiceParser{
		ServiceTransnum: servTnum.New(),
		ServiceResource: servRsrc.New(),
	}
	return &s
}
