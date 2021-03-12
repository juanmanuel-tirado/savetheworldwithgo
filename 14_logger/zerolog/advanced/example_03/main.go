package main

import (
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
    "time"
)

func main() {
    logger := log.Sample(&zerolog.BurstSampler{
        Burst: 2,
        Period: time.Second * 5,
        NextSampler: &zerolog.BasicSampler{N: 90000000},
    })

    for i:=0;i<99999999;i++{
        logger.Info().Int("i",i).Msg("")
    }
}
