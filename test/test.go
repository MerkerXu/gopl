package main

import (
	"fmt"
	"math"
	"os"
	"runtime"
	"strings"
	"time"
)

func main() {
	testFloat()
	testRune()
	fmt.Printf("%d\n", byte(65536>>16))
	fmt.Printf("%v\n", time.Now())
	hostname := GetHostname("")
	println("主机名:", hostname)
	fmt.Println("操作系统ARCH:", GetOSArch())
}

func testFloat() {
	var f float32 = 16777216
	fmt.Println(f == f+1)
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z)
	nan := math.NaN()
	fmt.Println(nan == nan, nan < nan, nan > nan)
	fmt.Println(math.IsNaN(z/z), math.IsInf(1/z, 0))
}

func testRune() {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
}

func GetHostname(hostnameOverride string) string {
	hostname := hostnameOverride
	if hostname == "" {
		nodename, err := os.Hostname()
		if err != nil {
			fmt.Errorf("Couldn't determine hostname: %v", err)
		}
		hostname = nodename
	}
	return strings.ToLower(strings.TrimSpace(hostname))
}

func GetOSArch() string {
	osArch := runtime.GOARCH
	return osArch
}
