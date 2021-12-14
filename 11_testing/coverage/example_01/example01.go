package example_01

func Periods(year int) string {
    switch {
    case year < -3000:
        return "Copper Age"
    case year < -2000:
        return "Bronze Age"
    case year < -1000:
        return "Iron Age"
    case year < 0:
        return "Classic Age"
    case year < 476:
        return "Roman Age"
    case year < 1492:
        return "Middle Age"
    case year < 1800:
        return "Modern Age"
    default:
        return "unknown"
    }
}
