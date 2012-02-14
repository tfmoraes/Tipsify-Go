LDFLAGS += -L ${GOPATH}
GCFLAGS += -I ${GOPATH}
include ${GOROOT}/src/Make.inc

TARG=tipsify_sorter

GOFILES=\
        tipsify_sorter.go\

include ${GOROOT}/src/Make.cmd
