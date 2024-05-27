package main

import (
	"fmt"
	"time"
)

func Log(data interface{}, color string) {
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	var colorCode string
	switch color {
	case "g":
		colorCode = "32"
	case "r":
		colorCode = "31"
	case "y":
		colorCode = "33"
	default:
		colorCode = ""
	}

	if colorCode != "" {
		fmt.Printf("%s ==> \033[%sm%s\033[0m\n", formattedTime, colorCode, data)
	} else {
		fmt.Printf("%s ==> %s\n", formattedTime, data)
	}

}
