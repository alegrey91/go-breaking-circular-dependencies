package location

type GPSCoords [2]float64

type PlusCode string

type Address struct {
    Street      string
    City        string
    PostCode    string
    Country     string
}

// For each type we have its own "Coords" function
func (g GPSCoords) Coords() (float64, float64, error) {
    return 2.0, 3.9, nil
}

func (pc PlusCode) Coords() (float64, float64, error) {
    return 10.0, 9.0, nil
}

func (a Address) Coords() (float64, float64, error) {
    return 11.0, 4.0, nil
}
