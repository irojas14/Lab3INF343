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
	port    = ":50500"
	port2   = ":50501"
	port3   = ":50502"
	local   = "localhost"
	local1  = local + port
	local2  = local + port2
	local3  = local + port3
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
	FulcrumId    int
	curAddr string
	curFilesPath string
)

var (
	RelojesVectoresDict map[string]*pb.RelojVector = make(map[string]*pb.RelojVector)
)

type server struct {
	pb.UnimplementedFulcrumServer
}

// METODOS REMOTOS DEL SERVICIO

func (s *server) Comando(ctx context.Context, in *pb.SolicitudComando) (*pb.RespuestaComandoFulcrum, error) {

	fmt.Printf("En Comando Fulcrum id: %v. Comando: %v\n", FulcrumId, in.Cmd.String())

	var RelojVectorRes *pb.RelojVector = nil

	switch in.Cmd.Tipo {
	case pb.TipoComando_AddCity:
		RelojVectorRes = AddCity(in.Cmd)
	case pb.TipoComando_UpdateName:
		RelojVectorRes = UpdateName(in.Cmd)
	case pb.TipoComando_DeleteCity:
		log.Printf("En Delete City: %v\n", FulcrumId)
		RelojVectorRes = DeleteCity(in.Cmd)
	case pb.TipoComando_UpdateNumber:
		RelojVectorRes = UpdateNumber(in.Cmd)
	default:
		log.Printf("Tipo de Comando No Existente: %v\n", in.Cmd.Tipo)
	}
	return &pb.RespuestaComandoFulcrum{RelojVec: RelojVectorRes}, nil
}

func (s *server) GetNumberRebelds(ctx context.Context, in *pb.SolicitudGetNumberRebelds) (*pb.RespuestaGetNumberRebelds, error) {

	log.Printf("Realizando Get Number Rebelds")
	nombreArchivo := in.Consulta.Coord.NombrePlaneta + "_" + "Info"

	existeBool := funcs.IsInServer(nombreArchivo, curFilesPath)

	if existeBool {
		RebelNum, errObtencion := funcs.ObtenerRebels(curFilesPath + "/" + nombreArchivo, in.Consulta.Coord)

		if errObtencion != nil {
			log.Printf("Ocurrió un error al buscar los rebeldes en Coord: %v y Fulcrum: %v: Error: %v - Cód: %v\n", in.Consulta.Coord, FulcrumId, errObtencion, RebelNum)
			return nil, errObtencion
		}

		if RebelNum == -3 {
			log.Printf("No se encontró la ciudad %v en el planeta %v en el Fulcrum %v\n", in.Consulta.Coord.NombreCiudad, in.Consulta.Coord.NombrePlaneta, FulcrumId)
			return &pb.RespuestaGetNumberRebelds{
				Consulta: &pb.Consulta {NumRebels: -3},
			}, nil
		}

		if RelojesVectoresDict[nombreArchivo] == nil {
			RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
				Nombre: nombreArchivo,
				X : 0, Y: 0, Z: 0,
			}
		}

		return &pb.RespuestaGetNumberRebelds{
			Consulta: &pb.Consulta{
				ArchivoName: nombreArchivo,
				FulcrumDir: curAddr,
				NumRebels: RebelNum,
				RelojVec: RelojesVectoresDict[nombreArchivo],
			},
		}, nil

	} else {
		log.Printf("Coord %v no presente en Fulcrum: %v\n", in.Consulta.Coord, FulcrumId)
		return &pb.RespuestaGetNumberRebelds{}, nil
	}
}

// FUNCIONES AUXILIARES

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
				X:      0, Y: 0, Z: 0,
			}
		}

		err := funcs.InsertarComandoEnRegistro(curFilesPath+"/"+nombreArchivo, cmd) // WIP
		if err != nil {
			return nil
		}

		// Registramos el cambio en el Log asociado
		err = funcs.InsertarComandoEnLog(curFilesPath+"/"+nombreArchivoLog, cmd)
		if err != nil {
			return nil
		}
	} else {
		// Si no existe

		// Creamos su Reloj de Vectores
		RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
			Nombre: nombreArchivo,
			X:      0, Y: 0, Z: 0,
		}

		// Creamos el Archivo de Registro Planetario
		funcs.CrearTxt(curFilesPath + "/" + nombreArchivo) // creamos el archivo de Info

		// Insertamos el comando en el Registro Planetario
		err := funcs.InsertarComandoEnRegistro(curFilesPath+"/"+nombreArchivo, cmd)
		if err != nil {
			return nil
		}

		// Creamos el archivo de Log asociado
		funcs.CrearTxt(curFilesPath + "/" + nombreArchivoLog) // creamos el archivo Log

		// Registramos el cambio en el Log
		err = funcs.InsertarComandoEnLog(curFilesPath+"/"+nombreArchivoLog, cmd) //
		if err != nil {
			return nil
		}
	}
	// Modificamos el reloj de vectores correspondiente para reflejar el cambio
	ModificarRelojVector(nombreArchivo)

	return RelojesVectoresDict[nombreArchivo]
}

func UpdateName(cmd *pb.Comando) *pb.RelojVector {

	nombreArchivo := cmd.Coord.NombrePlaneta + "_" + "Info"
	nombreArchivoLog := cmd.Coord.NombrePlaneta + "_" + "Log"

	// Vemos que el archivo existe
	if funcs.IsInServer(nombreArchivo, curFilesPath) {

		// Si existe, vemos si existe la ciudad para actualizarla
		b, err := funcs.CambiarNombreCiudad(curFilesPath+"/"+nombreArchivo, cmd)

		// si no existe la ciudad, retornamos nil
		if !b || err != nil {
			return nil
		}

		// El cambio se llevo a cabo de los Registros Planetarios. Registramos el cambio en Log
		err = funcs.InsertarComandoEnLog(curFilesPath+"/"+nombreArchivoLog, cmd) //

		if err != nil {
			return nil
		}

		if RelojesVectoresDict[nombreArchivo] == nil {
			RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
				Nombre: nombreArchivo,
				X:      0, Y: 0, Z: 0,
			}
		}

		ModificarRelojVector(nombreArchivo)
	} else {
		return nil
	}
	return RelojesVectoresDict[nombreArchivo]
}

func DeleteCity(cmd *pb.Comando) *pb.RelojVector {
	nombreArchivo := cmd.Coord.NombrePlaneta + "_" + "Info"
	nombreArchivoLog := cmd.Coord.NombrePlaneta + "_" + "Log"

	// Vemos que el archivo existe
	if funcs.IsInServer(nombreArchivo, curFilesPath) {

		log.Printf("El archivo del planeta sí existe: Fulcrum %v - Planeta: %v\n", FulcrumId, cmd.Coord.NombrePlaneta)
		// Si existe, vemos si existe la ciudad para actualizarla
		b, err := funcs.BorrarCiudad(curFilesPath+"/"+nombreArchivo, cmd)

		// si no existe la ciudad, retornamos nil
		if !b || err != nil {
			log.Printf("La ciudad no existe: %v\n", cmd.Coord.NombreCiudad)
			return nil
		}

		// El cambio se llevo a cabo de los Registros Planetarios. Registramos el cambio en Log
		err = funcs.InsertarComandoEnLog(curFilesPath+"/"+nombreArchivoLog, cmd) //

		if err != nil {
			return nil
		}

		if RelojesVectoresDict[nombreArchivo] == nil {
			RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
				Nombre: nombreArchivo,
				X:      0, Y: 0, Z: 0,
			}
		}

		ModificarRelojVector(nombreArchivo)
	} else {
		log.Printf("El archivo del planeta no existe: %v\n", cmd.Coord.NombrePlaneta)
		return nil
	}
	return RelojesVectoresDict[nombreArchivo]
}

// -------------------Maxi porfa corrobora que está bien esto------------------------------------
func UpdateNumber(cmd *pb.Comando) *pb.RelojVector {
	nombreArchivo := cmd.Coord.NombrePlaneta + "_" + "Info"
	nombreArchivoLog := cmd.Coord.NombrePlaneta + "_" + "Log"

	// Vemos que el archivo existe
	if funcs.IsInServer(nombreArchivo, curFilesPath) {

		// Si existe, vemos si existe la ciudad para actualizarla
		b, err := funcs.CambiarNumberoDeSoldados(curFilesPath+"/"+nombreArchivo, cmd)

		// si no existe la ciudad, retornamos nil
		if !b || err != nil {
			return nil
		}

		// El cambio se llevo a cabo de los Registros Planetarios. Registramos el cambio en Log
		err = funcs.InsertarComandoEnLog(curFilesPath+"/"+nombreArchivoLog, cmd) //

		if err != nil {
			return nil
		}

		if RelojesVectoresDict[nombreArchivo] == nil {
			RelojesVectoresDict[nombreArchivo] = &pb.RelojVector{
				Nombre: nombreArchivo,
				X:      0, Y: 0, Z: 0,
			}
		}

		ModificarRelojVector(nombreArchivo)
	} else {
		return nil
	}
	return RelojesVectoresDict[nombreArchivo]
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

func main() {
	var srvAddr string
	if len(os.Args) == 3 {
		if os.Args[1] == "1" {
			srvAddr = local + port
			curFilesPath = curFiles + filesF1
			FulcrumId = 1
		} else if os.Args[1] == "2" {
			srvAddr = local + port2
			curFilesPath = curFiles + filesF2
			FulcrumId = 2
		} else {
			srvAddr = local + port3
			curFilesPath = curFiles + filesF3
			FulcrumId = 3
		}
	} else {
		if os.Args[1] == "1" {
			srvAddr = f1Addrs
			curFilesPath = curFiles + filesF1
			FulcrumId = 1
		} else if os.Args[1] == "2" {
			srvAddr = f2Addrs
			curFilesPath = curFiles + filesF2
			FulcrumId = 2
		} else {
			srvAddr = f3Addrs
			curFilesPath = curFiles + filesF3
			FulcrumId = 3
		}
	}
	curAddr = srvAddr
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
