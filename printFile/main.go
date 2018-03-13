package main

import (
	"fmt"
	"os"
)

func main() {
	fileName := os.Args[1]
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	data := make([]byte, 100)
	count, _ := f.Read(data)
	fmt.Println(string(data))
	fmt.Println("Bytes:", count)

}
