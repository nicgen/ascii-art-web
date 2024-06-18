package ASCIIWEB

import (
	"ASCII"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func LoadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// renderTemplate renders the template with given data
func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("src/html/templates", tmpl) //Set the path of the html files we want to tmplPath, we join "templates" and the html file name
	t, err := template.ParseFiles(tmplPath)               //Parse the template file to analyse it to find where to put data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(data)
	t.Execute(w, data) //execute the template and use the data inside the parse template
}

// homeHandler serves the main page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index.html", nil)
}

// asciiArtHandler handles the conversion of text to ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	r.ParseForm()
	text := ""
	fmt.Println("TEXT:", r.FormValue("text"))
	text = strings.ReplaceAll(r.FormValue("text"), "\r\n", "\n") //replace \r\n of textarea new line by \n
	fmt.Println("TEXT:", r.FormValue("text"))
	banner := r.FormValue("banner")
	fmt.Println("BANNER:", r.FormValue("banner"))

	// Determine the banner file to use
	var themeFile string
	switch banner {
	case "standard":
		themeFile = "src/themes/standard.txt"
	case "shadow":
		themeFile = "src/themes/shadow.txt"
	case "thinkertoy":
		themeFile = "src/themes/thinkertoy.txt"
	default:
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	// Convert the text to ASCII art
	// var file_content []string
	file_content := ASCII.FileToLine(themeFile)
	// fmt.Println(file_content)
	ascii_art := ASCII.TransformAscii(text, file_content)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to convert text to ASCII art: %v", err), http.StatusInternalServerError)
	// 	return
	// }

	data := map[string]string{
		"text":         text,
		"banner":       banner,
		"ascii_export": ascii_art,
	}

	fmt.Println("[DATA]Input:", data["text"])
	fmt.Println("[DATA]banner:", data["banner"])
	fmt.Println("[DATA]ascii:\n", data["ascii_export"])

	RenderTemplate(w, "result.html", data)
}
