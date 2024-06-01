package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var path1, path2 string
	fmt.Println("Please Enter the Path Source")
	fmt.Scanln(&path1)
	fmt.Println("Please Enter the Path Destination")
	fmt.Scanln(&path2)
	entries, err := os.ReadDir(path1)
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			fmt.Println(e.Name())
			Path := path2 + "\\" + e.Name()
			err := os.MkdirAll(Path, os.ModeDir)
			if err != nil || os.IsExist(err) {
				log.Fatal(err)
			}
		}
	}
}
