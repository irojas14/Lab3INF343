package funcs

import (
	"bufio"
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

func InsertarComandoEnRegistro(NombreDelArchivo string, comando *pb.Comando) error {
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
			if strings.Contains(line, comando.Coord.NombrePlaneta+" "+comando.Coord.NombreCiudad) {
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
			return err
		}
		if !seReemplazo {
			_, err = file.WriteString(linea + "\n")
			if isError(err) {
				log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
				return err
			}
		}
		seReemplazo = false
	}

	err = file.Sync()
	if isError(err) {
		log.Fatalf("Error al hacer 'Sync' en Insertar Cambios en Registro. Archivo: %v\n", NombreDelArchivo)
		return err
	}
	fix_file(NombreDelArchivo)
	return nil
}

func CambiarNombreCiudad(NombreDelArchivo string, comando *pb.Comando) (bool, error) {
	var file, err = os.OpenFile(NombreDelArchivo, os.O_WRONLY, 0644)

	if isError(err) {
		log.Fatalf("Error al abrir archivo en Insertar Cambios en Registro. Archivo: %v - Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	defer file.Close()

	// Aqui veremos si linea ya existe dentro del archivo del planeta.
	b, err := ioutil.ReadFile(NombreDelArchivo)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	fmt.Printf("Principio Lines: %v - len: %v\n", lines, len(lines))

	var nuevaLinea string = ""
	for i, line := range lines {
		if strings.Contains(line, comando.Coord.NombreCiudad) {
			fmt.Println("La ciudad existe")
			items := strings.Split(line, " ")
			log.Printf("Items: %v: len: %v\n", items, len(items))
			log.Printf("Comando Nuevo Valor: %v\n", comando.NuevoValorStr)
			items[1] = comando.NuevoValorStr
			nuevaLinea = strings.Join(items, " ")
			log.Printf("Nueva Línea: %v\n", nuevaLinea)
			lines[i] = nuevaLinea
		}
	}

	if nuevaLinea == "" {
		fmt.Print("No Existe la Ciudad")
		return false, nil
	}

	output := strings.Join(lines, "\n")

	fmt.Printf("Final Output: %v --\n", output)

	_, err = file.WriteString(output)
	if isError(err) {
		log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	fix_file(NombreDelArchivo)
	return true, nil
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
			return err
		}
	case pb.TipoComando_DeleteCity:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err
		}
	case pb.TipoComando_UpdateName:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + comando.NuevoValorStr + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err
		}
	case pb.TipoComando_UpdateNumber:
		linea := comando.Nombre + " " + comando.Coord.NombrePlaneta + " " + comando.Coord.NombreCiudad + " " + strconv.FormatInt(comando.NuevoValorInt, 10) + "\n"
		_, err = file.WriteString(linea)
		if isError(err) {
			log.Fatalf("Error al escribir linea de archivo log %v. Error: %v\n", NombreDelArchivo, err)
			return err
		}
	}

	err = file.Sync()
	if isError(err) {
		log.Fatalf("Error al hacer 'Sync' en Insertar Cambios en Log. Archivo: %v. Error: %v\n", NombreDelArchivo, err)
		return err
	}
	return nil
}

func BorrarCiudad(NombreDelArchivo string, comando *pb.Comando) (bool, error) {

	var file, err = os.OpenFile(NombreDelArchivo, os.O_RDWR, 0644)
	file_new, err_new := os.OpenFile("temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}
	if err_new != nil {
		return false, err_new
	}
	defer file.Close()

	cityExist := false
	datawriter := bufio.NewWriter(file_new)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return false, err
	}
	for _, line := range lines {
		if strings.Contains(line, comando.Coord.NombreCiudad) {
			cityExist = true
			continue
		}
		datawriter.WriteString(line + "\n")
	}
	datawriter.Flush()
	file_new.Close()
	file.Close()
	e1 := os.Remove(NombreDelArchivo)
	if e1 != nil {
		return false, e1
	}
	e := os.Rename("temp.txt", NombreDelArchivo)
	if e != nil {
		return false, e
	}
	fix_file(NombreDelArchivo)
	return cityExist, nil
}

func CambiarNumberoDeSoldados(NombreDelArchivo string, comando *pb.Comando) (bool, error) {
	var file, err = os.OpenFile(NombreDelArchivo, os.O_WRONLY, 0644)

	if isError(err) {
		log.Fatalf("Error al abrir archivo en Insertar Cambios en Registro. Archivo: %v - Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	defer file.Close()

	// Aqui veremos si linea ya existe dentro del archivo del planeta.
	b, err := ioutil.ReadFile(NombreDelArchivo)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	fmt.Printf("Principio Lines: %v - len: %v\n", lines, len(lines))

	var nuevaLinea string = ""
	for i, line := range lines {
		if strings.Contains(line, comando.Coord.NombreCiudad) {
			fmt.Println("La ciudad existe")
			items := strings.Split(line, " ")
			log.Printf("Items: %v: len: %v\n", items, len(items))
			log.Printf("Comando Nuevo Valor: %v\n", comando.NuevoValorInt)
			items[2] = strconv.FormatInt(comando.NuevoValorInt, 10)
			nuevaLinea = strings.Join(items, " ")
			log.Printf("Nueva Línea: %v\n", nuevaLinea)
			lines[i] = nuevaLinea
		}
	}

	if nuevaLinea == "" {
		fmt.Print("No Existe la Ciudad")
		return false, nil
	}

	output := strings.Join(lines, "\n")

	fmt.Printf("Final Output: %v --\n", output)

	_, err = file.WriteString(output)
	if isError(err) {
		log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	fix_file(NombreDelArchivo)
	return true, nil
}

func fix_file(NombreDelArchivo string) error {
	file, err := os.Open("./files/" + NombreDelArchivo)
	file_new, err_new := os.OpenFile("./files/temp.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	if err_new != nil {
		return err_new
	}
	defer file.Close()
	datawriter := bufio.NewWriter(file_new)
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	for _, line := range lines {
		if line == "\n" || line == "" || line == " " {
			continue
		}
		datawriter.WriteString(line + "\n")
	}
	datawriter.Flush()
	file_new.Close()
	file.Close()
	e1 := os.Remove("./files/" + NombreDelArchivo)
	if e1 != nil {
		return e1
	}
	e := os.Rename("./files/temp.txt", "./files/"+NombreDelArchivo)
	if e != nil {
		return e
	}
	return nil
}

/*
func BorrarCiudad(NombreDelArchivo string, comando *pb.Comando) (bool, error) {

	var file, err = os.OpenFile(NombreDelArchivo, os.O_WRONLY, 0644)

	if isError(err) {
		log.Fatalf("Error al abrir archivo en Insertar Cambios en Registro. Archivo: %v - Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	defer file.Close()

	// Aqui veremos si linea ya existe dentro del archivo del planeta.
	b, err := ioutil.ReadFile(NombreDelArchivo)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")

	fmt.Printf("Principio Lines: %v - len: %v\n", lines, len(lines))

	var deleteFlag string = ""
	for i, line := range lines {
		if strings.Contains(line, comando.Coord.NombreCiudad) {
			lines[i] = ""
			deleteFlag = " "
		}
	}

	if deleteFlag == "" {
		fmt.Print("No Existe la Ciudad")
		return false, nil
	}

	output := strings.Join(lines, "\n")

	fmt.Printf("Final Output: %v --\n", output)

	_, err = file.WriteString(output)
	if isError(err) {
		log.Fatalf("Error al escribir linea de archivo %v. Error: %v\n", NombreDelArchivo, err)
		return false, err
	}
	return true, nil
}
*/
