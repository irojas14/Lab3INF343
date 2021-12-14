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
	Port    = ":50500"
	f1Addrs = "dist149.inf.santiago.usm.cl" + Port
	f2Addrs = "dist151.inf.santiago.usm.cl" + Port
	f3Addrs = "dist152.inf.santiago.usm.cl" + Port
)

const (
	port1 = ":50500"
	port2 = ":50501"
	port3 = ":50502"
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

var (
	curElection int = 0;
)

type server struct {
	pb.UnimplementedMosEisleyServer
}

func (s *server) Comando(ctx context.Context, in *pb.SolicitudComando) (*pb.RespuestaComandoMosEisley, error){
	// Switch para revisar el tipo del comando y actuar.
	log.Printf("En Mos Eisley Comando, a elegir Fulcrum")
	faddr := EleccionFulcrum();
	log.Printf("Dirección Elegida %v\n", faddr)
	return &pb.RespuestaComandoMosEisley{DirFulcrum: faddr,}, nil
}




func (s *server) GetNumberRebelds(ctx context.Context, in *pb.SolicitudGetNumberRebelds) (*pb.RespuestaGetNumberRebelds, error) {
	log.Printf("En GetNumberRebelds: Consulta coord: %v\n", in.Consulta.Coord)	

	// Preguntamos y Definimos la direccion Fulcrum
	var addr string
	if (in.Consulta.RelojVec != nil || in.Consulta.FulcrumDir != "") {
		addr = in.Consulta.FulcrumDir
	} else {
		addr = EleccionFulcrum()
	}

	// Realizamos la Consulta
	r, err := ConsultaDelVivo(addr, in.Consulta)

	
	// Retornamos los resutlados
	if (err != nil) {
		log.Fatalf("Error al Consultar Vivamente los Rebels en %v desde el Fulcrum %v\n", in.Consulta.Coord, addr)
	}

	return r, err
}

func (s *server) GetNumberRebeldsInformante(ctx context.Context, in *pb.SolicitudGetNumRebelsInformante) (*pb.RespuestaGetNumRebelsInformante, error) {

	log.Printf("En Get Number Rebelds Informante: In: %v\n", in)
	
	var addr string
	if in.Cambio.FulcrumDir != "" {
		addr = in.Cambio.FulcrumDir
	} else {
		addr = EleccionFulcrum()
	}

	log.Printf("Addrs final: %v\n", addr)

	consulta := &pb.Consulta{
		ArchivoName: in.Cambio.ArchivoName,
		Coord: in.Cambio.Cmd.Coord,
		RelojVec: in.Cambio.RelojVec,
		FulcrumDir: addr,
	}

	// Realizamos la Consulta
	var r *pb.RespuestaGetNumRebelsInformante = nil
	rConsulta, errConsulta := ConsultaDelVivo(addr, consulta)

	log.Printf("rConsulta Recibida: %v\n", rConsulta)

	if errConsulta != nil {
		log.Fatalf("Ocurrió un error al querer confirmar el read de %v en el fulcrum %v\n", in.Cambio.Cmd.Coord, in.Cambio.FulcrumDir)
		return nil, errConsulta
	}

	if rConsulta != nil {
		log.Printf("Traduciendo Consulta a versión de Informante")
		r = &pb.RespuestaGetNumRebelsInformante{
			Cambio: &pb.Cambio{
				ArchivoName: rConsulta.Consulta.ArchivoName,
				FulcrumDir: rConsulta.Consulta.FulcrumDir,
				RelojVec: rConsulta.Consulta.RelojVec,
				Cmd: in.Cambio.Cmd,
			},
			NumRebels: rConsulta.Consulta.NumRebels,
		}
	}
	return r, nil
}

func ConsultaDelVivo(addr string, consulta *pb.Consulta) (*pb.RespuestaGetNumberRebelds, error) {
	connFulcrum, errFulcrum := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

	// CONEXION
	// Nos conectamos al Fulcrum con direccion "addr"
	if errFulcrum != nil {
		log.Fatalf("Error al conectarse al fulcrum %v\n", addr)
		return nil, errFulcrum
	}
	defer connFulcrum.Close()

	// Creamos el Cliente Fulcrum

	c := pb.NewFulcrumClient(connFulcrum)

	rRebelds, errRebelds := c.GetNumberRebelds(context.Background(), &pb.SolicitudGetNumberRebelds{
		Consulta: consulta,
	})

	if errRebelds != nil {
		log.Fatalf("Error Solicitar Get Number Rebelds en coord: %v, en el Fulcrum %v\n", consulta.Coord, addr)
		return nil, errRebelds
	}

	if (rRebelds == nil) {
		log.Fatalf("No Existe el Planeta %v en el Fulcrum %v\n", consulta.Coord.NombrePlaneta, addr)
	}

	return rRebelds, nil
}

func ConsultaDelNoVivo(in *pb.SolicitudGetNumberRebelds) (*pb.RespuestaGetNumberRebelds, error) {
	valores := make([]*pb.RespuestaGetNumberRebelds, 0)
	
	for _, addr := range(curAddrs) {
		
		// Creamos la conexion

		connFulcrum, errFulcrum := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

		// Nos conectamos al Fulcrum con direccion "addr"

		if errFulcrum != nil {
			log.Fatalf("Error al conectarse al fulcrum %v\n", addr)
			connFulcrum.Close()
			continue;
		}
		defer connFulcrum.Close()


		// Creamos el Cliente Fulcrum

		c := pb.NewFulcrumClient(connFulcrum)

		// Realizamos el llamdo remoto a Fulcrum

		rRebelds, errRebelds := c.GetNumberRebelds(context.Background(), &pb.SolicitudGetNumberRebelds{
			Consulta : in.Consulta,
		})

		if errRebelds != nil {
			log.Printf("Error Solicitar Get Number Rebelds en coord: %v, en el Fulcrum %v\n", in.Consulta.Coord, addr)
			continue
		}

		if rRebelds.Consulta == nil {
			log.Printf("Planeta %v no presente en el Fulcrum %v\n", in.Consulta.Coord.NombrePlaneta, addr)
			continue
		}

		if (rRebelds != nil && rRebelds.Consulta.NumRebels != -3) {
			valores = append(valores, rRebelds)
		}
	}

	val_len := len(valores)
	if val_len > 0 {

		max := &pb.RespuestaGetNumberRebelds{
			Consulta: &pb.Consulta{
				RelojVec: &pb.RelojVector{X:-1, Y:-1, Z:-1,},
				},
			}

		for i := 0; i < val_len; i++ {
			if RelojVecComparison(valores[i], max) == 0 {
				max = valores[i]
			}
		}

		log.Printf("Respuesta Elegida: %v con vector %v\n", max, max.Consulta.RelojVec)
		
		return max, nil

	} else {

		return &pb.RespuestaGetNumberRebelds{
			Consulta: &pb.Consulta{
				NumRebels: -3,
			},
		}, nil
	}
}

// FUNCIONES AUXILIARES

func EleccionFulcrum() string {

	faddr := curAddrs[curElection];
	curElection++;
	if (curElection >= 3) {
		curElection = 0;
	}
	return faddr
}

func RelojVecComparison(left *pb.RespuestaGetNumberRebelds, right *pb.RespuestaGetNumberRebelds) int {

	lCon := left.Consulta
	rCon := right.Consulta

	if lCon.RelojVec.X == rCon.RelojVec.X && lCon.RelojVec.Y == rCon.RelojVec.Y && lCon.RelojVec.Z == rCon.RelojVec.Z {
		return 1
	}

	if lCon.RelojVec.X <= rCon.RelojVec.X && lCon.RelojVec.Y <= rCon.RelojVec.Y && lCon.RelojVec.Z <= rCon.RelojVec.Z {
		return 1
	}

	if lCon.RelojVec.X >= rCon.RelojVec.X && lCon.RelojVec.Y >= rCon.RelojVec.Y && lCon.RelojVec.Z >= rCon.RelojVec.Z {
		return 0
	}
	return 1
}

func main() {
	srvAddr := address
	curAddrs = RemoteAddrs
	if len(os.Args) == 2 {
		srvAddr = local
		curAddrs = localAddrs
	}
	log.Printf("Dirección MosEisley: %v\n", srvAddr)
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