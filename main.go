package main

import (
	"fmt"

	"github.com/mskutle/hello-go/utils"
)

func main() {
	name := "Bobby"
	fmt.Println(utils.Reverse(name))
}
