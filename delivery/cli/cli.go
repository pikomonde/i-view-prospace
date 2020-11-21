package cli

import (
	"fmt"

	servPars "github.com/pikomonde/i-view-prospace/service/parser"
)

// Cli is used to contains cli delivery
type Cli struct {
	ServiceParser *servPars.ServiceParser
}

// Start starts Cli delivery
func (c *Cli) Start() {
	sPars := c.ServiceParser

	sPars.Parse("glob is I")
	sPars.Parse("prok is V")
	sPars.Parse("pish is X")
	sPars.Parse("tegj is L")

	sPars.Parse("glob glob Silver is 34 Credits")
	sPars.Parse("glob prok Gold is 57800 Credits")
	sPars.Parse("pish pish Iron is 3910 Credits")
	sPars.Parse("lolo kkkk Porridge is 123 Credits")

	fmt.Println(sPars.Parse("how much is pish tegj glob glob ?"))

	fmt.Println(sPars.Parse("how many Credits is glob prok Silver ?"))
	fmt.Println(sPars.Parse("how many Credits is glob prok Gold ?"))
	fmt.Println(sPars.Parse("how many Credits is glob prok Iron ?"))

	fmt.Println(sPars.Parse("how much wood could a woodchuck chuck if a woodchuck could chuck wood ?"))
}
