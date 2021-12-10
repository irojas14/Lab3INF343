package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	pb "github.com/irojas14/Lab3INF343/Proto"
)

func CmdToStr(cmd pb.TipoComando) string {
	var r string
	switch cmd {
	case pb.TipoComando_AddCity:
		r = "AddCity"
	case pb.TipoComando_UpdateName:
		r = "UpdateName"
	case pb.TipoComando_UpdateNumber:
		r = "UpdateNumber"
	case pb.TipoComando_DeleteCity:
		r = "DeleteCity"
	}
	return r
}

func isError(err error) bool {
	if err != nil {
		fmt.Println(err.Error())
	}

	return (err != nil)
}

func CrearTxt(NombreDelArchivo string) {
	// check if file exists
	var _, err = os.Stat(NombreDelArchivo)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(NombreDelArchivo)
		if isError(err) {
			return
		}
		defer file.Close()
	}

	fmt.Println("Se crea el archivo llamado: ", NombreDelArchivo)
}

func IsInServer(nombreRegistro string, ruta string) bool {
	all_files, err := ioutil.ReadDir(ruta)
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range all_files {
		if f.Name() == nombreRegistro {
			return true
		}
	}
	return false
}

//Hay que editar esta funcion para adaptarla a esta tarea
func InsertarComandoEnRegistro(NombreDelArchivo string, comando *pb.Comando) error{
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(NombreDelArchivo, os.O_WRONLY, 0644)
	seReemplazo := false
	if isError(err) {
		log.Fatalf("Error al abrir archivo en Insertar Cambios en Registro. Archivo: %v - Error: %v\n", NombreDelArchivo, err)
		return err
	}
	defer file.Close()
	
	switch comando.Tipo {
		case pb.TipoComando_AddCity:
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + strconv.FormatInt(comando.NuevoValorInt, 10)
			
			// Aqui veremos si linea ya existe dentro del archivo del planeta.
			b, err := ioutil.ReadFile(NombreDelArchivo)
			if err != nil {
				panic(err)
			}

			lines := strings.Split(string(b), "\n")

			fmt.Printf("Principio Lines: %v - len: %v\n", lines, len(lines))

			for i, line := range lines {
				if strings.Contains(line, comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad) {
					fmt.Println("Line contiene Par Planeta-Ciudad")
					lines[i] = linea
					seReemplazo = true
				}
			}

			output := strings.Join(lines, "\n")

			fmt.Printf("Final Output: %v --\n", output)
			
			_, err = file.WriteString(output)
			if isError(err) {
				log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
				return err;
			}
			if !seReemplazo {
				_, err = file.WriteString(linea + "\n")
				if isError(err) {
					log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
					return err;
				}
			}
			seReemplazo=false
		}
			/*
		case "UpdateName":
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = FormatInt32(jugada) + "\n"
			_, err = file.WriteString(linea)
			if isError(err) {
				return
			}
		case "UpdateNumber":
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = FormatInt32(jugada) + "\n"
			_, err = file.WriteString(linea)
			if isError(err) {
				return
			}
		case "DeleteCity":
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = FormatInt32(jugada) + "\n"
			_, err = file.WriteString(linea)
			if isError(err) {
				return
			}
		}
		*/

	/*
		for _, jugada := range JugadasDelJugador {
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = FormatInt32(jugada) + "\n"
			_, err = file.WriteString(linea)
			if isError(err) {
				return
			}
		}
	*/

	err = file.Sync()
	if isError(err) {
		log.Fatalf("Error al hacer 'Sync' en Insertar Cambios en Registro. Archivo: %v\n", NombreDelArchivo)
		return err
	}
	return nil
}

func FormatInt32(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}

func InsertarComandoEnLog(NombreDelArchivo string, comando *pb.Comando) error {
	
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(NombreDelArchivo, os.O_APPEND|os.O_WRONLY, 0644)
	if isError(err) {
		log.Fatalf("Error al abrir archivo en Insertar Cambios en Log. Archivo: %v - Error: %v\n", NombreDelArchivo, err)
		return err
	}
	defer file.Close()

	switch comando.Tipo {
	case pb.TipoComando_AddCity:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + strconv.FormatInt(comando.NuevoValorInt, 10) + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err;
		}
	case pb.TipoComando_DeleteCity:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err;
		}
	case pb.TipoComando_UpdateName:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + comando.NuevoValorStr + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err;
		}
	case pb.TipoComando_UpdateNumber:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + strconv.FormatInt(comando.NuevoValorInt, 10) + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err;
		}	
	}

	err = file.Sync()
	if isError(err) {
		log.Fatalf("Error al hacer 'Sync' en Insertar Cambios en Log. Archivo: %v. Error: %v\n", NombreDelArchivo, err)
		return err
	}
	return nil
}