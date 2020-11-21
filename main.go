package main

import (
	"github.com/pikomonde/i-view-prospace/delivery"
	servPars "github.com/pikomonde/i-view-prospace/service/parser"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	// setup repository

	// setup service
	sPars := servPars.New()

	// setup delivery
	d := delivery.New(delivery.Opt{
		IsService:     true,
		ServiceParser: sPars,
	})
	d.Start()

}
