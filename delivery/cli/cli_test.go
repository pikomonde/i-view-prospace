package cli_test

import (
	"strings"
	"testing"
	"time"

	"github.com/pikomonde/i-view-prospace/delivery"
)

func TestDelivery(t *testing.T) {
	testCase := "blob is I\n"
	testCase += "how much is blob blob ?\n"

	d := delivery.New(delivery.Opt{
		IOReader: strings.NewReader(testCase),
	})
	go d.Start()

	// TODO: test input.Scanner() error

	// Give a time to finish the d.Start() to stuck on select case.
	time.Sleep(500 * time.Millisecond)
}
