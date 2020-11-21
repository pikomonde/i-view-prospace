package cli

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pikomonde/i-view-prospace/helper/log"
	servPars "github.com/pikomonde/i-view-prospace/service/parser"
)

// Cli is used to contains cli delivery
type Cli struct {
	ServiceParser *servPars.ServiceParser
}

// Start starts Cli delivery
func (c *Cli) Start() {
	sPars := c.ServiceParser

	scanner := bufio.NewScanner(os.Stdin)
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
