package main

/*
#include<stdio.h>
void f_char(char in) {
	printf("char %c\n", in);
}
void f_schar(signed char in) {
	printf("signed char %X\n", in);
}
void f_uchar(unsigned char in) {
	printf("unsigned char %c\n", in);
}
void f_short(short in) {
	printf("short %d\n", in);
}
void f_ushort(unsigned short in) {
	printf("unsigned short %u\n", in);
}
void f_int(int in) {
	printf("int %d\n", in);
}
void f_uint(unsigned int in) {
	printf("unsigned int %u\n", in);
}
void f_long(long in) {
	printf("long %X\n", in);
}
void f_ulong(unsigned long in) {
	printf("unsigned long %X\n", in);
}
void f_longlong(long long in) {
	printf("long long %X\n", in);
}
void f_ulonglong(unsigned long long in) {
	printf("unsigned long long %X\n", in);
}
void f_float(float in) {
	printf("float %e\n", in);
}
void f_double(double in) {
	printf("float %e\n", in);
}
*/
import "C"
import "math"

func main() {
	C.f_char(C.char(42))
	C.f_schar(C.schar(-42))
	C.f_uchar(C.uchar(42))

	C.f_short(C.short(-42))
	C.f_ushort(C.ushort(42))

	C.f_int(C.int(-42))
	C.f_uint(C.uint(math.MaxInt32))

	C.f_long(C.long(-42))
	C.f_ulong(C.ulong(math.MaxUint64))

	C.f_longlong(C.longlong(math.MaxInt64))
	C.f_ulonglong(C.ulonglong(math.MaxUint64))

	C.f_float(C.float(math.MaxFloat32))
	C.f_double(C.double(math.MaxFloat64))
}
