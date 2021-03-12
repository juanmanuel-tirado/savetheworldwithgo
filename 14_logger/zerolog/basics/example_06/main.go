package main

import (
    "github.com/pkg/errors"
    "github.com/rs/zerolog/log"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/pkgerrors"
)

func failA() error {
    return failB()
}

func failB() error {
    return failC()
}

func failC() error {
    return errors.New("C failed")
}

func main() {
    zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack

    err := failA()
    log.Error().Stack().Err(err).Msg("")
}
