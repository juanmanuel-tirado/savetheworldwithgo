package main

import (
    "os"
    "strings"
    "time"
    "fmt"
    "github.com/rs/zerolog"
)

func main() {
    output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
    output.FormatLevel = func(i interface{}) string {
        return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
    }
    output.FormatMessage = func(i interface{}) string {
        return fmt.Sprintf(">>>%s<<<", i)
    }
    output.FormatFieldName = func(i interface{}) string {
        return fmt.Sprintf("[%s]:", i)
    }
    output.FormatFieldValue = func(i interface{}) string {
        return strings.ToUpper(fmt.Sprintf("[%s]", i))
    }

    log := zerolog.New(output).With().Timestamp().Logger()

    log.Info().Str("foo", "bar").Msg("Save the world with Go!!!")
}
