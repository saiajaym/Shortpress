package Shortpress

// #cgo CFLAGS: -Wall -O2
//#include<stdlib.h>
// #include "smaz.c"
import "C"
import "unsafe"

const MAX_BUFFER = 4096

//Compress a given short string and returs a compressed string
//Max len of string should be below 4096 bytes
func Compress(str string) string {
	in := C.CString(str)
	cint := C.int(len(str))
	var output string
	conv := C.CString(output)
	defer C.free(unsafe.Pointer(in))
	defer C.free(unsafe.Pointer(conv))
	C.smaz_compress(in, cint, conv, MAX_BUFFER)
	return C.GoString(conv)
}

//Decompress a compressed string and returns original
//Max len of string should be below 4096 bytes
func Decompress(str string) string {
	in := C.CString(str)
	cint := C.int(len(str))
	var output string
	conv := C.CString(output)
	defer C.free(unsafe.Pointer(in))
	defer C.free(unsafe.Pointer(conv))
	C.smaz_decompress(in, cint, conv, MAX_BUFFER)
	return C.GoString(conv)
}
