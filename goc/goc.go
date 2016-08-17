//http://blog.giorgis.io/cgo-examples
package main

/*
#include <stdio.h>
#include <stdint.h>
#include <stdlib.h>
int ic = 5;
unsigned int uic = 7;
int16_t is = 12345;
char* cstring = "C string example";
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {

	var ig int = 10

	igc := int(C.ic)
	fmt.Println("value:", igc, "type:", reflect.TypeOf(igc))

	icg := C.int(ig)
	fmt.Println("value:", icg, "type:", reflect.TypeOf(icg))

	icp := (*C.int)(unsafe.Pointer(&ig))
	fmt.Println("value:", reflect.ValueOf(icp), "type:", reflect.TypeOf(icp))

	uigc := uint(C.uic)
	fmt.Println("value:", uigc, "type:", reflect.TypeOf(uigc))

	i64t := int16(C.is)
	fmt.Println("value:", i64t, "type:", reflect.TypeOf(i64t))


	var gstring string = "Go string example"

	//Go to C String, Output: *C.char
	cs := C.CString(gstring)
	defer C.free(unsafe.Pointer(cs))
	fmt.Println(cs)

	//C to Go String, Output: string
	gs := C.GoString(C.cstring)
	fmt.Println(gs)

	//C string, length to Go string
	gs2 := C.GoStringN(C.cstring, (C.int)(len(gs)))
	fmt.Println(gs2)

}