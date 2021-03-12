package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "math/rand"
)

type ComponentHook struct {
    component string
}

func (h ComponentHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
    if level == zerolog.DebugLevel {
        e.Str("component", h.component)
    }
}

type RandomHook struct{}

func (r RandomHook) Run(e *zerolog.Event, level zerolog.Level, msg string) {
    e.Int("random",rand.Int())
}

func main() {
    logger := log.Hook(ComponentHook{"moduleA"})
    logger = logger.Hook(RandomHook{})
    logger.Info().Msg("Info message")
    logger.Debug().Msg("Debug message")
}
