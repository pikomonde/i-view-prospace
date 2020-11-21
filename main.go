package main

import (
	"fmt"

	servPars "github.com/pikomonde/i-view-prospace/service/parser"
	servRsrc "github.com/pikomonde/i-view-prospace/service/resource"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.InfoLevel)

	// setup repository

	// setup service
	sPars := servPars.New()
	sPars.Parse("glob is I")
	sPars.Parse("prok is V")
	sPars.Parse("pish is X")
	sPars.Parse("tegj is L")

	sRsrc := servRsrc.New()
	sRsrc.AddResourcePrice("Silver", sPars.ServiceTransnum.MustGalaticToInt([]string{"glob", "glob"}), 34)
	sRsrc.AddResourcePrice("Gold", sPars.ServiceTransnum.MustGalaticToInt([]string{"glob", "prok"}), 57800)
	sRsrc.AddResourcePrice("Iron", sPars.ServiceTransnum.MustGalaticToInt([]string{"pish", "pish"}), 3910)
	sRsrc.AddResourcePrice("Porridge", sPars.ServiceTransnum.MustGalaticToInt([]string{"lolo", "kkkk"}), 123)

	fmt.Println(sRsrc.GetResourcePrice(sPars.ServiceTransnum.MustGalaticToInt([]string{"glob", "prok"}), "Silver"))
	fmt.Println(sRsrc.GetResourcePrice(sPars.ServiceTransnum.MustGalaticToInt([]string{"glob", "prok"}), "Gold"))
	fmt.Println(sRsrc.GetResourcePrice(sPars.ServiceTransnum.MustGalaticToInt([]string{"glob", "prok"}), "Iron"))

	// setup delivery

}
