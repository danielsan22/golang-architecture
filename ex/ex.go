package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var fileName string = "file-01.txt"
var newFile string = "file-02.txt"

func main() {

	file, err := os.Open(fileName)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	newFile, err := os.Create(newFile)

	if err != nil {
		log.Fatal(err)
	}
	defer newFile.Close()

	bytes, err := io.Copy(newFile, file)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("coppied %v bytes\n", bytes)

}
