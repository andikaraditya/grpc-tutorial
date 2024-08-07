package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/andikaraditya/grpc-tutorial/invoicer"
	"google.golang.org/grpc"
)

type myInvoicerServer struct {
	invoicer.UnimplementedInvoicerServer
}

func (s myInvoicerServer) Create(ctx context.Context, req *invoicer.CreateRequest) (*invoicer.CreateResponse, error) {

	return &invoicer.CreateResponse{
		Pdf:  []byte(req.From),
		Docx: []byte(req.To),
	}, nil
}

func main() {
	port := 3000
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal(fmt.Sprintf("Cannot create listener: %s", err))
	}

	serverRegistrar := grpc.NewServer()
	service := &myInvoicerServer{}

	invoicer.RegisterInvoicerServer(serverRegistrar, service)
	fmt.Printf("Listening to http://localhost:%d", port)
	err = serverRegistrar.Serve(lis)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error serve: %s", err))
	}
}
