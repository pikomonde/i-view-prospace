package main

import (
	"syscall/js"

	servPars "github.com/pikomonde/i-view-prospace/service/parser"
	"github.com/sirupsen/logrus"
)

var sPars *servPars.ServiceParser

func main() {
	c := make(chan bool)
	logrus.SetLevel(logrus.InfoLevel)
	sPars = servPars.New()

	js.Global().Set("parse", js.FuncOf(parse))

	<-c
}

func parse(this js.Value, inputs []js.Value) interface{} {
	return sPars.Parse(inputs[0].String())
}
