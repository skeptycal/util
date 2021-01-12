package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

const serviceName = "Medium service"
const serviceDescription = "Simple service, just for fun"

var (
	serviceIsRunning bool
	programIsRunning bool
	writingSync      sync.Mutex
)

// var logger service.Logger

type program struct{}

func (p program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	fmt.Println(s.String() + " started")

	writingSync.Lock()
	serviceIsRunning = true
	writingSync.Unlock()

	go p.run()
	return nil
}

func (p program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	writingSync.Lock()
	serviceIsRunning = false
	writingSync.Unlock()

	for programIsRunning {
		fmt.Println(s.String() + " stopping...")
		time.Sleep(1 * time.Second)
	}

	fmt.Println(s.String() + " stopped")
	return nil
}

func (p program) run() {

	for serviceIsRunning {
		// Do work here

		writingSync.Lock()
		programIsRunning = true
		writingSync.Unlock()
		fmt.Println("Service is running")
		time.Sleep(1 * time.Second)
		writingSync.Lock()
		programIsRunning = false
		writingSync.Unlock()
	}
}

func main() {

	log.Info("logger started...")

	serviceConfig := &service.Config{
		Name:        serviceName,
		DisplayName: serviceName,
		Description: serviceDescription,
	}
	prg := &program{}
	s, err := service.New(prg, serviceConfig)
	if err != nil {
		log.Errorf("Cannot create the service: %v\n", err.Error())
	}
	err = s.Run()
	if err != nil {
		log.Errorf("Cannot start the service: " + err.Error())
	}
}
