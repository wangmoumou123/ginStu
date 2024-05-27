package tools

import (
	"fmt"
	"strings"
	"time"
)

// Log prints the log message with the specified color and allows multiple data inputs
func Log(color string, data ...interface{}) {
	formattedTime := time.Now().Format("2006-01-02 15:04:05")
	var colorCode string

	// Set color code based on input
	switch color {
	case "g":
		colorCode = "32" // green
	case "r":
		colorCode = "31" // red
	case "y":
		colorCode = "33" // yellow
	default:
		colorCode = "" // no color
	}

	// Convert the variadic arguments to a single string
	message := strings.TrimSpace(fmt.Sprint(data...))

	// Print with color if color code is provided, otherwise print without color
	if colorCode != "" {
		fmt.Printf("%s ==> \033[%sm%s\033[0m\n", formattedTime, colorCode, message)
	} else {
		fmt.Printf("%s ==> %s\n", formattedTime, message)
	}

}
