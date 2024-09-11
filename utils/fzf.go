package utils

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func FzfSelect(options []string) (int, string, error) {
	cmd := exec.Command("fzf")

	reader, writer := io.Pipe()
	cmd.Stdin = reader

	go func() {
		defer writer.Close()
		for index, option := range options {
			formattedOption := fmt.Sprintf("%d. %s\n", index+1, option)
			_, err := writer.Write([]byte(formattedOption))
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error has occurred")
				os.Exit(1)
			}
		}
	}()

	out, err := cmd.Output()
	if err != nil || len(out) == 0 {
		return -1, "", nil
	}
	outSplit := strings.Split(string(out), ".")

	index := outSplit[0]

	selectedOption, err := strconv.Atoi(index)
	if err != nil {
		return -1, "", err
	}

	option := outSplit[1]

	return selectedOption - 1, option, nil
}
