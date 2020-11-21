package main

import (
	"fmt"

	servRsrc "github.com/pikomonde/i-view-prospace/service/resource"
	servTnum "github.com/pikomonde/i-view-prospace/service/transnum"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	// setup repository

	// setup service
	sTnum := servTnum.ServiceTransnum{}
	sTnum.AddGalacticUnit("glob", 'I')
	sTnum.AddGalacticUnit("prok", 'V')
	sTnum.AddGalacticUnit("pish", 'X')
	sTnum.AddGalacticUnit("tegj", 'L')

	sRsrc := servRsrc.ServiceResource{}
	sRsrc.AddResourcePrice("Silver", sTnum.MustGalaticToInt([]string{"glob", "glob"}), 34)
	sRsrc.AddResourcePrice("Gold", sTnum.MustGalaticToInt([]string{"glob", "prok"}), 57800)
	sRsrc.AddResourcePrice("Iron", sTnum.MustGalaticToInt([]string{"pish", "pish"}), 3910)
	sRsrc.AddResourcePrice("Pourage", sTnum.MustGalaticToInt([]string{"lolo", "kkkk"}), 123)

	fmt.Println(sRsrc.GetResourcePrice(sTnum.MustGalaticToInt([]string{"glob", "prok"}), "Silver"))
	fmt.Println(sRsrc.GetResourcePrice(sTnum.MustGalaticToInt([]string{"glob", "prok"}), "Gold"))
	fmt.Println(sRsrc.GetResourcePrice(sTnum.MustGalaticToInt([]string{"glob", "prok"}), "Iron"))

	// setup delivery

}
