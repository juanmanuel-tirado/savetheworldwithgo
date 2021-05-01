package main

/*
 #include <stdlib.h>
 typedef struct {
	char* name;
	int id;
 } User;
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("John")
	defer C.free(unsafe.Pointer(name))
	x := C.User{name: name, id: C.int(1234)}
	fmt.Println(C.GoString(x.name), x.id)
	// Should we release x?
	//C.free(unsafe.Pointer(&x))
}
