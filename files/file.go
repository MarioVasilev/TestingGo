package files

import (
	"fmt"
	"os"
)

// Create new file by given name
func CreateFile(name string) {
	file, _ := os.Create(name)
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()
	printSomeStuff("Testing trans package communication")
}
