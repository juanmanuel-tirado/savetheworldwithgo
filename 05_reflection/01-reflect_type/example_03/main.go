package main

import (
    "fmt"
    "reflect"
)

type Adder interface{
    Add (int, int) int
}

type Calculator struct{}

func(c *Calculator) Add(a int, b int) int {
    return a + b
}

func main() {

   var ptrAdder *Adder
   adderType := reflect.TypeOf(ptrAdder).Elem()

   c := Calculator{}

   calcType := reflect.TypeOf(c)
   calcTypePtr := reflect.TypeOf(&c)

   fmt.Println(calcType, calcType.Implements(adderType))
   fmt.Println(calcTypePtr, calcTypePtr.Implements(adderType))
}
