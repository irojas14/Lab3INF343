package main

import (
	"context"
	"fmt"
	"log"
	"os"

	pb "github.com/irojas14/Lab3INF343/Proto"
	"google.golang.org/grpc"
)

const (
	mos_eisley_port = ":50054"
	mos_eisley_addr_local = "localhost" + mos_eisley_port
	mos_eisley_addr = "dist150.inf.santiago.usm.cl" + mos_eisley_port
)

var (
	cur_mos_eisley string
)

type Consulta struct {
	archivo_name string
	coord        *pb.Ubicacion
	rebel_num    int64
	reloj_vec    *pb.RelojVector
	fulcrum_dir  string
}

var (
	consultas []Consulta
)

// FUNCION LEIA

func GetNumberRebelds(coord *pb.Ubicacion) error {
	
	// Conectarse y pedir la informacion a MOS EISLEY

	// Creamos la conexion y el comando
	connMos, errConnMos := grpc.Dial(cur_mos_eisley, grpc.WithInsecure(), grpc.WithBlock())
	if errConnMos != nil {
		log.Fatalf("No se pudo conectar: %v\n", errConnMos)
		return nil
	}
	defer connMos.Close()

	// Creamos el MosEisleyClient
	c := pb.NewMosEisleyClient(connMos);

	// Realizamos el llamado remoto
	rRebelds, errRebels := c.GetNumberRebelds(context.Background(), &pb.SolicitudGetNumberRebelds{
		Coord: coord,
	})

	if (errRebels != nil) {
		log.Fatalf("Error al pedir el número de rebeldes de %v %v. Error: %v\n", coord.NombrePlaneta, coord.NombreCiudad, errRebels)
		return nil
	}
	if rRebelds != nil {

		if rRebelds.NumRebels != -3 {
		
			if (rRebelds.RelojVec != nil) {

				consultas = append(consultas, Consulta{
				archivo_name: rRebelds.ArchivoName,
				coord:        coord,
				rebel_num:    rRebelds.NumRebels,
				reloj_vec:    rRebelds.RelojVec,
				fulcrum_dir:  rRebelds.FulcrumDir,
				})

				log.Printf("Los Rebeldes en la ciudad de %v en el planeta %v son %v! - Reloj Vector: %v - Fulcrum: %v\n", 
				coord.NombreCiudad, coord.NombrePlaneta, rRebelds.NumRebels, rRebelds.RelojVec, rRebelds.FulcrumDir)

			} else {
				log.Fatalf("No se encontró reloj de Vectores")	
			}
		} else {
			log.Printf("No se encontró la ciudad solicitada")
		}

	} else {
		log.Printf("La Respuesta fue vacío. No se obtuvo Información")
	}
	return nil
}

// FUNCIONES CONSOLA
func ConsolaGetNumberRebelds() {
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

	GetNumberRebelds(coord);

	log.Println("Comando GetNumberRebelds Terminado")
}

// CONSOLA LEIA

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
	fmt.Println("========== CONSOLA LEIA ==========")

	// Opciones
	fmt.Println("COMANDO 'Get Number Rebelds': PRESIONAR 'A' + 'ENTER'")
	fmt.Println("COMANDO Para Salir: PRESIONAR 'E' + 'ENTER'")
}

func ConsolaProcesamiento(option string) bool {
	if option == "A" || option == "a" {
		ConsolaGetNumberRebelds()
	} else if option == "E" || option == "e" {
		return true
	} else {
		fmt.Printf("Option %v no válida\n", option)
	}
	return false
}

func main() {
	if len(os.Args) == 2 {
		cur_mos_eisley = mos_eisley_addr_local
	} else {
		cur_mos_eisley = mos_eisley_addr
	}

	log.Printf("Dir Mos Eisley: %v\n", cur_mos_eisley)

	Consola()
}
