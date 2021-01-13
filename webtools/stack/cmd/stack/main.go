package main

import (
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
	"github.com/skeptycal/util/webtools/stack"
)

// var logger service.Logger

const serviceName = "Medium service"
const serviceDescription = "Simple service, just for fun"

func main() {

	log.Info("logger started...")

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}

	s, err := stack.NewService(serviceConfig)

	if err != nil {
		log.Errorf("Cannot create the service: %v\n", err.Error())
	}
	err = s.Run()
	if err != nil {
		log.Errorf("Cannot start the service: " + err.Error())
	}
}
