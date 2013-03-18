package main

import (
	"fmt"
	"net"
	"thrift"
	"bigselect/api"
)

type GoServer struct {
	handler   api.IBigSelect
	processor *api.BigSelectProcessor
}

func NewGoServer() *GoServer {
	handler := NewBigSelectHandler()
	processor := api.NewBigSelectProcessor(handler)
	return &GoServer{handler: handler, processor: processor}
}

func Simple(processor *api.BigSelectProcessor, transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory, ch chan int) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:9090")
	if err != nil {
		fmt.Print("Error resolving address: ", err.Error(), "\n")
		return
	}
	serverTransport, err := thrift.NewTServerSocketAddr(addr)
	if err != nil {
		fmt.Print("Error creating server socket: ", err.Error(), "\n")
		return
	}
	server := thrift.NewTSimpleServer4(processor, serverTransport, transportFactory, protocolFactory)
	// Use this for a multithreaded server
	// TServer server = new TThreadPoolServer(new TThreadPoolServer.Args(serverTransport).processor(processor));

	fmt.Print("Starting the simple server... on ", addr, "\n")
	for {
		err = server.Serve()
		if err != nil {
			fmt.Print("Error during simple server: ", err.Error(), "\n")
			return
		}
	}
	fmt.Print("Done with the simple server\n")
	ch <- 1
}

func Secure(processor *api.BigSelectProcessor) {
	addr, _ := net.ResolveTCPAddr("tcp", "localhost:9091")
	serverTransport, _ := thrift.NewTNonblockingServerSocketAddr(addr)
	server := thrift.NewTSimpleServer2(processor, serverTransport)
	fmt.Print("Starting the secure server... on ", addr, "\n")
	server.Serve()
	fmt.Print("Done with the secure server\n")
}

func RunServer(transportFactory thrift.TTransportFactory, protocolFactory thrift.TProtocolFactory) {
	server := NewGoServer()
	ch := make(chan int)
	go Simple(server.processor, transportFactory, protocolFactory, ch)
	//go Secure(server.processor)
	_ = <-ch
}
