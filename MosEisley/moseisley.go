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
	Port    = ":50055"
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
	addr := EleccionFulcrum()
	if (in.Consulta.RelojVec != nil || in.Consulta.FulcrumDir != "") {
		addr = in.Consulta.FulcrumDir
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
	
	addr := EleccionFulcrum()
	if in.Cambio.FulcrumDir != "" {
		addr = in.Cambio.FulcrumDir
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
			ArchivoName: rConsulta.ArchivoName,
			FulcrumDir: rConsulta.FulcrumDir,
			NumRebels: rConsulta.NumRebels,
			RelojVec: rConsulta.RelojVec,
			Cmd: in.Cambio.Cmd,
		}
	}
	return r, nil
}

func Merge() {

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
	for _, addr := range(curAddrs){
		connFulcrum, errFulcrum := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

		// Nos conectamos al Fulcrum con direccion "addr"
		if errFulcrum != nil {
			log.Fatalf("Error al conectarse al fulcrum %v\n", addr)
			continue;
		}
		defer connFulcrum.Close()

		// Creamos el Cliente Fulcrum

		c := pb.NewFulcrumClient(connFulcrum)

		rRebelds, errRebelds := c.GetNumberRebelds(context.Background(), &pb.SolicitudGetNumberRebelds{
			Consulta : in.Consulta,
		})

		if errRebelds != nil {
			log.Fatalf("Error Solicitar Get Number Rebelds en coord: %v, en el Fulcrum %v\n", in.Consulta.Coord, addr)
			continue
		}

		if (rRebelds != nil && rRebelds.NumRebels != -3) {
			valores = append(valores, rRebelds)
		}
	}

	diff := false
	if len(valores) > 0 {
		last := valores[0].NumRebels;
		
		for i := 1; i < len(valores); i++ {
			if valores[i].NumRebels != last && valores[i].NumRebels != -3 && last != -3 {
				diff = true
				break
			}
			last = valores[i].NumRebels
		}

	} else {
		return &pb.RespuestaGetNumberRebelds{
			NumRebels: -3,
		}, nil
	}

	var respuestaGetRebelds *pb.RespuestaGetNumberRebelds = nil
	if (diff) {
		Merge()
		respuestaGetRebelds = valores[0]
	} else {
		respuestaGetRebelds = valores[0]
	}
	return respuestaGetRebelds, nil
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