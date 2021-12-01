package main

const (
	port    = ":50054"
	local   = "localhost" + port
	address = "dist150.inf.santiago.usm.cl" + port
)

const (
	nameNodeFile = "registro.txt"
)

const (
	Port    = ":50055"
	f1Addrs = "dist149.inf.santiago.usm.cl" + Port
	f2Addrs = "dist151.inf.santiago.usm.cl" + Port
	f3Addrs = "dist152.inf.santiago.usm.cl" + Port
)

const (
	dnLocal = "localhost"
	local1  = dnLocal + Port
	local2  = dnLocal + Port
	local3  = dnLocal + Port
)

var (
	RemoteAddrs [3]string = [3]string{f1Addrs, f2Addrs, f3Addrs}
	localAddrs  [3]string = [3]string{local1, local2, local3}
	curAddrs    [3]string
)

var FulcrumAddresses [3]string = [3]string{f1Addrs, f2Addrs, f3Addrs}