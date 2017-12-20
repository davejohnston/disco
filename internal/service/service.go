// Package Service provides a runnable service that
// starts both a GRPC and RESTful Endpoint.
package service

import (
	"github.com/davejohnston/disco/internal/api"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
	"sync"
)

// Service is a simple service that exposes one public method
// Run().
// When called it will start a GRPC and Restful API listener
//
type Service struct {
	GrpcListenIf     string
	RestListenIf     string
	EdgeGrpcListenIf string
	Controller       api.CommandServer
}

// Run starts the GRPC and Restful listeners
// If the listen addresses have not been configured
// then it will assume defaults of port 8080 for GRPC
// and port 9999 for Rest
func (s Service) Run() {

	if s.GrpcListenIf == "" {
		s.GrpcListenIf = "localhost:8080"
	}
	if s.EdgeGrpcListenIf == "" {
		s.EdgeGrpcListenIf = "localhost:8081"
	}
	if s.RestListenIf == "" {
		s.RestListenIf = ":9999"
	}

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		err := s.startGrpcServer()
		if err != nil {
			glog.Fatal(err)
		}
	}()

	go func() {
		defer wg.Done()
		err := s.startRestAPI()
		if err != nil {
			glog.Fatal(err)
		}
	}()

	// Wait for both goroutines to exit
	wg.Wait()
}

// startGrpcServer starts the grpc server listening on the
// address and port defined by GrpcListenIf
func (s Service) startGrpcServer() error {

	lis, err := net.Listen("tcp", s.GrpcListenIf)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterCommandServer(grpcServer, s.Controller)
	if err != nil {
		return err
	}
	log.Printf("Serving GRPC on Port %s", s.GrpcListenIf)
	return grpcServer.Serve(lis)
}

// startRestAPI starts the rest api listening on the
// address and port defined by RestListenIf.  The RestAPI
// acts as a gateway to the grpc service.
func (s Service) startRestAPI() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := api.RegisterCommandHandlerFromEndpoint(ctx, mux, s.GrpcListenIf, opts)
	if err != nil {
		return err
	}
	log.Printf("Serving Rest on Port %s", s.RestListenIf)
	return http.ListenAndServe(s.RestListenIf, mux)
}
