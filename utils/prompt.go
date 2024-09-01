package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// returns index and item
func Prompt(options []string) (int, string) {

	reader := bufio.NewReader(os.Stdin)
	for index, option := range options {
		fmt.Printf("%d. %s\n", index+1, option)
	}

	fmt.Println("Select an option:")
	input, _ := reader.ReadString('\n')
	input = strings.Split(input, "\n")[0]

	selectedOption, err := strconv.Atoi(input)

	if err != nil || selectedOption-1 < 0 || selectedOption-1 > len(options) {
		fmt.Println("Invalid option")
		fmt.Println(err.Error())
		return Prompt(options)
	}

	return selectedOption - 1, options[selectedOption-1]
}
