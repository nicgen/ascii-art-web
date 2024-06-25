package ASCII

import (
	"os"
)

// Takes argument(s) and print/export it as an ascii art
func Ascii() {
	// Read arguments
	valid, input, output, theme, _, _ := CheckArgs(os.Args[1:])
	// * options: valid, input, output(file), theme(file), option(s)
	// fmt.Println("[Ascii]", "\nvalid: ", valid, "\ninput:", input, "\noutput:", output, "\ntheme:", theme, "\nalign:", align, "\ncolor:", color)

	// check if the arguments are valid
	if !valid {
		PrintHelp()
		os.Exit(0)
	}

	// Open the ASCII theme file
	file_content, err := FileToLine(theme)
	if err != nil {
		panic(err)
	}
	// Export/Print the result
	ExportAscii(input, output, file_content)
}
