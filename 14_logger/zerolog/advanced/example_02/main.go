package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {
    logger := log.Sample(&zerolog.BasicSampler{N:200})
    for i:=0;i<1000;i++{
        logger.Info().Int("i",i).Msg("")
    }
}

