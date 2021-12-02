package main

import (
	"log"
	"net"
	"os"

	pb "github.com/irojas14/Lab3INF343/Proto"
	"google.golang.org/grpc"
)

const (
	port = ":50500"
	port2 = ":50501"
	port3 = ":50502"
	local = "localhost"
	local1   = local + port
	local2 = local + port2
	local3 = local + port3
	f1Addrs = "dist149.inf.santiago.usm.cl" + port
	f2Addrs = "dist151.inf.santiago.usm.cl" + port
	f3Addrs = "dist152.inf.santiago.usm.cl" + port
)

var (
	filesF1  = "1"
	filesF2  = "2"
	filesF3  = "3"
	curFiles = "Fulcrum/"
)

type server struct {
	pb.UnimplementedFulcrumServer
}

func main() {
	var srvAddr string
	if len(os.Args) == 3 {
		if (os.Args[1] == "1") {
			srvAddr = local + port
		} else if (os.Args[1] == "2") {
			srvAddr = local + port2
		} else {
			srvAddr = local + port3
		}
	} else {
		if (os.Args[1] == "1") {
			srvAddr = f1Addrs
		} else if (os.Args[1] == "2") {
			srvAddr = f2Addrs
		} else {
			srvAddr = f3Addrs
		}
	}

	lis, err := net.Listen("tcp", srvAddr)

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterFulcrumServer(s, server{})
	log.Printf("Juego Inicializado: escuchando en %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}