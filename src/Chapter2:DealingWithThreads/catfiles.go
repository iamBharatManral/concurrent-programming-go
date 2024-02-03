package main

import (
	"fmt"
	"os"
	"time"
)

func readAndOutput(fileName string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading the file %v. Skipping the file\n", fileName)
		return
	}
	fmt.Printf("%s\n", data)
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		go readAndOutput(os.Args[i])
	}
	time.Sleep(2 * time.Second)
}
