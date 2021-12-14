package main

/*
 #include <stdio.h>
 typedef struct {
	 int id;
	 int age;
 } User;

 void print_user(User *u) {
	 printf("Id: %d\n", u->id);
	 printf("Age: %d\n", u->age);
 }
*/
import "C"
import "unsafe"

type User struct {
	Id  int32
	Age int32
}

func main() {
	u := User{1234, 33}
	userToC := (*C.User)(unsafe.Pointer(&u))
	C.print_user(userToC)
}
