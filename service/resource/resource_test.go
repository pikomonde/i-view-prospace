package resource_test

import (
	"testing"

	"github.com/pikomonde/i-view-prospace/service/resource"
	"github.com/stretchr/testify/assert"
)

func getResourcePrices() map[string][2]int {
	return map[string][2]int{
		"Silver": {2, 34},
		"Gold":   {4, 57800},
		"Iron":   {20, 3910},
	}
}

func getErrorResourcePrices() map[string][2]int {
	return map[string][2]int{
		"Burger": {0, 34},
		"Oil":    {0, 0},
	}
}

func get() {

}

func TestAddResourcePrice(t *testing.T) {
	sRes := resource.New()
	sResError := resource.New()

	// Positive testing
	for k, v := range getResourcePrices() {
		err := sRes.AddResourcePrice(k, v[0], v[1])
		assert.Equal(t, nil, err)
	}
	assert.Equal(t, len(getResourcePrices()), len(sRes.Dict))

	// Negative testing
	for k, v := range getErrorResourcePrices() {
		err := sResError.AddResourcePrice(k, v[0], v[1])
		assert.Equal(t, resource.ErrInvalidZeroResourceUnit, err)
	}
	assert.Equal(t, 0, len(sResError.Dict))
}

func TestGetResourcePrice(t *testing.T) {
	sRes := resource.New()
	for k, v := range getResourcePrices() {
		sRes.AddResourcePrice(k, v[0], v[1])
	}

	// Positive testing
	for k, v := range getResourcePrices() {
		total, err := sRes.GetResourcePrice(10, k)
		assert.Equal(t, float64(v[1])/float64(v[0])*10, total)
		assert.Equal(t, nil, err)
	}

	// Negative testing
	for k := range getErrorResourcePrices() {
		total, err := sRes.GetResourcePrice(10, k)
		assert.Equal(t, float64(0), total)
		assert.Equal(t, resource.ErrNoResourceFound, err)
	}
}

func TestGetResourcePriceNotInit(t *testing.T) {
	sRes := resource.New()
	total, err := sRes.GetResourcePrice(10, "-")
	assert.Equal(t, float64(0), total)
	assert.Equal(t, resource.ErrNoResourceFound, err)
}
