package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/raphaelmb/go-dist-sys/grades"
	"github.com/raphaelmb/go-dist-sys/log"
	"github.com/raphaelmb/go-dist-sys/registry"
	"github.com/raphaelmb/go-dist-sys/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
	}

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	logProvider, err := registry.GetProvider(registry.LogService)
	if err == nil {
		fmt.Printf("Logging service found at: %v\n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()
	fmt.Println("Shutting down Grading Service")
}
