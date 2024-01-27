package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	data := make([]int, 0)

	for i := 0; i < 10000000; i++ {
		data = append(data, 1)
	}

	file, err := os.Create("20mb.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Convert integers to a string and then to bytes
	for _, num := range data {
		str := strconv.Itoa(num)        // convert number to string
		_, err := file.WriteString(str) // write string to file
		if err != nil {
			log.Fatal(err)
		}
	}
}
