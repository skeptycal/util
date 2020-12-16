package main

import (
	"fmt"
	"io"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/anansi"
)

func main() {
	Formatter := new(log.TextFormatter)
	Formatter.TimestampFormat = "02-01-2006 15:04:05"
	Formatter.FullTimestamp = true
	log.SetFormatter(Formatter)

	log.Info("Some info...")
	log.Warning("This is a warning")
	log.Error("Not fatal. An error. Won't stop execution")
	// log.Fatal("MAYDAY MAYDAY MAYDAY. Execution will be stopped here")
	// log.Panic("Do not panic")

	b, err := io.Copy(anansi.Output, os.Stdin)
	if err != nil {
		log.Error(fmt.Errorf("Failed to copy. %v bytes written.", b))
	}
}
