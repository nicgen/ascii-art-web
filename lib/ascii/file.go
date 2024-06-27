package ASCII

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// Open a file & return a slice of line
func FileToLine(theme string) (file_content []string, err error) {
	// Open theme file
	file, err := os.Open(theme)
	if err != nil {
		return nil, errors.New("the selected theme does not exist")
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
	return file_content, nil
}

func WriteFile(output, content string) {
	err := os.WriteFile(output, []byte(content), 0644)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println("The result file: written successfully on " + output + ".")
}

// Print or save to a file
func ExportAscii(input, output string, file_content []string) error {
	content, err := TransformAscii(input, file_content)
	if err != nil {
		return err
	}
	if output != "" {
		WriteFile(output, content)
	} else {
		fmt.Printf("%s", content)
	}
	return nil
}

// Return ASCII and export to a txt file
func BothAscii(input, output, ext string, file_content []string) (string, error) {
	fmt.Println("[BothAscii]: ", output, " ext:", ext)
	content, err := TransformAscii(input, file_content)
	var content_formatted string
	if err != nil {
		return "", err
	}
	if output != "" {
		if ext == ".md" {
			content_formatted = "# ASCII_ART_EXPORT\n\n```text\n" + content + "\n```"
		} else {
			content_formatted = content
		}
		WriteFile(output, content_formatted)
	} else {
		fmt.Printf("%s", content)
	}
	return content, nil
}
