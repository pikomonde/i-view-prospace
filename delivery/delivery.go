package delivery

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/pikomonde/i-view-prospace/delivery/cli"
	"github.com/pikomonde/i-view-prospace/helper/log"
	servPars "github.com/pikomonde/i-view-prospace/service/parser"
)

// Delivery contains on the app's delivery
type Delivery struct {
	Cli *cli.Cli
}

// Opt is used as an option for the app
type Opt struct {
	ServiceParser *servPars.ServiceParser
}

// New returns the delivery
func New(opt Opt) *Delivery {
	cli := cli.Cli{
		ServiceParser: opt.ServiceParser,
	}
	if opt.ServiceParser == nil {
		cli.ServiceParser = servPars.New()
	}

	return &Delivery{
		Cli: &cli,
	}
}

// Start starts the app
func (d *Delivery) Start() {
	go d.Cli.Start()

	term := make(chan os.Signal)
	signal.Notify(term, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-term:
		log.Println("Exiting gracefully...")
	}
}
