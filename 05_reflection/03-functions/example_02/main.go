//package main
//
//import (
//    "fmt"
//    "reflect"
//)
//
//type User struct {
//    UserId string
//    Email string
//}
//
//const Trojan int64 = 42
//
//func NewStructCounter(i interface{}) interface{} {
//    v := reflect.ValueOf(i)
//    t := reflect.TypeOf(i)
//    var stf []reflect.StructField
//
//    counterField := reflect.StructField{
//        Name: "Trojan",
//        Type: reflect.TypeOf(Trojan),
//    }
//
//    for i:=0;i<v.NumField();i++ {
//        newField := reflect.StructField{
//            Name: t.Field(i).Name,
//            Type: t.Field(i).Type,
//        }
//        stf = append(stf, newField)
//    }
//    stf = append(stf, counterField)
//
//    newType := reflect.StructOf(stf)
//    newStruct := reflect.New(newType).Interface()
//
//    // build something new from the interface
//    toReturn := reflect.ValueOf(newStruct)
//    // fill it
//    for i := 0; i<v.NumField(); i++ {
//        toReturn.Elem().Field(i).Set(v.Field(i))
//    }
//    // set trojan
//    toReturn.Elem().FieldByName("Trojan").SetInt(Trojan)
//
//    //return newStruct.Interface()
//    return toReturn
//}
//
//func main() {
//    u := User{"John", "john@gmail.com"}
//    fmt.Println(u)
//
//    uCounter := NewStructCounter(u)
//    fmt.Println(uCounter)
//
//    v := reflect.ValueOf(uCounter)
//    fmt.Println(v)
//
//    v.Field(2).SetInt(666)
//    fmt.Println(uCounter)
//
//    //sr := reflect.ValueOf(uCounter)
//    //fmt.Println(sr.Elem().Field(0).Interface())
//    //sr.Elem().Field(0).Set(reflect.ValueOf("hola"))
//    //fmt.Println(sr)
//
//    //t := reflect.TypeOf(uCounter).Elem()
//    //fmt.Println(t)
//    //newU := reflect.New(t).Elem().Interface()
//    //fmt.Println(newU)
//
//}
