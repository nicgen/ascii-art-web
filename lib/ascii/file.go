package ASCII

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

// Open a file & return a slice of line
func FileToLine(theme string) (file_content []string) {
	// Open theme file
	file, err := os.Open(theme)
	if err != nil {
		panic(err)
	}
	// Close file on exit and check for its returned error
	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()

	// Bufio
	// var file_content []string // new slice to write the content of the file (theme) by lines without empty lines
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" { // if line is not empty, add to the slice
			file_content = append(file_content, scanner.Text())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	file_content = append(file_content, "")
	return
}

func WriteFile(output, content string) {
	err := os.WriteFile(output, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("The result file: written successfully on " + output + ".")
}

// Print or save to a file
func ExportAscii(input, output string, file_content []string) {
	content := TransformAscii(input, file_content)
	if output != "" {
		WriteFile(output, content)
	} else {
		fmt.Printf("%s", content)
	}
}

// Return ASCII and export to a txt file
func BothAscii(input, output string, file_content []string) string {
	// fmt.Println("[BothAscii]: ", output)
	content := TransformAscii(input, file_content)
	if output != "" {
		WriteFile(output, content)
	} else {
		fmt.Printf("%s", content)
	}
	return content
}
