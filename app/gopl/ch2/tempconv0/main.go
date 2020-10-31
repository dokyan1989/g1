package tempconv

// Celsius ...
type Celsius float64

// Fahrenheit ...
type Fahrenheit float64

// Some temperatures
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 0
)

// CToF ...
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC ...
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
