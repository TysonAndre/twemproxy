MAINTAINERCLEANFILES = Makefile.in

noinst_LIBRARIES = libreplication.a
noinst_HEADERS = libreplication.h

# Builds a golang library.
# TODO: port to automake
libreplication_a_SOURCES = libreplication.go

bin_PROGRAMS = replication_test

# replication_test is a stub that the library can be built and linked.
# It is not a part of twemproxy
replication_test$(EXEEXT): libreplication.a replication_test.c
	gcc -pthread replication_test.c replication.a -o replication_test

# replication.a (shared library) and replication.h (header) will be used to compile twemproxy with replication support.
libreplication.a: libreplication.go
	go build -buildmode=c-archive libreplication.go
