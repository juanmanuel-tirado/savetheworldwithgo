package main

import (
    "bytes"
    "encoding/json"
    "errors"
    "fmt"
    "reflect"
    "strings"
)

func Marshal(input interface{}) ([]byte, error) {
    var buffer bytes.Buffer
    t := reflect.TypeOf(input)
    v := reflect.ValueOf(input)

    buffer.WriteString("{")
    for i:=0;i < t.NumField();i++ {
        encodedField,err := encodeField(t.Field(i),v.Field(i))

        if err != nil {
            return nil, err
        }
        if len(encodedField) != 0 {
            if i >0 && i<= t.NumField()-1 {
                buffer.WriteString(", ")
            }
            buffer.WriteString(encodedField)
        }
    }
    buffer.WriteString("}")
    return buffer.Bytes(),nil
}

func encodeField(f reflect.StructField, v reflect.Value) (string, error) {

    if f.PkgPath!=""{
        return "",nil
    }

    if f.Type.Kind() != reflect.String {
        return "", nil
    }

    tag, found := f.Tag.Lookup("pretty")
    if !found {
        return "", nil
    }

    result := f.Name+":"
    var err error = nil
    switch tag {
    case "upper":
        result = result + strings.ToUpper(v.String())
    case "lower":
        result = result + strings.ToLower(v.String())
    default:
        err = errors.New("invalid tag value")
    }
    if err != nil {
        return "", err
    }

    return result, nil
}

type User struct {
    UserId string `pretty:"upper"`
    Email string `pretty:"lower"`
    password string `pretty:"lower"`
}

type Record struct {
    Name string `pretty:"lower" json:"name"`
    Surname string `pretty:"upper" json:"surname"`
    Age int `pretty:"other" json:"age"`
}


func main() {
    u := User{"John", "John@Gmail.com", "admin"}

    marSer, _ := Marshal(u)
    fmt.Println("pretty user", string(marSer))

    r := Record{"John", "Johnson",33}
    marRec, _:= Marshal(r)
    fmt.Println("pretty rec", string(marRec))

    jsonRec, _ := json.Marshal(r)
    fmt.Println("json rec",string(jsonRec))
}
