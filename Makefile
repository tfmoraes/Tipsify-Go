LDFLAGS += -L ${GOPATH}
GCFLAGS += -I ${GOPATH}
include ${GOROOT}/src/Make.inc

TARG=ply_reader

GOFILES=\
        ply_reader.go\

include ${GOROOT}/src/Make.cmd
