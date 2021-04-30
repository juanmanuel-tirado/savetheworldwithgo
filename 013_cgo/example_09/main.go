package main

/*
 #include <stdlib.h>
 typedef struct {
	char* name;
	int id;
 } User;

 User* create_user(char* name) {
	 User* u;
	 u = (User*)malloc(sizeof(User*));
	 u -> name = name;
	 u -> id = 1234;
	 return u;
 }
*/
import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	name := C.CString("John")
	u := C.create_user(name)
	defer C.free(unsafe.Pointer(u))
	fmt.Printf("User name: %s\n", C.GoString(u.name))
	fmt.Printf("User id: %d\n", u.id)
}
