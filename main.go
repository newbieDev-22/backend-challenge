package main

import (
	"fmt"

	leftrightencode "github.com/newbieDev-22/backend-challenge/left-right-encode"
)

func main() {
	numString := "=LLRR"
	decoded := leftrightencode.LeftRightDecode(numString)
	fmt.Println(decoded)
}
