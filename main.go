package gocamp

/*
#cgo CFLAGS: -std=gnu99

#include <stdlib.h>
#include "amp.h"
*/
import "C"
import "unsafe"

func Encode(args []string) *C.char {
	var char *C.char

	size := unsafe.Sizeof(char)

	ptr := C.malloc(C.size_t(len(args)) * C.size_t(size))
	defer C.free(ptr)

	for i := 0; i < len(args); i++ {
		element := (**C.char)(unsafe.Pointer(uintptr(ptr) + uintptr(i)*size))
		*element = C.CString(string(args[i]))
	}

	buf := C.amp_encode((**C.char)(ptr), C.int(len(args)))

	return buf
}

func Decode(buf *C.char) []string {
	var msg C.amp_t

	C.amp_decode(&msg, buf)

	var result []string

	for i := 0; C.short(i) < msg.argc; i++ {
		result = append(result, C.GoString(C.amp_decode_arg(&msg)))
	}

	return result
}
