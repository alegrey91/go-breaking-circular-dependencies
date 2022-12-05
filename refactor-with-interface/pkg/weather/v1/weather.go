package v1

type Prediction uint8

const (
    Sunny Prediction = iota
    Rain
    Overcast
    Snow
    Unknown
)

func PredictAtCoords(lat, long float64) (Prediction, error) {
    // Totally random return value.
    return Sunny, nil
}

func PredictAtPlusCode(code string) (Prediction, error) {
    // Totally random return value.
    return Snow, nil
}
