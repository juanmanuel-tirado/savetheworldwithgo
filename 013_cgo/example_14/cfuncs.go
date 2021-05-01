package main

/*
#include <stdio.h>
void goFunction_cgo(int number) {
	printf("go_Function_cgo with number %d\n",number);
	void goFunction(int);
	goFunction(number);
}
*/
import "C"
