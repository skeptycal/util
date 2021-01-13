// Package stack generates a very basic web stack with
// Go and Javascript.
//
// An interface and configuration are provided and a service is
// returned
package stack

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/kardianos/service"
	log "github.com/sirupsen/logrus"
)

var (
	serviceIsRunning bool
	programIsRunning bool
	writingSync      sync.Mutex
)

func NewService(c *service.Config) (service.Service, error) {
	i := NewInterface()
	s, err := service.New(i, c)
	if err != nil {
		log.Errorf("Cannot create the service: %v\n", err.Error())
	}
	return s, err
}

func NewInterface() Program {
	return &program{}
}

type Program interface {
	Start(s service.Service) error
	Stop(s service.Service) error
	run()
}

type program struct{}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.

	writingSync.Lock()
	serviceIsRunning = true
	writingSync.Unlock()

	fmt.Println(s.String() + " started")

	go p.run()
	return nil
}

func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return within a few seconds.
	writingSync.Lock()
	serviceIsRunning = false
	writingSync.Unlock()

	// for programIsRunning {
	// wait for cleanup ...
	fmt.Println(s.String() + " stopping...")
	time.Sleep(500 * time.Millisecond)
	// }
	fmt.Println(s.String() + " stopped")
	return nil
}

func (p *program) run() {
	// for serviceIsRunning {
	// 	writingSync.Lock()
	// 	programIsRunning = true
	// 	writingSync.Unlock()

	router := httprouter.New()
	router.ServeFiles("/js/*filepath", http.Dir("js"))
	router.ServeFiles("/css/*filepath", http.Dir("css"))

	router.GET("/", serveHomepage)

	err := http.ListenAndServe(":81", router)
	if err != nil {
		log.Errorf("Problem starting web server: %v", err.Error())
	}
	// 	writingSync.Lock()
	// 	programIsRunning = false
	// 	writingSync.Unlock()
	// }
}

func serveHomepage(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

}
