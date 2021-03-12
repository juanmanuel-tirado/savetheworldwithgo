package main

import (
    "os"
    "github.com/rs/zerolog"
)

func main() {
    mainLogger := zerolog.New(os.Stdout).With().Logger()
    mainLogger.Info().Msg("This is the main logger")

    subLogger := mainLogger.With().Str("component","componentA").Logger()
    subLogger.Info().Msg("This is the sublogger")
}
