package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
}

// The name of each package-level entity is visible
// not only throughout the source file that contains its
// declaration, but throughout all the files of the package.
