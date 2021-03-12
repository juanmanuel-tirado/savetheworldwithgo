package main

import (
    "bufio"
    "bytes"
    "encoding/json"
    "fmt"
    "net/http"
    "strings"
)

const (
	URL = "http://localhost:8082/topics/%s/partitions/%d"
	CONTENT_TYPE = "application/vnd.kafka.json.v2+json"
)

type User struct{
    Name string `json:"name"`
    Email string `json:"email"`
}

func BuildBody (users []User) string {
    values := make([]string,len(users))
    for i, u := range users {
        encoded, err := json.Marshal(&u)
        if err != nil {
            panic(err)
        }
        values[i] = fmt.Sprintf("{\"value\":%s}", encoded)
    }
    result := strings.Join(values,",")
    return fmt.Sprintf("{\"records\": [%s]}", result)
}

func main() {
    users := []User{{"John", "john@gmail.com"},{"Mary","mary@email.com"}}
    body := BuildBody(users)
    fmt.Println(body)
    bufferBody := bytes.NewBuffer([]byte(body))

    resp, err := http.Post(fmt.Sprintf(URL,"helloTopic",0),CONTENT_TYPE, bufferBody)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println(resp.Status)
    bodyAnswer := bufio.NewScanner(resp.Body)
    for bodyAnswer.Scan() {
        fmt.Println(bodyAnswer.Text())
    }
}
