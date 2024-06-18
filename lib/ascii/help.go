package ASCII

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// todo add a error msg
// print help
func PrintHelp() {
	const minwidth, padding = 0, 3
	w := tabwriter.NewWriter(os.Stdout, minwidth, 0, padding, ' ', tabwriter.DiscardEmptyColumns)
	// output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint (|tabwriter.Debug)

	fmt.Fprintln(w, "Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII.")
	fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "Usage:\t")
	fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "\tgo <option> [input] <theme>\t")
	fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "Options:\t")
	fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "\tTheme :\tchange the theme of the text\t")
	fmt.Fprintln(w, " \t \t[standard|shadow|thinkertoy]\t")
	fmt.Fprintln(w, "\t")
	// fmt.Fprintln(w, "\tAlign :\talign the text\t")
	// fmt.Fprintln(w, " \t \t--align=[center|justify|left]\t")
	// fmt.Fprintln(w, "\t")
	// fmt.Fprintln(w, "\tColor :\tchange the color of the text\t")
	// fmt.Fprintln(w, " \t \t--color=[black|red|green|orange|blue|purple|cyan| <text to color>\t")
	// fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "\tOutput:\twrite the result in a file\t")
	fmt.Fprintln(w, " \t \t--output=fileName.txt\t")
	fmt.Fprintln(w, "\t")
	fmt.Fprintln(w, "\tHelp :\tshow this guide\t")
	fmt.Fprintln(w, " \t \t--help\t")
	// Flush the Writer to ensure all data is written to the output.
	w.Flush()
}
