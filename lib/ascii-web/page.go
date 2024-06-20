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
	tmplPath := filepath.Join("static/html/templates", tmpl) //Set the path of the html files we want to tmplPath, we join "templates" and the html file name
	t, err := template.ParseFiles(tmplPath)                  //Parse the template file to analyse it to find where to put data
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// fmt.Println(data)
	t.Execute(w, data) //execute the template and use the data inside the parse template
}

// homeHandler serves the main page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// RenderTemplate(w, "index.html", nil)
	if r.URL.Path != "/" {
		w.WriteHeader(404) // return error 404 (forced, bad practice)
		RenderTemplate(w, "404.html", nil)
	} else {
		RenderTemplate(w, "index.html", nil)
	}
}

// asciiArtHandler handles the conversion of text to ASCII art
func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	r.ParseForm()
	// text := ""
	// fmt.Printf("TEXT:%v, type: %T", r.FormValue("text"), r.FormValue("text"))
	text := strings.ReplaceAll(r.FormValue("text"), "\r\n", `\n`) //replace \r\n of textarea new line by \n
	// text := strings.ReplaceAll(r.Form.Get("text"), "\r\n", `\n`)
	// text := r.FormValue("text")
	// fmt.Println("TEXT:", r.Form.Get("text"))
	banner := r.Form.Get("banner")
	fmt.Println("BANNER:", r.Form.Get("banner"))

	// Determine the banner file to use
	var themeFile string
	switch banner {
	case "standard":
		themeFile = "static/themes/standard.txt"
	case "shadow":
		themeFile = "static/themes/shadow.txt"
	case "thinkertoy":
		themeFile = "static/themes/thinkertoy.txt"
	default:
		http.Error(w, "Invalid banner", http.StatusBadRequest)
		return
	}

	// Convert the text to ASCII art
	// var file_content []string
	file_content := ASCII.FileToLine(themeFile)
	// fmt.Println(file_content)
	ascii_art := ASCII.BothAscii(text, "static/export/ascii-art.txt", file_content)
	// ASCII.BothAscii(text, "ascii_art.txt", file_content)

	data := map[string]string{
		"text":         text,
		"banner":       banner,
		"ascii_export": ascii_art,
	}

	fmt.Printf("input: %v theme: %v", data["text"], data["banner"])
	// fmt.Printf("[DATA]Input: %v", data["text"])
	// fmt.Println("[DATA]banner:", data["banner"])
	// fmt.Println("[DATA]ascii:\n", data["ascii_export"])
	RenderTemplate(w, "index.html", data)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Failed to convert text to ASCII art: %v", err), http.StatusInternalServerError)
	// 	return
	// }
}
