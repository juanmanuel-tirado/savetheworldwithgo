package main

import (
    "encoding/json"
    "github.com/rs/zerolog/log"
)

type AStruct struct {
    FieldA string
    FieldB int
    fieldC bool
}

type AJSONStruct struct {
    FieldA string   `json:"fieldA,omitempty"`
    FieldB int      `json:"fieldB,omitempty"`
    fieldC bool
}

func main() {
    a := AStruct{"a string", 42, false}
    b := AJSONStruct{"a string", 42, false}

    log.Info().Interface("a",a).Msg("AStruct")
    log.Info().Interface("b",b).Msg("AJSONStruct")

    encoded, _ := json.Marshal(b)
    log.Info().RawJSON("encoded",encoded).Msg("Encoded JSON")
}
