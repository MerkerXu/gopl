package main

import (
    "fmt"
)

type Flags uint

const (
    FlagUp Flags = 1 << iota    // is up
    FlagBroadcast               // supports broadcast access capability
    FlagLoopback                // is a loopback interface
    FlagPointToPoint            // belongs to a point-to-point link
    FlagMulticast               // supports multicast access capability
)

func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func TurnDown(v *Flags) { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagMulticast }
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

const (
    _ = 1 << (10*iota)
    KiB
    MiB
    GiB
    TiB
    PiB
    EiB
    ZiB
    YiB
)

func main() {
    var v Flags = FlagMulticast | FlagUp
    fmt.Printf("%b %t\n", v, IsUp(v))
    TurnDown(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    SetBroadcast(&v)
    fmt.Printf("%b %t\n", v, IsUp(v))
    fmt.Printf("%b %t\n", v, IsCast(v))

    fmt.Printf("%#X\n", GiB)
    fmt.Println(YiB/ZiB)
}
