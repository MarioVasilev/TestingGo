package file

import (
	"fmt"
	"os"
)

func file(name string) {
	file, _ := os.Create(name)
	fmt.Fprint(file, "This is how you write to a file, by the way")
	file.Close()
}
