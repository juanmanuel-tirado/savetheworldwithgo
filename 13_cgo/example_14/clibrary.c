#include "clibrary.h"

#include <stdio.h>

void a_callback_func(callback_f f) {
    printf("C a_callback_func calls a Go gateway function\n");
    printf("       with argument 42\n");
    f(42);
}
