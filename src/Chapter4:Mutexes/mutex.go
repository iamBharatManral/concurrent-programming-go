package main

import (
	"fmt"
	"time"
)

func run(reader *int, writer *int, count *int) {
	read(reader, writer, count)
	write(reader, writer, count)
}

func read(reader *int, writer *int, count *int) {
	if *writer == 0 {
		*reader += 1
		fmt.Printf("Read: %v\n", *count)
		*reader -= 1
		return
	}
	read(reader, writer, count)
}

func write(reader *int, writer *int, count *int) {
	if *reader == 0 {
		*writer = 1
		*count += 1
		fmt.Printf("Write: %v\n", *count)
		*writer = 0
		return
	}
	write(reader, writer, count)
}

func main() {
	reader := 0
	writer := 0
	count := 0
	for i := 1; i < 10; i++ {
		go read(&reader, &writer, &count)
	}
	for i := 1; i < 10; i++ {
    		go write(&reader, &writer, &count)
    }
	time.Sleep(3 * time.Second)
}
