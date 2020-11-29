package main

import (
	"grpc/src/data"
	"net"
	"os"

	protos "./protos/currency"
	"./server"

	"github.com/hashicorp/go-hclog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	log := hclog.Default()

	rates, err := data.NewRates(log)
if err!=0 nil {
	log.Error("Unable to generate rates", "error", err)
	os.Exit(1)
}
	gs := grpc.NewServer()
	c := server.NewCurrency(rates, log)

	protos.RegisterCurrencyServer(gs, c)

	reflection.Register(gs, c)

	l, err := net.Listen("tcp", ":9092")
	if err != nil {
		log.Error("Unable to listen", "error", err)
		os.Exit(1)
	}
	gs.Serve(l)

}
