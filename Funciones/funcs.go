package funcs

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

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

//Hay que editar esta función para adaptarla a esta tarea
func InsertarComandoEnRegistro(NombreDelArchivo string, comando string, numeroSoladosRebeldes int32, nombreCiudad string) bool {
	// Open file using READ & WRITE permission.
	var file, err = os.OpenFile(NombreDelArchivo, os.O_APPEND|os.O_WRONLY, 0644)
	if isError(err) {
		return false
	}
	defer file.Close()

	/*
		switch comando {
		case "AddCity":
			// Hacer algo
			// Write some text line-by-line to file.
			var linea string = FormatInt32(jugada) + "\n"
			_, err = file.WriteString(linea)
			if isError(err) {
				return
			}
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
		return false
	}
	return true

}

func FormatInt32(n int32) string {
	return strconv.FormatInt(int64(n), 10)
}
