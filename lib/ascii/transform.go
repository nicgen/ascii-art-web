package ASCII

import (
	"errors"
	"strings"
)

// Handle the \n as a new line
func TransformAscii(input string, file_content []string) (string, error) {
	output := ""
	tmp := strings.Split(input, `\n`)
	for _, w := range tmp {
		if w != "" {
			result, err := CreateAscii(w, file_content)
			if err != nil {
				return "", err
			}
			output += result
		} else {
			output += "\n"
		}
	}
	return output, nil
}

// Create an ASCII art from an input and a theme file
func CreateAscii(input string, file_content []string) (output string, err error) {
	output = "" // tmp file for result
	for i := 0; i < 8; i++ {
		for _, j := range input {
			if rune(j) >= 32 && rune(j) <= 126 {
				index := (int(j)-32)*8 + i // index related to ascii (start at 32), 8 lines jump
				output += file_content[index]
				// fmt.Println("[CreateAscii]\n", output)
			} else {
				return "", errors.New("range output of managed ascii table")
			}
		}
		output += "\n"
	}
	return
}
