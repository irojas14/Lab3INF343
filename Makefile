ifdef OS
	RM = del /Q
	FixPath = $(subst /,\,$1)
else
	ifeq ($(shell uname), Linux)
		RM = rm -f
		FixPath = $1
	endif
endif

leia: 
	go run LeiaOrgana/leiaorgana.go
moseisley: 
	go run MosEisley/moseisley.go
informante: 
	go run Informante/informante.go
fulcrum1:
	go run Fulcrum/fulcrum.go 1
fulcrum2:
	go run Fulcrum/fulcrum.go 2
fulcrum3:
	go run Fulcrum/fulcrum.go 3
clean:
	find . -type f -name '*.txt' ! -name 'dummy.txt' -delete