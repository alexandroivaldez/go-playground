// A package is required for every go project.
// main is a special package name, that tells go that this is the main entry point.
// Acts like a basic main in any other language. Sort of like README.md, it's what's ran first.
package main

// fmt is a native package from go
import "fmt"

func main() {
	fmt.Print("Hello World")
}
