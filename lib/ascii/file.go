package ASCII

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
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

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filePath := "static/export/file.txt" // Replace with the actual path to your file
	file, err := os.Open(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	defer file.Close()

	fileStat, err := file.Stat()
	if err != nil {
		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
		return
	}

	fileName := filepath.Base(filePath)

	// Set the headers
	w.Header().Set("Content-Type", "application/txt") // Replace with the appropriate MIME type for your file
	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

	// Serve the file
	http.ServeFile(w, r, filePath)
}

// func DownloadHandler(w http.ResponseWriter, r *http.Request) {
// 	format := r.URL.Query().Get("format")
// 	if format != "nfo" && format != "txt" {
// 		http.Error(w, "Invalid format", http.StatusBadRequest)
// 		return
// 	}

// 	filePath := fmt.Sprintf("/static/export/file%s", format) // Replace with actual path
// 	fmt.Println("FILEPATH: ", filePath)
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		http.Error(w, "File not found", http.StatusNotFound)
// 		return
// 	}
// 	defer file.Close()

// 	fileStat, err := file.Stat()
// 	if err != nil {
// 		http.Error(w, "Failed to get file info", http.StatusInternalServerError)
// 		return
// 	}

// 	fileName := filepath.Base(filePath)

// 	// Set the headers
// 	contentType := "text/plain"
// 	if format == "nfo" {
// 		contentType = "application/x-nfo"
// 	}

// 	w.Header().Set("Content-Type", contentType)
// 	w.Header().Set("Content-Length", fmt.Sprintf("%d", fileStat.Size()))
// 	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))

// 	// Serve the file
// 	fmt.Println("FILEPATH: ", filePath)
// 	http.ServeFile(w, r, filePath)
// }

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
		// WriteFile(output, content_formatted)
		WriteFile("static/export/file.txt", content_formatted)
		// WriteFile("static/export/file.nfo", content_formatted)
	} else {
		fmt.Printf("%s", content)
	}
	return content, nil
}
