package main

import(
    "github.com/rs/zerolog/log"
)

func main() {
    // Panic or fatal messages stop the execution flow
    // log.Panic().Msg("This is a panic message")
    // log.Fatal().Msg("This is a fatal message")
    log.Error().Msg("This is an error message")
    log.Warn().Msg("This is a warning message")
    log.Info().Msg("This is an information message")
    log.Debug().Msg("This is a debug message")
    log.Trace().Msg("This is a trace message")
}
