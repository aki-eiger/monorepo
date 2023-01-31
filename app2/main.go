package main

import (
	"fmt"

	"github.com/aki-eiger/monorepo/module1"
)

func main() {
	fmt.Println("Start!")
	str := module1.CommonFunctionality("Blah 2")
	fmt.Println(str)

}
