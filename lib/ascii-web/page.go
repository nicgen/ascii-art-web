package ASCIIWEB

import (
	"ASCII"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

// renderTemplate renders the template with given data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("static/html/templates", tmpl+".html") //Set the path of the html files we want to tmplPath, we join "templates" and the html file name
	t, err := template.ParseFiles(tmplPath)                          //Parse the template file to analyse it to find where to put data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, data) //execute the template and use the data inside the parse template
}

// homeHandler serves the main page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		Error(w, http.StatusNotFound, "not found")
		return
	} else {
		RenderTemplate(w, "index", nil)
	}
}

// asciiArtHandler handles the conversion of text to ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		Error(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}
	r.ParseForm()
	text := strings.ReplaceAll(r.FormValue("text"), "\r\n", `\n`) //replace \r\n of textarea new line by \n

	banner := r.Form.Get("banner")
	// Determine the banner file to use
	var themeFile string
	switch banner {
	case "standard":
		themeFile = "static/themes/standard.txt"
	case "shadow":
		themeFile = "static/themes/shadow.txt"
	case "thinkertoy":
		themeFile = "static/themes/thinkertoy.txt"
	}

	// file type

	file_type := r.Form.Get("type")
	// Determine the banner file to use
	var file_ext string
	switch file_type {
	case "texte":
		file_ext = ".txt"
	case "markdown":
		file_ext = ".nfo"
	case "nfo":
		file_ext = ".nfo"
		// case "thinkertoy":
		// 	file_ext = "static/themes/thinkertoy.txt"
	}

	url_dl := "static/export/ascii-art" + banner + file_ext
	fmt.Println(url_dl)

	// Convert the text to ASCII art
	file_content, err := ASCII.FileToLine(themeFile)
	if err != nil {
		Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ascii_art, err := ASCII.BothAscii(text, "static/export/ascii-art"+banner+file_ext, file_content)
	if err != nil {
		Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	data := map[string]string{
		"text":         text,
		"banner":       banner,
		"ascii_export": ascii_art,
		"url_dl":       url_dl,
		"file_ext":     file_ext,
	}

	RenderTemplate(w, "index", data)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to convert text to ASCII art: %v", err), http.StatusInternalServerError)
	// 	return
	// }
}

func Error(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	error_id := strconv.Itoa(status)
	tmpl, err := template.ParseFiles("static/html/templates/error.html")
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error parsing template: %v", err)
		return
	}
	data := map[string]string{
		"error_id": error_id,
		"message":  message,
	}
	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
		log.Printf("Error executing template: %v", err)
	}
}
