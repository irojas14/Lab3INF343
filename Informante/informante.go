package main

import (
	"log"
	"os"

	pb "github.com/irojas14/Lab3INF343/Proto"
	"google.golang.org/grpc"
)

const (
	port   = ":50052"
	local  = "localhost" + port
	maddrs = "dist149.inf.santiago.usm.cl" + port
)

func main() {
	
	srvAddr := maddrs
	if len(os.Args) == 2 {
		srvAddr = local
	}

	conn, err := grpc.Dial(srvAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewMosEisleyClient(conn)

	log.Printf("%v\n", c)
}
