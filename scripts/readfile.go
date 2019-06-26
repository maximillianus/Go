package main

import (
	"fmt"
	"io/ioutil"
	"log"
    "os"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println("FILES:", files)

    fi,err := os.Stat("./logging.go")
    fmt.Println("Stat:", fi.Size())
	for _, file := range files {
		fmt.Println(file.Name())
        fmt.Printf("FileSize: %d\n", file.Size())
	}
}
