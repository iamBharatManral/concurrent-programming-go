package main

import (
	"fmt"
	"os"
	"regexp"
	"time"
)

func readAndOutput(fileName string, pattern string) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading the file %v. Skipping the file\n", fileName)
		return
	}
	re := regexp.MustCompile(pattern)
	fmt.Printf("%v: %q\n", fileName, re.FindAll(data, -1))
}

func main() {
	for i := 2; i < len(os.Args); i++ {
		go readAndOutput(os.Args[i], os.Args[1])
	}
	time.Sleep(2 * time.Second)
}
