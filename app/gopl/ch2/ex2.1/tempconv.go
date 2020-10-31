package tempconv

import (
	"fmt"
)

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

// Kelvin ...
type Kelvin float64

// Some temperatures
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 0
)

func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }
func (k Kelvin) String() string     { return fmt.Sprintf("%g°K", k) }
