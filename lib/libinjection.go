package lib

/*
#cgo CFLAGS: -I./libinjection
#cgo LDFLAGS: -L./libinjection -linjection
#include "libinjection.h"
#include "libinjection_sqli.h"
*/
import "C"
import (
	"bytes"
	"unsafe"
)

func TestSQLi(text string) (bool, string) {
	var out [8]C.char
	pointer := (*C.char)(unsafe.Pointer(&out[0]))
	if found := C.libinjection_sqli(C.CString(text), C.size_t(len(text)), pointer); found == 1 {
		output := C.GoBytes(unsafe.Pointer(&out[0]), 8)
		return true, string(output[:bytes.Index(output, []byte{0})])
	}

	return false, ""
}

func TestXSS(text string) bool {
	if found := C.libinjection_xss(C.CString(text), C.size_t(len(text))); found == 1 {
		return true
	}

	return false
}

func TestSQLiString() string {
	return "asdf asd ; -1' and 1=1 union/* foo */select load_file('/etc/passwd')--"
}

func TestXSSString() string {
	return "<img onload=\"alert(1)\" />"
}
