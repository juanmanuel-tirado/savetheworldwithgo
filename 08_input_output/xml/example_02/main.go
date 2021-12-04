package main

import (
    "encoding/xml"
    "errors"
    "fmt"
)

type MyMap map[string]string

func (s MyMap) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
    tokens := []xml.Token{start}

    for key, value := range s {
        t := xml.StartElement{Name: xml.Name{"", key}}
        tokens = append(tokens, t, xml.CharData(value), xml.EndElement{t.Name})
    }

    tokens = append(tokens, xml.EndElement{start.Name})

    for _, t := range tokens {
        err := e.EncodeToken(t)
        if err != nil {
            return err
        }
    }

    return e.Flush()
}

func (a MyMap) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

    key := ""
    val := ""

    for {

        t, _ := d.Token()
        switch tt := t.(type) {

        case xml.StartElement:
            key = tt.Name.Local
        case xml.CharData:
            val = string(tt)
        case xml.EndElement:
            if len(key) != 0{
                a[key] = val
                key,val = "", ""
            }
            if tt.Name == start.Name {
                return nil
            }

        default:
            return errors.New(fmt.Sprintf("uknown %T",t))
        }
    }

   return nil
}


func main() {

    var theMap MyMap = map[string]string{"one": "1","two":"2","three":"3"}
    aMap, _ := xml.MarshalIndent(&theMap, "", "    ")
    fmt.Println(string(aMap))

    var recoveredMap MyMap = make(map[string]string)

    err := xml.Unmarshal(aMap, &recoveredMap)
    if err != nil {
        panic(err)
    }

    fmt.Println(recoveredMap)
}
