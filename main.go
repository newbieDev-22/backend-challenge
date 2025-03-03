package main

import (
	"fmt"

	maxpathlogic "github.com/newbieDev-22/backend-challenge/max-path-logic"
)

func main() {
	array2D := [][]int{
		{59},
		{73, 41},
		{52, 40, 53},
		{26, 53, 6, 34},
	}

	maxPath := maxpathlogic.FindMaxPath(array2D)
	fmt.Println(maxPath)
}
