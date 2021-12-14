// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.18.1
// source: Proto/starwars.proto

package Proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type TipoComando int32

const (
	TipoComando_AddCity      TipoComando = 0
	TipoComando_UpdateName   TipoComando = 1
	TipoComando_UpdateNumber TipoComando = 2
	TipoComando_DeleteCity   TipoComando = 3
)

// Enum value maps for TipoComando.
var (
	TipoComando_name = map[int32]string{
		0: "AddCity",
		1: "UpdateName",
		2: "UpdateNumber",
		3: "DeleteCity",
	}
	TipoComando_value = map[string]int32{
		"AddCity":      0,
		"UpdateName":   1,
		"UpdateNumber": 2,
		"DeleteCity":   3,
	}
)

func (x TipoComando) Enum() *TipoComando {
	p := new(TipoComando)
	*p = x
	return p
}

func (x TipoComando) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (TipoComando) Descriptor() protoreflect.EnumDescriptor {
	return file_Proto_starwars_proto_enumTypes[0].Descriptor()
}

func (TipoComando) Type() protoreflect.EnumType {
	return &file_Proto_starwars_proto_enumTypes[0]
}

func (x TipoComando) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use TipoComando.Descriptor instead.
func (TipoComando) EnumDescriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{0}
}

type SolicitudComando struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cmd *Comando `protobuf:"bytes,1,opt,name=cmd,proto3" json:"cmd,omitempty"`
}

func (x *SolicitudComando) Reset() {
	*x = SolicitudComando{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudComando) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudComando) ProtoMessage() {}

func (x *SolicitudComando) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudComando.ProtoReflect.Descriptor instead.
func (*SolicitudComando) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{0}
}

func (x *SolicitudComando) GetCmd() *Comando {
	if x != nil {
		return x.Cmd
	}
	return nil
}

type RespuestaComandoMosEisley struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DirFulcrum string `protobuf:"bytes,1,opt,name=dir_fulcrum,json=dirFulcrum,proto3" json:"dir_fulcrum,omitempty"`
}

func (x *RespuestaComandoMosEisley) Reset() {
	*x = RespuestaComandoMosEisley{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaComandoMosEisley) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaComandoMosEisley) ProtoMessage() {}

func (x *RespuestaComandoMosEisley) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaComandoMosEisley.ProtoReflect.Descriptor instead.
func (*RespuestaComandoMosEisley) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{1}
}

func (x *RespuestaComandoMosEisley) GetDirFulcrum() string {
	if x != nil {
		return x.DirFulcrum
	}
	return ""
}

type RespuestaComandoFulcrum struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RelojVec *RelojVector `protobuf:"bytes,1,opt,name=reloj_vec,json=relojVec,proto3" json:"reloj_vec,omitempty"`
}

func (x *RespuestaComandoFulcrum) Reset() {
	*x = RespuestaComandoFulcrum{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaComandoFulcrum) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaComandoFulcrum) ProtoMessage() {}

func (x *RespuestaComandoFulcrum) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaComandoFulcrum.ProtoReflect.Descriptor instead.
func (*RespuestaComandoFulcrum) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{2}
}

func (x *RespuestaComandoFulcrum) GetRelojVec() *RelojVector {
	if x != nil {
		return x.RelojVec
	}
	return nil
}

type SolicitudGetNumRebelsInformante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cambio *Cambio `protobuf:"bytes,1,opt,name=cambio,proto3" json:"cambio,omitempty"`
}

func (x *SolicitudGetNumRebelsInformante) Reset() {
	*x = SolicitudGetNumRebelsInformante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudGetNumRebelsInformante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudGetNumRebelsInformante) ProtoMessage() {}

func (x *SolicitudGetNumRebelsInformante) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudGetNumRebelsInformante.ProtoReflect.Descriptor instead.
func (*SolicitudGetNumRebelsInformante) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{3}
}

func (x *SolicitudGetNumRebelsInformante) GetCambio() *Cambio {
	if x != nil {
		return x.Cambio
	}
	return nil
}

type RespuestaGetNumRebelsInformante struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cambio    *Cambio `protobuf:"bytes,1,opt,name=cambio,proto3" json:"cambio,omitempty"`
	NumRebels int64   `protobuf:"varint,2,opt,name=num_rebels,json=numRebels,proto3" json:"num_rebels,omitempty"`
}

func (x *RespuestaGetNumRebelsInformante) Reset() {
	*x = RespuestaGetNumRebelsInformante{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaGetNumRebelsInformante) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaGetNumRebelsInformante) ProtoMessage() {}

func (x *RespuestaGetNumRebelsInformante) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaGetNumRebelsInformante.ProtoReflect.Descriptor instead.
func (*RespuestaGetNumRebelsInformante) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{4}
}

func (x *RespuestaGetNumRebelsInformante) GetCambio() *Cambio {
	if x != nil {
		return x.Cambio
	}
	return nil
}

func (x *RespuestaGetNumRebelsInformante) GetNumRebels() int64 {
	if x != nil {
		return x.NumRebels
	}
	return 0
}

type SolicitudGetNumberRebelds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consulta *Consulta `protobuf:"bytes,1,opt,name=consulta,proto3" json:"consulta,omitempty"`
}

func (x *SolicitudGetNumberRebelds) Reset() {
	*x = SolicitudGetNumberRebelds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SolicitudGetNumberRebelds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SolicitudGetNumberRebelds) ProtoMessage() {}

func (x *SolicitudGetNumberRebelds) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SolicitudGetNumberRebelds.ProtoReflect.Descriptor instead.
func (*SolicitudGetNumberRebelds) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{5}
}

func (x *SolicitudGetNumberRebelds) GetConsulta() *Consulta {
	if x != nil {
		return x.Consulta
	}
	return nil
}

type RespuestaGetNumberRebelds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consulta *Consulta `protobuf:"bytes,1,opt,name=consulta,proto3" json:"consulta,omitempty"`
}

func (x *RespuestaGetNumberRebelds) Reset() {
	*x = RespuestaGetNumberRebelds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RespuestaGetNumberRebelds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RespuestaGetNumberRebelds) ProtoMessage() {}

func (x *RespuestaGetNumberRebelds) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RespuestaGetNumberRebelds.ProtoReflect.Descriptor instead.
func (*RespuestaGetNumberRebelds) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{6}
}

func (x *RespuestaGetNumberRebelds) GetConsulta() *Consulta {
	if x != nil {
		return x.Consulta
	}
	return nil
}

type Comando struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tipo          TipoComando `protobuf:"varint,1,opt,name=tipo,proto3,enum=Proto.TipoComando" json:"tipo,omitempty"`
	Nombre        string      `protobuf:"bytes,2,opt,name=nombre,proto3" json:"nombre,omitempty"`
	Coord         *Ubicacion  `protobuf:"bytes,3,opt,name=coord,proto3" json:"coord,omitempty"`
	NuevoValorStr string      `protobuf:"bytes,4,opt,name=nuevo_valor_str,json=nuevoValorStr,proto3" json:"nuevo_valor_str,omitempty"`
	NuevoValorInt int64       `protobuf:"varint,5,opt,name=nuevo_valor_int,json=nuevoValorInt,proto3" json:"nuevo_valor_int,omitempty"`
}

func (x *Comando) Reset() {
	*x = Comando{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comando) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comando) ProtoMessage() {}

func (x *Comando) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comando.ProtoReflect.Descriptor instead.
func (*Comando) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{7}
}

func (x *Comando) GetTipo() TipoComando {
	if x != nil {
		return x.Tipo
	}
	return TipoComando_AddCity
}

func (x *Comando) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *Comando) GetCoord() *Ubicacion {
	if x != nil {
		return x.Coord
	}
	return nil
}

func (x *Comando) GetNuevoValorStr() string {
	if x != nil {
		return x.NuevoValorStr
	}
	return ""
}

func (x *Comando) GetNuevoValorInt() int64 {
	if x != nil {
		return x.NuevoValorInt
	}
	return 0
}

type RelojVector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Nombre string `protobuf:"bytes,1,opt,name=nombre,proto3" json:"nombre,omitempty"`
	X      int64  `protobuf:"varint,2,opt,name=x,proto3" json:"x,omitempty"`
	Y      int64  `protobuf:"varint,3,opt,name=y,proto3" json:"y,omitempty"`
	Z      int64  `protobuf:"varint,4,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *RelojVector) Reset() {
	*x = RelojVector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RelojVector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RelojVector) ProtoMessage() {}

func (x *RelojVector) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RelojVector.ProtoReflect.Descriptor instead.
func (*RelojVector) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{8}
}

func (x *RelojVector) GetNombre() string {
	if x != nil {
		return x.Nombre
	}
	return ""
}

func (x *RelojVector) GetX() int64 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *RelojVector) GetY() int64 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *RelojVector) GetZ() int64 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Ubicacion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NombrePlaneta string `protobuf:"bytes,1,opt,name=nombre_planeta,json=nombrePlaneta,proto3" json:"nombre_planeta,omitempty"`
	NombreCiudad  string `protobuf:"bytes,2,opt,name=nombre_ciudad,json=nombreCiudad,proto3" json:"nombre_ciudad,omitempty"`
}

func (x *Ubicacion) Reset() {
	*x = Ubicacion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Ubicacion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Ubicacion) ProtoMessage() {}

func (x *Ubicacion) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Ubicacion.ProtoReflect.Descriptor instead.
func (*Ubicacion) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{9}
}

func (x *Ubicacion) GetNombrePlaneta() string {
	if x != nil {
		return x.NombrePlaneta
	}
	return ""
}

func (x *Ubicacion) GetNombreCiudad() string {
	if x != nil {
		return x.NombreCiudad
	}
	return ""
}

type Consulta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchivoName string       `protobuf:"bytes,1,opt,name=archivo_name,json=archivoName,proto3" json:"archivo_name,omitempty"`
	Coord       *Ubicacion   `protobuf:"bytes,2,opt,name=coord,proto3" json:"coord,omitempty"`
	NumRebels   int64        `protobuf:"varint,3,opt,name=num_rebels,json=numRebels,proto3" json:"num_rebels,omitempty"`
	RelojVec    *RelojVector `protobuf:"bytes,4,opt,name=reloj_vec,json=relojVec,proto3" json:"reloj_vec,omitempty"`
	FulcrumDir  string       `protobuf:"bytes,5,opt,name=fulcrum_dir,json=fulcrumDir,proto3" json:"fulcrum_dir,omitempty"`
}

func (x *Consulta) Reset() {
	*x = Consulta{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Consulta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Consulta) ProtoMessage() {}

func (x *Consulta) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Consulta.ProtoReflect.Descriptor instead.
func (*Consulta) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{10}
}

func (x *Consulta) GetArchivoName() string {
	if x != nil {
		return x.ArchivoName
	}
	return ""
}

func (x *Consulta) GetCoord() *Ubicacion {
	if x != nil {
		return x.Coord
	}
	return nil
}

func (x *Consulta) GetNumRebels() int64 {
	if x != nil {
		return x.NumRebels
	}
	return 0
}

func (x *Consulta) GetRelojVec() *RelojVector {
	if x != nil {
		return x.RelojVec
	}
	return nil
}

func (x *Consulta) GetFulcrumDir() string {
	if x != nil {
		return x.FulcrumDir
	}
	return ""
}

type Cambio struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ArchivoName string       `protobuf:"bytes,1,opt,name=archivo_name,json=archivoName,proto3" json:"archivo_name,omitempty"`
	Cmd         *Comando     `protobuf:"bytes,2,opt,name=cmd,proto3" json:"cmd,omitempty"`
	RelojVec    *RelojVector `protobuf:"bytes,3,opt,name=reloj_vec,json=relojVec,proto3" json:"reloj_vec,omitempty"`
	FulcrumDir  string       `protobuf:"bytes,4,opt,name=fulcrum_dir,json=fulcrumDir,proto3" json:"fulcrum_dir,omitempty"`
}

func (x *Cambio) Reset() {
	*x = Cambio{}
	if protoimpl.UnsafeEnabled {
		mi := &file_Proto_starwars_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cambio) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cambio) ProtoMessage() {}

func (x *Cambio) ProtoReflect() protoreflect.Message {
	mi := &file_Proto_starwars_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cambio.ProtoReflect.Descriptor instead.
func (*Cambio) Descriptor() ([]byte, []int) {
	return file_Proto_starwars_proto_rawDescGZIP(), []int{11}
}

func (x *Cambio) GetArchivoName() string {
	if x != nil {
		return x.ArchivoName
	}
	return ""
}

func (x *Cambio) GetCmd() *Comando {
	if x != nil {
		return x.Cmd
	}
	return nil
}

func (x *Cambio) GetRelojVec() *RelojVector {
	if x != nil {
		return x.RelojVec
	}
	return nil
}

func (x *Cambio) GetFulcrumDir() string {
	if x != nil {
		return x.FulcrumDir
	}
	return ""
}

var File_Proto_starwars_proto protoreflect.FileDescriptor

var file_Proto_starwars_proto_rawDesc = []byte{
	0x0a, 0x14, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x61, 0x72, 0x77, 0x61, 0x72, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x34, 0x0a,
	0x10, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64,
	0x6f, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x6d, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x52, 0x03,
	0x63, 0x6d, 0x64, 0x22, 0x3c, 0x0a, 0x19, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61,
	0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x4d, 0x6f, 0x73, 0x45, 0x69, 0x73, 0x6c, 0x65, 0x79,
	0x12, 0x1f, 0x0a, 0x0b, 0x64, 0x69, 0x72, 0x5f, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x69, 0x72, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x22, 0x4a, 0x0a, 0x17, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x2f, 0x0a, 0x09,
	0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x12, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x22, 0x48, 0x0a,
	0x1f, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d,
	0x52, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65,
	0x12, 0x25, 0x0a, 0x06, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x52,
	0x06, 0x63, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x22, 0x67, 0x0a, 0x1f, 0x52, 0x65, 0x73, 0x70, 0x75,
	0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73,
	0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x06, 0x63, 0x61,
	0x6d, 0x62, 0x69, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x43, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x52, 0x06, 0x63, 0x61, 0x6d, 0x62, 0x69,
	0x6f, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75, 0x6d, 0x5f, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6e, 0x75, 0x6d, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73,
	0x22, 0x48, 0x0a, 0x19, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x47, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2b, 0x0a,
	0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61,
	0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x22, 0x48, 0x0a, 0x19, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x2b, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x73, 0x75,
	0x6c, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x73,
	0x75, 0x6c, 0x74, 0x61, 0x22, 0xc1, 0x01, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f,
	0x12, 0x26, 0x0a, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x12,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x6f, 0x6d, 0x61, 0x6e,
	0x64, 0x6f, 0x52, 0x04, 0x74, 0x69, 0x70, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62,
	0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63, 0x69, 0x6f,
	0x6e, 0x52, 0x05, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x65, 0x76,
	0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x56, 0x61, 0x6c, 0x6f, 0x72, 0x53, 0x74, 0x72,
	0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x75, 0x65, 0x76, 0x6f, 0x5f, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x5f,
	0x69, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x6e, 0x75, 0x65, 0x76, 0x6f,
	0x56, 0x61, 0x6c, 0x6f, 0x72, 0x49, 0x6e, 0x74, 0x22, 0x4f, 0x0a, 0x0b, 0x52, 0x65, 0x6c, 0x6f,
	0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x01, 0x7a, 0x22, 0x57, 0x0a, 0x09, 0x55, 0x62, 0x69,
	0x63, 0x61, 0x63, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x0a, 0x0e, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65,
	0x5f, 0x70, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d,
	0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x50, 0x6c, 0x61, 0x6e, 0x65, 0x74, 0x61, 0x12, 0x23, 0x0a,
	0x0d, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x5f, 0x63, 0x69, 0x75, 0x64, 0x61, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6e, 0x6f, 0x6d, 0x62, 0x72, 0x65, 0x43, 0x69, 0x75, 0x64,
	0x61, 0x64, 0x22, 0xc6, 0x01, 0x0a, 0x08, 0x43, 0x6f, 0x6e, 0x73, 0x75, 0x6c, 0x74, 0x61, 0x12,
	0x21, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76, 0x6f, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x26, 0x0a, 0x05, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x62, 0x69, 0x63, 0x61, 0x63,
	0x69, 0x6f, 0x6e, 0x52, 0x05, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x75,
	0x6d, 0x5f, 0x72, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6e, 0x75, 0x6d, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x2f, 0x0a, 0x09, 0x72, 0x65, 0x6c,
	0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x52, 0x08, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b, 0x66, 0x75,
	0x6c, 0x63, 0x72, 0x75, 0x6d, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x44, 0x69, 0x72, 0x22, 0x9f, 0x01, 0x0a, 0x06,
	0x43, 0x61, 0x6d, 0x62, 0x69, 0x6f, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x72, 0x63, 0x68, 0x69, 0x76,
	0x6f, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x72,
	0x63, 0x68, 0x69, 0x76, 0x6f, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x03, 0x63, 0x6d, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x52, 0x03, 0x63, 0x6d, 0x64, 0x12, 0x2f, 0x0a, 0x09, 0x72,
	0x65, 0x6c, 0x6f, 0x6a, 0x5f, 0x76, 0x65, 0x63, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12,
	0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x6f, 0x6a, 0x56, 0x65, 0x63, 0x12, 0x1f, 0x0a, 0x0b,
	0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x5f, 0x64, 0x69, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0a, 0x66, 0x75, 0x6c, 0x63, 0x72, 0x75, 0x6d, 0x44, 0x69, 0x72, 0x2a, 0x4c, 0x0a,
	0x0b, 0x54, 0x69, 0x70, 0x6f, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x0b, 0x0a, 0x07,
	0x41, 0x64, 0x64, 0x43, 0x69, 0x74, 0x79, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x10, 0x01, 0x12, 0x10, 0x0a, 0x0c, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x10, 0x02, 0x12, 0x0e, 0x0a, 0x0a, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x69, 0x74, 0x79, 0x10, 0x03, 0x32, 0x97, 0x02, 0x0a, 0x09,
	0x4d, 0x6f, 0x73, 0x45, 0x69, 0x73, 0x6c, 0x65, 0x79, 0x12, 0x44, 0x0a, 0x07, 0x43, 0x6f, 0x6d,
	0x61, 0x6e, 0x64, 0x6f, 0x12, 0x17, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c,
	0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x20, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43,
	0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x4d, 0x6f, 0x73, 0x45, 0x69, 0x73, 0x6c, 0x65, 0x79, 0x12,
	0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65,
	0x6c, 0x64, 0x73, 0x12, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x69,
	0x63, 0x69, 0x74, 0x75, 0x64, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65,
	0x62, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x6c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x74, 0x65, 0x12, 0x26, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f,
	0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x62,
	0x65, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x74, 0x65, 0x1a, 0x26, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47,
	0x65, 0x74, 0x4e, 0x75, 0x6d, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x73, 0x49, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x74, 0x65, 0x32, 0xa5, 0x01, 0x0a, 0x07, 0x46, 0x75, 0x6c, 0x63, 0x72, 0x75,
	0x6d, 0x12, 0x42, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x12, 0x17, 0x2e, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x43, 0x6f,
	0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x1a, 0x1e, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65,
	0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x43, 0x6f, 0x6d, 0x61, 0x6e, 0x64, 0x6f, 0x46, 0x75,
	0x6c, 0x63, 0x72, 0x75, 0x6d, 0x12, 0x56, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x4e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x12, 0x20, 0x2e, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x53, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x74, 0x75, 0x64, 0x47, 0x65, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x1a, 0x20, 0x2e, 0x50, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x75, 0x65, 0x73, 0x74, 0x61, 0x47, 0x65, 0x74,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x65, 0x62, 0x65, 0x6c, 0x64, 0x73, 0x42, 0x26, 0x5a,
	0x24, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x72, 0x6f, 0x6a,
	0x61, 0x73, 0x31, 0x34, 0x2f, 0x4c, 0x61, 0x62, 0x33, 0x49, 0x4e, 0x46, 0x33, 0x34, 0x33, 0x2f,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_Proto_starwars_proto_rawDescOnce sync.Once
	file_Proto_starwars_proto_rawDescData = file_Proto_starwars_proto_rawDesc
)

func file_Proto_starwars_proto_rawDescGZIP() []byte {
	file_Proto_starwars_proto_rawDescOnce.Do(func() {
		file_Proto_starwars_proto_rawDescData = protoimpl.X.CompressGZIP(file_Proto_starwars_proto_rawDescData)
	})
	return file_Proto_starwars_proto_rawDescData
}

var file_Proto_starwars_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_Proto_starwars_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_Proto_starwars_proto_goTypes = []interface{}{
	(TipoComando)(0),                        // 0: Proto.TipoComando
	(*SolicitudComando)(nil),                // 1: Proto.SolicitudComando
	(*RespuestaComandoMosEisley)(nil),       // 2: Proto.RespuestaComandoMosEisley
	(*RespuestaComandoFulcrum)(nil),         // 3: Proto.RespuestaComandoFulcrum
	(*SolicitudGetNumRebelsInformante)(nil), // 4: Proto.SolicitudGetNumRebelsInformante
	(*RespuestaGetNumRebelsInformante)(nil), // 5: Proto.RespuestaGetNumRebelsInformante
	(*SolicitudGetNumberRebelds)(nil),       // 6: Proto.SolicitudGetNumberRebelds
	(*RespuestaGetNumberRebelds)(nil),       // 7: Proto.RespuestaGetNumberRebelds
	(*Comando)(nil),                         // 8: Proto.Comando
	(*RelojVector)(nil),                     // 9: Proto.RelojVector
	(*Ubicacion)(nil),                       // 10: Proto.Ubicacion
	(*Consulta)(nil),                        // 11: Proto.Consulta
	(*Cambio)(nil),                          // 12: Proto.Cambio
}
var file_Proto_starwars_proto_depIdxs = []int32{
	8,  // 0: Proto.SolicitudComando.cmd:type_name -> Proto.Comando
	9,  // 1: Proto.RespuestaComandoFulcrum.reloj_vec:type_name -> Proto.RelojVector
	12, // 2: Proto.SolicitudGetNumRebelsInformante.cambio:type_name -> Proto.Cambio
	12, // 3: Proto.RespuestaGetNumRebelsInformante.cambio:type_name -> Proto.Cambio
	11, // 4: Proto.SolicitudGetNumberRebelds.consulta:type_name -> Proto.Consulta
	11, // 5: Proto.RespuestaGetNumberRebelds.consulta:type_name -> Proto.Consulta
	0,  // 6: Proto.Comando.tipo:type_name -> Proto.TipoComando
	10, // 7: Proto.Comando.coord:type_name -> Proto.Ubicacion
	10, // 8: Proto.Consulta.coord:type_name -> Proto.Ubicacion
	9,  // 9: Proto.Consulta.reloj_vec:type_name -> Proto.RelojVector
	8,  // 10: Proto.Cambio.cmd:type_name -> Proto.Comando
	9,  // 11: Proto.Cambio.reloj_vec:type_name -> Proto.RelojVector
	1,  // 12: Proto.MosEisley.Comando:input_type -> Proto.SolicitudComando
	6,  // 13: Proto.MosEisley.GetNumberRebelds:input_type -> Proto.SolicitudGetNumberRebelds
	4,  // 14: Proto.MosEisley.GetNumberRebeldsInformante:input_type -> Proto.SolicitudGetNumRebelsInformante
	1,  // 15: Proto.Fulcrum.Comando:input_type -> Proto.SolicitudComando
	6,  // 16: Proto.Fulcrum.GetNumberRebelds:input_type -> Proto.SolicitudGetNumberRebelds
	2,  // 17: Proto.MosEisley.Comando:output_type -> Proto.RespuestaComandoMosEisley
	7,  // 18: Proto.MosEisley.GetNumberRebelds:output_type -> Proto.RespuestaGetNumberRebelds
	5,  // 19: Proto.MosEisley.GetNumberRebeldsInformante:output_type -> Proto.RespuestaGetNumRebelsInformante
	3,  // 20: Proto.Fulcrum.Comando:output_type -> Proto.RespuestaComandoFulcrum
	7,  // 21: Proto.Fulcrum.GetNumberRebelds:output_type -> Proto.RespuestaGetNumberRebelds
	17, // [17:22] is the sub-list for method output_type
	12, // [12:17] is the sub-list for method input_type
	12, // [12:12] is the sub-list for extension type_name
	12, // [12:12] is the sub-list for extension extendee
	0,  // [0:12] is the sub-list for field type_name
}

func init() { file_Proto_starwars_proto_init() }
func file_Proto_starwars_proto_init() {
	if File_Proto_starwars_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_Proto_starwars_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudComando); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaComandoMosEisley); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaComandoFulcrum); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudGetNumRebelsInformante); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaGetNumRebelsInformante); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SolicitudGetNumberRebelds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RespuestaGetNumberRebelds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comando); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RelojVector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Ubicacion); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Consulta); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_Proto_starwars_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cambio); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_Proto_starwars_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_Proto_starwars_proto_goTypes,
		DependencyIndexes: file_Proto_starwars_proto_depIdxs,
		EnumInfos:         file_Proto_starwars_proto_enumTypes,
		MessageInfos:      file_Proto_starwars_proto_msgTypes,
	}.Build()
	File_Proto_starwars_proto = out.File
	file_Proto_starwars_proto_rawDesc = nil
	file_Proto_starwars_proto_goTypes = nil
	file_Proto_starwars_proto_depIdxs = nil
}
