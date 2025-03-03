package main

import (
	"fmt"

	leftrightencode "github.com/newbieDev-22/backend-challenge/left-right-decode"
)

func main() {
	numString := "=LLRR"
	decoded := leftrightencode.LeftRightDecode(numString)
	fmt.Println(decoded)
}
