package main

import (
	"fmt"

	"github.com/agarciamontoro/module-to-be-renamed/pack"
)

func main() {
	fmt.Println("running from module-to-be-renamed...")
	pack.ExposedFunction()
}
