
package main

import "fmt"

func ExposedFunction() {
	fmt.Println("running from module-to-be-renamed...")
}

func main() {
	fmt.Println("running...")
}
