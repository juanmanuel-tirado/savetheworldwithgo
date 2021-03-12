package main

import (
    "fmt"
    "reflect"
)


func BuildAdder (i interface{}) {
    fn := reflect.ValueOf(i).Elem()

    newF := reflect.MakeFunc(fn.Type(), func(in []reflect.Value)[]reflect.Value{

        if len(in) > 2 {
            return []reflect.Value{}
        }

        a, b := in[0], in[1]

        if a.Kind() != b.Kind() {
            return []reflect.Value{}
        }

        var result reflect.Value

        switch a.Kind() {
        case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
            result = reflect.ValueOf(a.Int() + b.Int())
        case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
            result = reflect.ValueOf(a.Uint() + b.Uint())
        case reflect.Float32, reflect.Float64:
            result = reflect.ValueOf(a.Float() + b.Float())
        case reflect.String:
            result = reflect.ValueOf(a.String() + b.String())
        default:
            result = reflect.ValueOf(interface{}(nil))
        }
        return []reflect.Value{result}
    })
    fn.Set(newF)
}

func main() {
    var intAdder func(int64,int64) int64
    var floatAdder func(float64, float64) float64
    var strAdder func(string, string) string

    BuildAdder(&intAdder)
    BuildAdder(&floatAdder)
    BuildAdder(&strAdder)

    fmt.Println(intAdder(1,2))
    fmt.Println(floatAdder(3.0,2.423))
    fmt.Println(strAdder("hello"," go"))
}
