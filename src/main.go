package main

import (
	"fmt"
	"math/big"
	"os"
	"os/exec"
	"runtime"
)

type ChainIdType interface {
	int64 | uint64 | *big.Int
}

var (
	LogFie            = ""
	ConfigFile        = ""
	SupportedNetworks = []int{42161, 1, 10, 100, 136}
)

// This function clears the terminal screen.
func ClearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
func main() {

	SplashScreen()
}

func Init() {

}
func SplashScreen() {
	ClearScreen()
	fmt.Println("Copper Protocol => Chain Connector")

	fmt.Printf("\n\n")
}
