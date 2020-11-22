package delivery_test

import (
	"syscall"
	"testing"
	"time"

	"github.com/pikomonde/i-view-prospace/delivery"
)

func TestDelivery(t *testing.T) {
	d := delivery.New(delivery.Opt{})
	go d.Start()

	// Give a time to finish the d.Start() to stuck on select case.
	time.Sleep(500 * time.Millisecond)

	// Simulate syscall.SIGINT
	syscall.Kill(syscall.Getpid(), syscall.SIGINT)

	// Give another time d.Start() receive syscall.SIGINT signal.
	time.Sleep(500 * time.Millisecond)
}
