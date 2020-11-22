package cli

import (
	"bufio"
	"fmt"
	"io"

	"github.com/pikomonde/i-view-prospace/helper/log"
	servPars "github.com/pikomonde/i-view-prospace/service/parser"
)

// Cli is used to contains cli delivery
type Cli struct {
	ServiceParser *servPars.ServiceParser
	IOReader      io.Reader
}

// Start starts Cli delivery
func (c *Cli) Start() {
	sPars := c.ServiceParser

	scanner := bufio.NewScanner(c.IOReader)
	for scanner.Scan() {
		out := sPars.Parse(scanner.Text())
		if out != "" {
			fmt.Println(out)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Error(nil, err)
	}
}
