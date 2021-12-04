package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "github.com/irojas14/Lab3INF343/Proto"
	"google.golang.org/grpc"
)

const (
	port    = ":50054"
	local   = "localhost" + port
	address = "dist150.inf.santiago.usm.cl" + port
)

const (
	nameNodeFile = "registro.txt"
)

const (
	Port    = ":50055"
	f1Addrs = "dist149.inf.santiago.usm.cl" + Port
	f2Addrs = "dist151.inf.santiago.usm.cl" + Port
	f3Addrs = "dist152.inf.santiago.usm.cl" + Port
)

const (
	port1 = "50050"
	port2 = "50051"
	port3 = "50052"
	Local = "localhost"
	local1  = Local + port1
	local2  = Local + port2
	local3  = Local + port3
)

var (
	RemoteAddrs [3]string = [3]string{f1Addrs, f2Addrs, f3Addrs}
	localAddrs  [3]string = [3]string{local1, local2, local3}
	curAddrs    [3]string
)

var FulcrumAddresses [3]string = [3]string{f1Addrs, f2Addrs, f3Addrs}

type server struct {
	pb.UnimplementedMosEisleyServer
}

func (s *server) Comando(ctx context.Context, in *pb.SolicitudComando) (*pb.RespuestaComandoMosEisley, error){
	// Switch para revisar el tipo del comando y actuar.
	return &pb.RespuestaComandoMosEisley{}, nil
}

func main() {
	srvAddr := address
	curAddrs = RemoteAddrs
	if len(os.Args) == 2 {
		srvAddr = local
		curAddrs = localAddrs
	}
	log.Printf("Direcciones Fulcrum: %v\n", curAddrs)

	lis, err := net.Listen("tcp", srvAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterMosEisleyServer(s, &server{})
	log.Printf("Juego Inicializado: escuchando en %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}