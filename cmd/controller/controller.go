package main

import (
	"flag"
	"github.com/davejohnston/disco/internal/command"
	"github.com/davejohnston/disco/internal/service"
)

func main() {
	flag.Parse()

	controller := command.NewMockCommandSvc()
	service := &service.Service{
		Controller: controller,
	}

	service.Run()
}
