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
	leia_port = ":50054"
	local     = "localhost" + leia_port
	leia_addr = "dist151.inf.santiago.usm.cl" + leia_port
)

type server struct {
	pb.UnimplementedFulcrumServer
}

type Consulta struct {
	archivo_name string
	coord        *pb.Ubicacion
	rebel_num    int64
	reloj_vec    *pb.RelojVector
	fulcrum_dir  string
}

var (
	Tipo      pb.ActorType
	consultas []Consulta
)

func main() {
	var srvAddr string
	if len(os.Args) == 3 {
		srvAddr = local
	} else {
		srvAddr = leia_addr
	}

	lis, err := net.Listen("tcp", srvAddr)

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterFulcrumServer(s, server{})
	log.Printf("Bienvenida Leia Organa: escuchando en %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// FUNCION LEIA

func GetNumberRebelds(c pb.MosEisleyClient, nombre_planeta string, nombre_ciudad string) error {
	if Tipo == pb.ActorType_Leia {
		// Conectarse y pedir la informacion a Mos Eisley

		coord := pb.Ubicacion{
			NombrePlaneta: nombre_planeta,
			NombreCiudad:  nombre_ciudad,
		}

		r, err := c.GetNumberRebelds(context.Background(), &pb.SolicitudGetNumberRebelds{
			Coord: &coord,
		})

		if err != nil {
			log.Fatalf("Error al preguntar por rebeldes en %v - %v. Error: %v\n", coord.NombrePlaneta, coord.NombreCiudad, err)
			return err
		}

		consultas = append(consultas, Consulta{
			archivo_name: r.ArchivoName,
			coord:        &coord,
			rebel_num:    r.NumRebels,
			reloj_vec:    r.RelojVec,
			fulcrum_dir:  r.FulcrumDir,
		})
	}
	return nil
}
