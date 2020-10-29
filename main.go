package main

import (
	"fmt"

	"github.com/agarciamontoro/module-already-renamed/pack"
)

func main() {
	fmt.Println("(module and imports renamed) running from module-already-renamed...")
	pack.ExposedFunction()
}
