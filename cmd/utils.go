package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func getUserInput(displayMsg string) string {
	fmt.Println(displayMsg)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	//fmt.Println(input)
	return input
}

func pathExists(path string) bool {
	// Check if directory/ path exists, if not create one
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}
	return true
}
