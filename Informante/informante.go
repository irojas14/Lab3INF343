package main

import (
	"context"
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
	comando      pb.Comando
	reloj_vec    *pb.RelojVector
	fulcrum_dir  string
}

var (
	Tipo    pb.ActorType
	cambios []Cambio
)

func main() {

	addr = mos_addr
	if len(os.Args) == 2 {
		addr = local
	}

	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v\n", err)
	}
	defer conn.Close()
	c := pb.NewMosEisleyClient(conn)

	log.Printf("%v\n", c)
}

// FUNCIONES INFORMANTE

func CreateBaseComandoAndConn(coord *pb.Ubicacion) (*pb.Comando, *grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("No se pudo conectar: %v\n", err)
		return nil, nil, err
	}
	cmd := CreateBaseCmd(pb.TipoComando_AddCity, coord)
	return cmd, conn, nil
}

func AddCity(coord *pb.Ubicacion, nuevo_valor int64) {
	cmd, conn, err := CreateBaseComandoAndConn(coord)
	if err != nil {
		return
	}
	defer conn.Close()

	cmd.NuevoValorInt = nuevo_valor
	c := pb.NewMosEisleyClient(conn)
	Execute(c, cmd)
}

func AddCityNonVal(coord *pb.Ubicacion) {
	cmd, conn, err := CreateBaseComandoAndConn(coord)
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

func Execute(c pb.MosEisleyClient, cmd *pb.Comando) error {
	r, err := c.Comando(context.Background(), &pb.SolicitudComando{
		Cmd: cmd,
	})
	if err != nil {
		log.Fatalf("Error al realizar el comando: %v\n", err)
		return err
	}
	log.Printf("Results: %v\n", r)
	return nil
}

func CreateBaseCmd(tipoCmd pb.TipoComando, coord *pb.Ubicacion) *pb.Comando {
	return &pb.Comando{
		Tipo:   tipoCmd,
		Nombre: funcs.CmdToStr(tipoCmd),
		Coord:  coord,
	}
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
