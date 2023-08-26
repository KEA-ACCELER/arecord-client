package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"
)

func getAccelerDir() string {
	cwd, err := os.Getwd()
	if err != nil {
		// Handle the error.
		log.Fatal(err)
	}
	// Print the current working directory.
	path := filepath.Join(cwd, "acceler")
	fmt.Println(path)
	return path

}

// LocalAbsToRoot converts an absolute path to a relative path to the root
func LocalAbsToRoot(abs string, root string) string {

	rootdir, _ := path.Split(root)
	result := abs[len(rootdir):]
	return "/" + result
}

// LocalRelToRoot converts a relative path to a relative path to the root
func LocalRelToRoot(rel string, root string) string {

	abs, err := filepath.Abs(rel)
	if err != nil {
		return ""

	}
	return LocalAbsToRoot(abs, root)

}
