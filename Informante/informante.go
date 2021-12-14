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
	mos_addr = "dist150.inf.santiago.usm.cl" + mos_port //Ver si esto es correcto
)

var (
	addr string
)

// Variables de los Informantes
var (
	cambios []*pb.Cambio = make([]*pb.Cambio, 0)
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
		NombreCiudad:  ciudad,
	}
	AddCity(coord, valorInt)

	log.Println("Comando 'AddCity' Terminado")
}

func ConsolaUpdateName() {
	var planeta string
	var ciudad string
	var valorStr string
	fmt.Print("Ingrese Nombre Planeta: ")
	fmt.Scanln(&planeta)
	fmt.Print("Ingrese Nombre Ciudad: ")
	fmt.Scanln(&ciudad)
	fmt.Print("Ingrese Nuevo Nombre: ")
	fmt.Scanln(&valorStr)

	coord := &pb.Ubicacion{
		NombrePlaneta: planeta,
		NombreCiudad:  ciudad,
	}
	UpdateName(coord, valorStr)

	log.Println("Comando 'UpdateName' Terminado")
}

// Falta completar
func ConsolaDeleteCity() {
	var planeta string
	var ciudad string
	fmt.Print("Ingrese Nombre Planeta: ")
	fmt.Scanln(&planeta)
	fmt.Print("Ingrese Nombre Ciudad: ")
	fmt.Scanln(&ciudad)
	
	coord := &pb.Ubicacion{
		NombrePlaneta: planeta,
		NombreCiudad:  ciudad,
	}

	DeleteCity(coord)

	log.Println("Comando 'DeleteCity' Terminado")

}

func ConsolaUpdateNumber() {
	var planeta string
	var ciudad string
	var valorInt int64
	fmt.Print("Ingrese Nombre Planeta: ")
	fmt.Scanln(&planeta)
	fmt.Print("Ingrese Nombre Ciudad: ")
	fmt.Scanln(&ciudad)
	fmt.Print("Ingrese Nuevo Número: ")
	fmt.Scanln(&valorInt)

	coord := &pb.Ubicacion{
		NombrePlaneta: planeta,
		NombreCiudad:  ciudad,
	}
	UpdateNumber(coord, valorInt)

	log.Println("Comando 'UpdateName' Terminado")
}

func ConsolaGetNumRebels() {
	var planeta string
	var ciudad string
	fmt.Print("Ingrese Nombre Planeta: ")
	fmt.Scanln(&planeta)
	fmt.Print("Ingrese Nombre Ciudad: ")
	fmt.Scanln(&ciudad)

	coord := &pb.Ubicacion{
		NombrePlaneta: planeta,
		NombreCiudad:  ciudad,
	}

	GetNumberRebels(coord)

	log.Println("Comando 'Get Num Rebels' Terminado")
}

// FUNCIONES INFORMANTE

func CreateBaseComandoAndConn(tipoComando pb.TipoComando, coord *pb.Ubicacion) (*pb.Comando, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v\n", err)
		return nil, nil, err
	}
	cmd := CreateBaseCmd(tipoComando, coord)
	return cmd, conn, nil
}

func AddCity(coord *pb.Ubicacion, nuevo_valor int64) {
	// Realizar RPC a Mos Eisley
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
	if err != nil {
		log.Fatalf("Error al ejecutar el comando hacia Mos Eisley: err: %v\n", err)
		return
	}

	// Realizar RPC a Fulcrum
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

	ResultadosComando(cmd, rComando, dirFulcrum)
}

func UpdateName(coord *pb.Ubicacion, nuevo_valor string) {
	// Realizar RPC a Mos Eisley
	cmd, conn, err := CreateBaseComandoAndConn(pb.TipoComando_UpdateName, coord)
	if err != nil {
		return
	}
	defer conn.Close()
	cmd.NuevoValorStr = nuevo_valor
	mc := pb.NewMosEisleyClient(conn)
	dirFulcrum, err := Execute(mc, cmd)
	conn.Close()
	log.Println("primer 'conn' cerrado")

	log.Printf("Dirección Fulcrum Recibida: dir: %v\n", dirFulcrum)
	if err != nil {
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

	// Realizar RPC a Fulcrum
	fc := pb.NewFulcrumClient(connFulcrum)

	log.Printf("'FC' cliente creado, a realizar fc.comando")
	rComando, errComando := fc.Comando(context.Background(), &pb.SolicitudComando{
		Cmd: cmd,
	})

	if errComando != nil {
		log.Fatalf("Error al realizar el comando: %v\n", errComando)
		return
	}

	ResultadosComando(cmd, rComando, dirFulcrum)
}

func UpdateNumber(coord *pb.Ubicacion, nuevo_valor int64) {
	// Realizar RPC a Mos Eisley
	cmd, conn, err := CreateBaseComandoAndConn(pb.TipoComando_UpdateNumber, coord)
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
	if err != nil {
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

	// Realizar RPC a Fulcrum
	fc := pb.NewFulcrumClient(connFulcrum)

	log.Printf("'FC' cliente creado, a realizar fc.comando")
	rComando, errComando := fc.Comando(context.Background(), &pb.SolicitudComando{
		Cmd: cmd,
	})

	if errComando != nil {
		log.Fatalf("Error al realizar el comando: %v\n", errComando)
		return
	}

	ResultadosComando(cmd, rComando, dirFulcrum)
}

func DeleteCity(coord *pb.Ubicacion) {
		// Realizar RPC a Mos Eisley
		cmd, conn, err := CreateBaseComandoAndConn(pb.TipoComando_DeleteCity, coord)
		if err != nil {
			return
		}
		defer conn.Close()
		mc := pb.NewMosEisleyClient(conn)
		dirFulcrum, err := Execute(mc, cmd)
		conn.Close()
		log.Println("primer 'conn' cerrado")
	
		// Realizar RPC a FULCRUM
		log.Printf("Dirección Fulcrum Recibida: dir: %v\n", dirFulcrum)
		if err != nil {
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
	
		// Realizar RPC a Fulcrum
		fc := pb.NewFulcrumClient(connFulcrum)
	
		log.Printf("'FC' cliente creado, a realizar fc.comando")
		rComando, errComando := fc.Comando(context.Background(), &pb.SolicitudComando{
			Cmd: cmd,
		})
	
		if errComando != nil {
			log.Fatalf("Error al realizar el comando: %v\n", errComando)
			return
		}
	
		ResultadosComando(cmd, rComando, dirFulcrum)
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

func ResultadosComando(cmd *pb.Comando, rComando *pb.RespuestaComandoFulcrum, dirFulcrum string) {
	
	if rComando == nil || rComando.RelojVec == nil {
		log.Printf("Respuesta Vacía. No se llevo a cabo ningún cambio")
		return
	}

	log.Printf( cmd.Nombre + " Realizado. Respueta: %v\n", rComando)
	if rComando != nil {
		log.Printf( cmd.Nombre + " Realizado. Agregando al arreglo de Cambios: Respueta: %v\n", rComando)
		cambios = append(cambios, &pb.Cambio{
			ArchivoName: rComando.RelojVec.Nombre,
			Cmd:      cmd,
			RelojVec:    rComando.RelojVec,
			FulcrumDir:  dirFulcrum,
		})
		log.Printf(cmd.Nombre + " Realizado con Éxito: Nuevo Cambio: %v\n", cambios[len(cambios)-1])
	} else {
		log.Printf(cmd.Nombre + " Fracasado. Respuesta Nula: %v\n", rComando)
	}
}

func GetNumberRebels(coord *pb.Ubicacion) {

	cambio := &pb.Cambio{
		Cmd: &pb.Comando{
			Coord: coord,
		},
	}
	for i := len(cambios) - 1; i >= 0; i-- {
		cm := cambios[i]
		if cm.Cmd.Coord.NombrePlaneta == coord.NombrePlaneta && cm.Cmd.Coord.NombreCiudad == coord.NombreCiudad {
			cambio = cambios[i]
			break
		}
	}

	conn, errConn := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())

	if errConn != nil {
		log.Fatalf("Error al conectarse para realizar un GetNumberRebel: Error: %v\n", errConn)
		return
	}
	defer conn.Close()

	c := pb.NewMosEisleyClient(conn)

	r, errSol := c.GetNumberRebeldsInformante(context.Background(), &pb.SolicitudGetNumRebelsInformante{
		Cambio: cambio,
	})

	if errSol != nil {
		log.Fatalf("Error al realizar un GetNumberRebel: Error: %v\n", errSol)
		return		
	}
	fmt.Printf("Resultado Obtenido: %v\n", r)
}

func Consola() {
	for {
		ShowConsola()

		var option string
		fmt.Print("Ingrese Opción: ")
		fmt.Scanln(&option)

		exitBool := ConsolaProcesamiento(option)
		if exitBool {
			break
		}
	}
}

func ShowConsola() {
	fmt.Println("========== CONSOLA INFORMANTE ==========")

	// Opciones
	fmt.Println("COMANDO 'AddCity': PRESIONAR 'A' + 'ENTER'")
	fmt.Println("COMANDO 'UpdateName': PRESIONAR 'N' + 'ENTER'")
	fmt.Println("COMANDO 'UpdateNumber': PRESIONAR 'F' + 'ENTER'")
	fmt.Println("COMANDO 'DeleteCity': PRESIONAR 'D' + 'ENTER'")
	fmt.Println("PREGUNTAR Por Rebeldes: PRESIONAR 'p' + 'ENTER")
	fmt.Println("SALIR: PRESIONAR 'E' + 'ENTER'")
}

func ConsolaProcesamiento(option string) bool {
	if option == "A" || option == "a" {
		ConsolaAddCity()
	} else if option == "E" || option == "e" {
		return true
	} else if option == "N" || option == "n" {
		ConsolaUpdateName()
	} else if option == "F" || option == "f" {
		ConsolaUpdateNumber()
	} else if option == "D" || option == "d" {
		ConsolaDeleteCity()
	} else if option == "P" || option == "p" {
		ConsolaGetNumRebels()
	}else {
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
