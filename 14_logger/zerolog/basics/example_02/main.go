package main

import (
    "github.com/rs/zerolog/log"
    "github.com/rs/zerolog"
)

func main() {
    zerolog.SetGlobalLevel(zerolog.DebugLevel)

    log.Debug().Msg("Debug message is displayed")
    log.Info().Msg("Info Message is displayed")

    zerolog.SetGlobalLevel(zerolog.InfoLevel)
    log.Debug().Msg("Debug message is no longer displayed")
    log.Info().Msg("Info message is displayed")
}
