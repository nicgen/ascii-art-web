package ASCII

import (
	"strings"
)

// Handle the \n as a new line
func TransformAscii(input string, file_content []string) string {
	output := ""
	tmp := strings.Split(input, `\n`)
	for _, w := range tmp {
		if w != "" {
			output += CreateAscii(w, file_content)
		} else {
			output += "\n"
		}
	}
	return output
}

// Create an ASCII art from an input and a theme file
func CreateAscii(input string, file_content []string) (output string) {
	output = "" // tmp file for result
	for i := 0; i < 8; i++ {
		for _, j := range input {
			if rune(j) >= 32 && rune(j) <= 126 {
				index := (int(j)-32)*8 + i // index related to ascii (start at 32), 8 lines jump
				output += file_content[index]
				// fmt.Println("[CreateAscii]\n", output)
			} else {
				continue
			}
		}
		output += "\n"
	}
	return
}
