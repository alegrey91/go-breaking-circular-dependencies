package main

import (
    "fmt"

    weather1 "github.com/alegrey91/swe-golang/refactorwithinterface/pkg/weather/v1"
    weather2 "github.com/alegrey91/swe-golang/refactorwithinterface/pkg/weather/v2"
    "github.com/alegrey91/swe-golang/refactorwithinterface/pkg/weather/v2/location"
)

// makePredictionV1 makes use of v1 package
func makePredictionV1() error {
    w, err := weather1.PredictAtCoords(5.0, 6.34)
    if err != nil {
        return err
    }
    fmt.Println("The weather prediction (v1) for London is: ", w)
    return nil
}

// makePredictionV1 makes use of v2 package
func makePredictionV2() error {
    loc := location.PlusCode("9C3XGV00+")
    pred, err := weather2.Predict(loc)
    if err != nil {
        return err
    }
    fmt.Println("The weather prediction (v2) for London is: ", pred)
    return nil
}

func main() {
    
    makePredictionV1()
    makePredictionV2()
}
