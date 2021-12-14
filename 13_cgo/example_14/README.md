# C function pointer callbacks

This example requires compiling and linking additional C code. The following lines should work in a Unix-compatible system.

```bash
gcc -c clibrary.c
ar cr libclibrary.a clibrary.o
go build *.go
./cfuncs
```
