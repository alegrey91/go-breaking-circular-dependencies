package v2

type Prediction uint8

const (
    Sunny Prediction = iota
    Rain
    Overcast
    Snow
    Unknown
)

// This interface is prepared to comply the methods of package "location"
type Locator interface {
    Coords() (float64, float64, error)
}

// Differently from before, "prediction" functions have been reduced to one.
func Predict(loc Locator) (Prediction, error) {
    _, _, err := loc.Coords()
    if err != nil {
        return Unknown, err
    }

    return Overcast, nil
}
