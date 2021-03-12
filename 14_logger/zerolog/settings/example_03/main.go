package main

import (
    "io/ioutil"
    "os"
    "fmt"
    "github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func main() {

    tempFile, err := ioutil.TempFile(os.TempDir(),"deleteme")
    if err != nil {
        log.Error().Err(err).
            Msg("there was an error creating a temporary file four our log")
    }
    defer tempFile.Close()
    fmt.Printf("The log file is allocated at %s\n", tempFile.Name())

    fileWriter := zerolog.New(tempFile).With().Logger()
    consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}

    multi := zerolog.MultiLevelWriter(consoleWriter, os.Stdout, fileWriter)

    logger := zerolog.New(multi).With().Timestamp().Logger()

    logger.Info().Msg("Save the world with Go!!!")
}
