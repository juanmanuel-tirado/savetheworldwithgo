package main

import(
    "github.com/rs/zerolog/log"
)

func main() {
    log.Info().Str("mystr","this is a string").Msg("")
    log.Info().Int("myint",1234).Msg("")
    log.Info().Int("myint",1234).Str("str","some string").Msg("And a regular message")
}
