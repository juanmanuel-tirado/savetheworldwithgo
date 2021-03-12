// Package example_01 contains a documentation example.
package example_01

import "fmt"

// Msg represents a message.
type Msg struct{
    // Note is the note to be sent with the message.
    Note string
}

// Send a message to a target destination.
func (m *Msg) Send(target string) {
    fmt.Printf("Send %s to %s\n", m.Note, target)
}

// Receive a message from a certain origin.
func (m *Msg) Receive(origin string) {
    fmt.Printf("Received %s from %s\n", m.Note, origin)
}
