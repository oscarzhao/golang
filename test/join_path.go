package main

import (
	"fmt"
	"path"
)

func main() {
	dir := "/etc/tenx/"
	filePath := "/hello/a.txt"
	
	fmt.Printf("join %s and %s to %s\n", dir, filePath, path.Join(dir, filePath))
	fmt.Printf("joined path: %s\n", path.Join(dir, path.Base(filePath)))
}