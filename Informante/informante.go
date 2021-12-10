package main

import (
	"context"
	"fmt"
	"log"
	"os"

	funcs "github.com/irojas14/Lab3INF343/Funciones"
	pb "github.com/irojas14/Lab3INF343/Proto"
	"google.golang.org/grpc"
)

const (
	mos_port = ":50054"
	local    = "localhost" + mos_port
	mos_addr = "dist149.inf.santiago.usm.cl" + mos_port
)

var (
	addr string
)

// Variables de los Informantes

type Cambio struct {
	archivo_name string
	comando      *pb.Comando
	reloj_vec    *pb.RelojVector
	fulcrum_dir  string
}

var (
	Tipo    pb.ActorType
	cambios []*Cambio = make([]*Cambio, 0)
)

func main() {

	fmt.Print("INFORMANTE INICIADO")
	addr = mos_addr
	if len(os.Args) == 2 {
		addr = local
	}
	Consola()
}

// FUNCIONES CONSOLA

func ConsolaAddCity() {
	var planeta string
	var ciudad string
	var valorInt int64
	fmt.Print("Ingrese Nombre Planeta: ")
	fmt.Scanln(&planeta)
	fmt.Print("Ingrese Nombre Ciudad: ")
	fmt.Scanln(&ciudad)
	fmt.Print("Ingrese Número Rebeldes: ")
	fmt.Scanln(&valorInt)

	coord := &pb.Ubicacion{
		NombrePlaneta: planeta,
		NombreCiudad: ciudad,
	}
	AddCity(coord, valorInt)

	log.Println("Comando 'AddCity' Terminado")
}

// FUNCIONES INFORMANTE

func CreateBaseComandoAndConn(tipoComando pb.TipoComando,coord *pb.Ubicacion) (*pb.Comando, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v\n", err)
		return nil, nil, err
	}
	cmd := CreateBaseCmd(tipoComando, coord)
	return cmd, conn, nil
}

func AddCity(coord *pb.Ubicacion, nuevo_valor int64) {
	cmd, conn, err := CreateBaseComandoAndConn(pb.TipoComando_AddCity, coord)
	if err != nil {
		return
	}
	defer conn.Close()

	cmd.NuevoValorInt = nuevo_valor
	mc := pb.NewMosEisleyClient(conn)
	dirFulcrum, err := Execute(mc, cmd)
	conn.Close()	
	log.Println("primer 'conn' cerrado")

	log.Printf("Dirección Fulcrum Recibida: dir: %v\n", dirFulcrum)
	if (err != nil) {
		log.Fatalf("Error al ejecutar el comando hacia Mos Eisley: err: %v\n", err)
		return
	}

	log.Printf("Abriendo Sgte Conexión en dirección: %v\n", dirFulcrum)
	connFulcrum, errFulcrum := grpc.Dial(dirFulcrum, grpc.WithInsecure(), grpc.WithBlock())
	if errFulcrum != nil {
		log.Fatalf("No se pudo conectar: %v\n", err)
		return
	}
	defer connFulcrum.Close()

	fc := pb.NewFulcrumClient(connFulcrum)

	log.Printf("'FC' cliente creado, a realizar fc.comando")
	rComando, errComando := fc.Comando(context.Background(), &pb.SolicitudComando{
		Cmd: cmd,
	})

	if errComando != nil {
		log.Fatalf("Error al realizar el comando: %v\n", errComando)
		return
	}

	log.Printf("AddCity Realizado. Agregando al arreglo de Cambios: Respueta: %v\n", rComando)
	cambios = append(cambios, &Cambio{
		archivo_name: rComando.RelojVec.Nombre,
		comando: cmd,
		reloj_vec: rComando.RelojVec,
		fulcrum_dir: dirFulcrum,
	})
	
	if (cambios == nil) {
		log.Printf("'cambios' es null")
	} else if (len(cambios) == 0){
		log.Printf("Len de 'cambios' es 0")
	} else {
		log.Printf("AddCity Realizado con Éxito: Nuevo Cambio: %v\n", cambios[len(cambios) - 1])
	}
}

func AddCityNonVal(coord *pb.Ubicacion) {
	cmd, conn, err := CreateBaseComandoAndConn(pb.TipoComando_AddCity, coord)
	if err != nil {
		return
	}
	defer conn.Close()

	c := pb.NewMosEisleyClient(conn)
	Execute(c, cmd)
}

func UpdateName(coord *pb.Ubicacion, nuevo_valor string) {

}

func UpdateNumber(coord *pb.Ubicacion, nuevo_valor int64) {

}

func Execute(c pb.MosEisleyClient, cmd *pb.Comando) (string, error) {
	r, err := c.Comando(context.Background(), &pb.SolicitudComando{
		Cmd: cmd,
	})
	if err != nil {
		log.Fatalf("Error al realizar el comando: %v\n", err)
		return "", err
	}
	log.Printf("Results: %v\n", r)
	return r.DirFulcrum, nil
}

func CreateBaseCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion) *pb.Comando {
	return &pb.Comando{
		Tipo:   tipoCmd,
		Nombre: funcs.CmdToStr(tipoCmd),
		Coord:  coord,
	}
}

func Consola() {
	for {
		ShowConsola()

		var option string 
		fmt.Print("Ingrese Opción: ")
		fmt.Scanln(&option)

		exitBool := ConsolaProcesamiento(option)
		if (exitBool) {
			break
		}
	}
}

func ShowConsola() {
    fmt.Println("========== CONSOLA INFORMANTE ==========")

	// Opciones
    fmt.Println("COMANDO 'AddCity': PRESIONAR 'A' + 'ENTER'")
}

func ConsolaProcesamiento(option string) bool {
	if option == "A" || option == "a" {
		ConsolaAddCity()
	} else if option == "E" || option == "e" {
		return true;
	} else {
		fmt.Printf("Option %v no válida\n", option)
	}
	return false
}

/*
func CreateAddCityValueCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion, nuevo_valor int64) *pb.Comando {
	cmd := CreateBaseCmd(tipoCmd, coord)
	cmd.NuevoValorInt = nuevo_valor
	return cmd
}

func CreateAddCityNonValueCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion) *pb.Comando {
	cmd := CreateBaseCmd(tipoCmd, coord)
	return cmd
}



func CreateUpdateNameCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion, nuevo_valor string) *pb.Comando{
	cmd := CreateBaseCmd(tipoCmd, coord)
	cmd.NuevoValorStr = nuevo_valor
	return cmd
}

func CreateUpdateNumberCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion, nuevo_valor int64) *pb.Comando {
	cmd := CreateBaseCmd(tipoCmd, coord)
	cmd.NuevoValorInt = nuevo_valor
	return cmd
}

func CreateDeleteCityCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion) *pb.Comando {
	cmd := CreateBaseCmd(tipoCmd, coord)
	return cmd
}
*/
