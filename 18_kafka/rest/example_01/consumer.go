package main

import (
    "bufio"
    "bytes"
    "time"
    "fmt"
    "net/http"
)

const (
    HOST           = "http://localhost:8082"
    CONSUMER       = "testConsumer"
    GROUP          = "testGroup"
	NEW_CONSUMER   = "%s/consumers/%s"
    SUBSCRIBE_CONSUMER = "%s/consumers/%s/instances/%s/subscription"
	FETCH_CONSUMER = "%s/consumers/%s/instances/%s/records"
	DELETE_CONSUMER = "%s/consumers/%s/instances/%s"
    CONTENT_TYPE   = "application/vnd.kafka.json.v2+json"
)

func DoHelper(client *http.Client, url string, body []byte ) error {
    bufferBody := bytes.NewBuffer(body)
    req, err := http.NewRequest(http.MethodPost,url, bufferBody)
    if err != nil {
        return err
    }
    fmt.Printf("-->Call %s\n",req.URL)
    fmt.Printf("-->Body %s\n",string(body))
    resp, err := client.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    bodyResp := bufio.NewScanner(resp.Body)
    fmt.Printf("<--Response %s\n", resp.Status)
    for bodyResp.Scan() {
        fmt.Printf("<--Body %s\n",bodyResp.Text())
    }
    return nil
}

func main() {
    client := http.Client{}
    // New consumer
    url := fmt.Sprintf(NEW_CONSUMER,HOST,GROUP)
    body := fmt.Sprintf(`{"name":"%s", "format": "json"}`,CONSUMER)
    err := DoHelper(&client, url, []byte(body))
    if err != nil {
        panic(err)
    }
    time.Sleep(time.Second)
    // Subscribe to topic
    url = fmt.Sprintf(SUBSCRIBE_CONSUMER,HOST,GROUP,CONSUMER)
    body = `{"topics":["helloTopic"]}`
    err = DoHelper(&client, url, []byte(body))
    if err != nil {
        panic(err)
    }
    time.Sleep(time.Second)
    // Get records
    req, err := http.NewRequest(http.MethodGet, fmt.Sprintf(FETCH_CONSUMER,HOST,GROUP,CONSUMER), nil)
    if err != nil {
        panic(err)
    }
    req.Header.Add("Accept",CONTENT_TYPE)
    fmt.Printf("-->Call %s\n",req.URL)
    respRecords, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer respRecords.Body.Close()
    fmt.Printf("<--Response %s\n", respRecords.Status)
    recordsBodyResp := bufio.NewScanner(respRecords.Body)
    for recordsBodyResp.Scan() {
        fmt.Printf("<--Body %s\n",recordsBodyResp.Text())
    }
    // Delete consumer instance
    deleteReq, err := http.NewRequest(http.MethodDelete,fmt.Sprintf(DELETE_CONSUMER,HOST,GROUP,CONSUMER),nil)
    if err != nil {
        panic(err)
    }
    fmt.Printf("-->Call %s\n",deleteReq.URL)
    resp, err := client.Do(deleteReq)
    if err != nil {
        panic(err)
    }
    fmt.Printf("<--Response %s\n",resp.Status)
}
