package conversion

import "fmt"

type Celsius float64
type Fahrenheit float64

func (c Celsius) String() string { return fmt.Sprintf("%gºC", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%gºF", f) }

type Feet float64
type Meter float64

func (f Feet) String() string { return fmt.Sprintf("%gft", f) }
func (m Meter) String() string { return fmt.Sprintf("%gm", m) }

type Kilogram float64
type Pound  float64

func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
func (p Pound) String() string { return fmt.Sprintf("%glb", p) }
