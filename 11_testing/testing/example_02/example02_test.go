package example_02

import "testing"

type Quote string
type Movie string

const (
    Crush Quote = "To crush your enemies..."
    T1000 Quote = "I'll be back"
    Unknown Quote = "unknown"

    Conan Movie = "conan"
    Terminator2 Movie = "terminator2"
    Predator Movie = "predator"
)

func MovieQuote(movie Movie) Quote {
    switch movie {
    case Conan:
        return Crush
    case Terminator2:
        return T1000
    default:
        return Unknown
    }
}

func TestMovieQuote(t *testing.T) {
    movies := []Movie{Conan, Predator, Terminator2}
    for _, m := range movies {
        if q := MovieQuote(m); q==Unknown{
            t.Error("unknown  quote for movie", m)
        }
    }
}


