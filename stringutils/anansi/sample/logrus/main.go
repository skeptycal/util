package main

import (
	"github.com/skeptycal/anansi"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{ForceColors: true})

	// logrus.SetOutput(colorable.NewColorableStdout())
	logrus.SetOutput(anansi.Output)

	logrus.Info("succeeded")
	logrus.Warn("not correct")
	logrus.Error("something error")
	logrus.Fatal("panic")
}
