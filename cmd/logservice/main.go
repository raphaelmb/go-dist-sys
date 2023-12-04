package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/raphaelmb/go-dist-sys/log"
	"github.com/raphaelmb/go-dist-sys/registry"
	"github.com/raphaelmb/go-dist-sys/service"
)

func main() {
	log.Run("./app.log")
	host, port := "localhost", "4000"
	serviceAddress := fmt.Sprintf("http://%v:%v", host, port)
	r := registry.Registration{
		ServiceName: registry.LogService,
		ServiceURL:  serviceAddress,
	}
	ctx, err := service.Start(context.Background(), r, host, port, log.RegisterHandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()
	fmt.Println("Shutting down Log Service")
}
