package main

import "fmt"

func main() {
    var aBool bool = true
    var aString string = "yXXXy"
    var aComplex complex64 = 5i
    var aRune rune = '€'

    fmt.Println(aBool)
    fmt.Println(aString)
    fmt.Println(aComplex)
    fmt.Println(aRune)
    fmt.Printf("%U\n",aRune)
    fmt.Printf("%c\n",aRune)
}
