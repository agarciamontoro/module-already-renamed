package main

import (
	"fmt"

	"github.com/agarciamontoro/module-already-renamed/pack"
)

func main() {
	fmt.Println("running from module-already-renamed...")
	pack.ExposedFunction()
}
