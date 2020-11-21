package log_test

import (
	"errors"
	"testing"

	"github.com/pikomonde/i-view-prospace/helper/log"
	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	fields01 := log.Fields{
		"testKey01": "testValue01",
	}
	error01 := errors.New("error01")

	assert.NotPanics(t, func() {
		err := log.Error(fields01, error01)
		assert.Equal(t, err, error01)
	})

	assert.NotPanics(t, func() {
		err := log.Error(fields01, nil)
		assert.Equal(t, err, nil)
	})

	assert.NotPanics(t, func() {
		err := log.Error(nil, error01)
		assert.Equal(t, err, error01)
	})

	assert.NotPanics(t, func() {
		err := log.Error(nil, nil)
		assert.Equal(t, err, nil)
	})
}

func TestPrint(t *testing.T) {
	assert.NotPanics(t, func() {
		log.Print()
	})

	assert.NotPanics(t, func() {
		log.Printf()
	})

	assert.NotPanics(t, func() {
		log.Println()
	})
}
