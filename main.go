package main

import (
	"fmt"

	tnService "github.com/pikomonde/i-view-prospace/service/transnum"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	// setup repository

	// setup service
	sr := tnService.ServiceTransnum{}
	sr.AddGalacticUnit("glob", 'I')
	sr.AddGalacticUnit("prok", 'V')
	sr.AddGalacticUnit("pish", 'X')
	sr.AddGalacticUnit("tegj", 'L')

	fmt.Println(sr.GalaticToInt([]string{"glob", "glob"}))
	fmt.Println(sr.GalaticToInt([]string{"glob", "prok"}))
	fmt.Println(sr.GalaticToInt([]string{"pish", "pish"}))

	// setup delivery

}
