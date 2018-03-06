package main

import "fmt"

func main() {
	colours := map[string]string{
		"white":  "#ffffff",
		"red":    "#f45ghf",
		"yellow": "#f123456",
	}
	colours["black"] = "#000000"

	fmt.Println(colours)

}
