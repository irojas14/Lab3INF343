package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	funcs "github.com/irojas14/Lab3INF343/Funciones"
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

var (
	FulcrumId int
	curFilesPath string
)

var (
	RelojesVectoresDict map[string]*pb.RelojVector = make(map[string]*pb.RelojVector)
)

type server struct {
	pb.UnimplementedFulcrumServer
}

func (s *server) Comando(ctx context.Context, in *pb.SolicitudComando) (*pb.RespuestaComandoFulcrum, error) {

	log.Printf("En Comando Fulcrum id: %v\n", FulcrumId)

	var RelojVectorRes *pb.RelojVector = nil

	switch in.Cmd.Tipo {
	case pb.TipoComando_AddCity:
		RelojVectorRes = AddCity(in.Cmd)
	case pb.TipoComando_DeleteCity:
		DeleteCity(in.Cmd)
	case pb.TipoComando_UpdateName:
		UpdateName(in.Cmd)
	case pb.TipoComando_UpdateNumber:
		UpdateNumber(in.Cmd)
	}
	return &pb.RespuestaComandoFulcrum{RelojVec: RelojVectorRes,}, nil
}

func AddCity(cmd *pb.Comando) *pb.RelojVector {
	nombreArchivo := cmd.Coord.NombrePlaneta + "_" + "Info"
	nombreArchivoLog := cmd.Coord.NombrePlaneta + "_" + "Log"
	
	existeBool := funcs.IsInServer(nombreArchivo, curFilesPath)

	if existeBool {
		// Si existe
		// Insertamos el comando en el Registro

		// Si no existe por alguna razon, creamos su Reloj de Vectores
		if RelojesVectoresDict[nombreArchivo] == nil {
			RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
				Nombre: nombreArchivo,
				X: 0, Y: 0, Z: 0,
		}
	}

		err := funcs.InsertarComandoEnRegistro(curFilesPath + "/" + nombreArchivo, cmd) // WIP
		if (err != nil) {
			return nil
		}

		// Registramos el cambio en el Log asociado
		err = funcs.InsertarComandoEnLog(curFilesPath + "/" + nombreArchivoLog, cmd)
		if (err != nil) {
			return nil
		}
	} else {
		// Si no existe

		// Creamos su Reloj de Vectores
		RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
			Nombre: nombreArchivo, 
			X: 0, Y: 0, Z: 0,
		}

		// Creamos el Archivo de Registro Planetario
		funcs.CrearTxt(curFilesPath + "/" + nombreArchivo) // creamos el archivo de Info

		// Insertamos el comando en el Registro Planetario
		err := funcs.InsertarComandoEnRegistro(curFilesPath + "/" + nombreArchivo, cmd)
		if (err != nil) {
			return nil
		}

		// Creamos el archivo de Log asociado
		funcs.CrearTxt(curFilesPath + "/" + nombreArchivoLog) // creamos el archivo Log

		// Registramos el cambio en el Log
		err = funcs.InsertarComandoEnLog(curFilesPath + "/" + nombreArchivoLog, cmd) //
		if (err != nil) {
			return nil
		}
	}
	// Modificamos el reloj de vectores correspondiente para reflejar el cambio
	ModificarRelojVector(nombreArchivo)

	return RelojesVectoresDict[nombreArchivo]
}

func DeleteCity(cmd *pb.Comando) {
	
}

func UpdateName(cmd *pb.Comando) {

}

func ModificarRelojVector(nombreArchivo string) {
	fmt.Printf("Nombre Archivo a Buscar: %v\n", nombreArchivo)
	fmt.Printf("Dictionario de Relojes: %v - len:% v\n", RelojesVectoresDict, len(RelojesVectoresDict))
	switch FulcrumId {
	case 1:
		RelojesVectoresDict[nombreArchivo].X++
	case 2:
		RelojesVectoresDict[nombreArchivo].Y++
	case 3:
		RelojesVectoresDict[nombreArchivo].Z++
	}

}

func UpdateNumber(cmd *pb.Comando) {

}

func main() {
	var srvAddr string
	if len(os.Args) == 3 {
		if (os.Args[1] == "1") {
			srvAddr = local + port
			curFilesPath = curFiles + filesF1
			FulcrumId = 1;
		} else if (os.Args[1] == "2") {
			srvAddr = local + port2
			curFilesPath = curFiles + filesF2
			FulcrumId = 2;
		} else {
			srvAddr = local + port3
			curFilesPath = curFiles + filesF3
			FulcrumId = 3;
		}
	} else {
		if (os.Args[1] == "1") {
			srvAddr = f1Addrs
			curFilesPath = curFiles + filesF1
			FulcrumId = 1;
		} else if (os.Args[1] == "2") {
			srvAddr = f2Addrs
			curFilesPath = curFiles + filesF2
			FulcrumId = 2;
		} else {
			srvAddr = f3Addrs
			curFilesPath = curFiles + filesF3
			FulcrumId = 3;			
		}
	}
	log.Printf("Address Fulcrum: %v - Id: %v - curFilesPath: %v\n", srvAddr, FulcrumId, curFilesPath)

	lis, err := net.Listen("tcp", srvAddr)

	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}

	s := grpc.NewServer()

	pb.RegisterFulcrumServer(s, &server{})
	log.Printf("Fulcrum Inicializado: escuchando en %v\n", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}