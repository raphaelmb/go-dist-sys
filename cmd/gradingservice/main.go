package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/raphaelmb/go-dist-sys/grades"
	"github.com/raphaelmb/go-dist-sys/registry"
	"github.com/raphaelmb/go-dist-sys/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(), r, host, port, grades.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down Grading Service")
}
