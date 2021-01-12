package main

import (
	"fmt"
	"time"

	"github.com/kardianos/service"
)

const serviceName = "Medium service"
const serviceDescription = "Simple service, just for fun"

var logger service.Logger

type program struct{}

func (p program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	fmt.Println(s.String() + " started")
	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {
	for {
		// Do work here
		fmt.Println("Service is running")
		time.Sleep(1 * time.Second)
	}
}

func main() {

	logger.Info("logger started...")

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		logger.Errorf("Cannot create the service: %v\n", err.Error())
	}
	err = s.Run()
	if err != nil {
		logger.Errorf("Cannot start the service: " + err.Error())
	}
}
