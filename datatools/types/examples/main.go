package main

import (
	"fmt"

	"github.com/skeptycal/util/datatools/math/fastinvsqrt"
	"github.com/skeptycal/util/datatools/types/justforfunc18"
)

func main() {
	if runTempTests == true {
		justforfunc18.TempTests()
	}
	var b fastinvsqrt.Bits = make([]byte, 0, 4)
	fmt.Println(b)
}

const runTempTests = false

// type celsius float64

// func (c celsius) String() string    { return fmt.Sprintf("%.2f °C", c) }
// func (c celsius) Value() float64    { return float64(c) }
// func (c celsius) Farenheit() string { return fmt.Sprintf("%.2f °F", c.Value()*cf+32) }
// func NewCelsius(f float64) celsius  { return celsius((f - 32.0) * fc) }

// type temperature struct{ celsius }

// type anotherTemp temperature
