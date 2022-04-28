package cmd

import "fmt"

func PromptString(msg string) string {
	fmt.Println()
	fmt.Println(msg)
	var toReturn string
	fmt.Scan(&toReturn)
	return toReturn
}
