package main

import (
    "fmt"
    "os"
    "strconv"
    "gopl/ch2/conversion"
)

func main() {
    for _, arg := range os.Args[1:] {
        t, err := strconv.ParseFloat(arg, 64)
        if err != nil {
            fmt.Fprintf(os.Stderr, "conv: %v\n", err)
            os.Exit(1)
        }

        f := conversion.Fahrenheit(t)
        c := conversion.Celsius(t)
        fmt.Printf("%s = %s, %s = %s\n",
            f, conversion.FToC(f), c, conversion.CToF(c))

        ft := conversion.Feet(t)
        m  := conversion.Meter(t)
        fmt.Printf("%s = %s, %s = %s\n",
            ft, conversion.FToM(ft), m, conversion.MToF(m))

        kg := conversion.Kilogram(t)
        lb := conversion.Pound(t)
        fmt.Printf("%s = %s, %s = %s\n",
            kg, conversion.KToP(kg), lb, conversion.PToK(lb))
    }
}
