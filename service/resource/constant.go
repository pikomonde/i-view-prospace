package resource

import "errors"

type dictionary map[string]float64

// ErrInvalidZeroResourceUnit is returned when resouce unit is zero.
var ErrInvalidZeroResourceUnit = errors.New("Resource unit should not be zero")

// ErrNoResourceFound is returned when resource is not found in dictionary.
var ErrNoResourceFound = errors.New("Invalid resource")
