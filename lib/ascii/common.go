package ASCII

import (
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

// accepts: arguments / returns: theme, input, options
func CheckArgs(args []string) (valid bool, input, output, theme, align string, color []string) {
	// VARS
	input, output, theme, align = "", "", "standard", "left"
	args_len := len(args)

	switch args_len {
	case 0: // error no argument
		valid = false
	case 1: // input only
		valid = true
		if args[0] == "--help" {
			valid = false
		} else {
			valid = true
			input = args[0]
			theme = "src/themes/standard.txt"
		}
	default: // option(s) + input + theme
		valid, input, output, theme, align, color = CheckOption(args)
	}
	return
}

// valid bool, input, output, theme, align string, color []
func CheckOption(arg []string) (valid bool, input, output, theme, align string, color []string) {
	theme = "standard"
	valid = false          // default value
	var undefined []string // tmp value for !options (input/theme)
	// REGEX
	// r := regexp.MustCompile(`(?:^--(color|output|align|help)=?([\w\.\-_]+)?)`)
	r := regexp.MustCompile(`^--(?:color|output|align|help)(?:=[\w\.\-_]+)?$`) // detect if arg is a valid option
	t := 0                                                                     // toggle to check if next arg after color is an option
	for i, e := range arg {                                                    // loop trough arguments
		if t == 1 {
			t = 0 // reset toggle
			continue
		}
		matched := r.MatchString(e)
		if matched { // element (o) is an option
			isOption, optionType := OptionType(e)
			if isOption {
				switch optionType {
				case "align":
					// fmt.Println("option is align")
					x := strings.Split(e, "=")[1]
					if x == "left" || x == "right" || x == "center" || x == "justify" {
						align = x
						valid = true
					}
				case "color":
					if i < len(arg) { // check if arg[i+1] exists
						x, _ := OptionType(arg[i+1])
						if !x { // check if next arg is an option
							x := strings.Split(e, "=")[1]
							color = append(color, x)
							color = append(color, arg[i+1])
							t = 1 // because we want to skip the next value
							valid = true
						} else {
							PrintHelp()
							os.Exit(1)
						}
					} else {
						PrintHelp()
						os.Exit(2)
					}
				case "output":
					// fmt.Println("option is output")
					output = strings.Split(e, "=")[1]
					output = "./output/" + output
					valid = true
				case "help":
					PrintHelp()
					os.Exit(3)
				}
			}
		} else { // if not option must be input (+ theme)?
			undefined = append(undefined, e)
		}
	}
	// undefined contains the lasts arguments (input || input + theme)
	// fmt.Println("[UNDEFINED]", undefined[len(undefined)-2])
	if len(undefined) == 1 { // if only one value = input
		input = undefined[0]
		valid = true
	} else if len(undefined) == 2 { // if two values input + theme
		input = undefined[len(undefined)-2]
		if IsThemeValid(undefined[len(undefined)-1]) {
			theme = undefined[len(undefined)-1]
			valid = true
		} else {
			PrintHelp()
			os.Exit(4)
		}
	} else { // if more value undefined value = error (show help)
		PrintHelp()
		os.Exit(5)
	}
	theme = "./src/themes/" + theme + ".txt"
	// fmt.Println("[CheckOption]", "\nvalid: ", valid, "\ninput:", input, "\noutput:", output, "\ntheme:", theme, "\nalign:", align, "\ncolor:", color)
	return
}

func OptionType(str string) (isOption bool, optionType string) {
	isOption = true
	if strings.HasPrefix(str, "--align=") {
		optionType = "align"
	} else if strings.HasPrefix(str, "--color=") {
		optionType = "color"
	} else if strings.HasPrefix(str, "--output=") {
		optionType = "output"
	} else if str == "--help" {
		optionType = "help"
	} else {
		isOption = false
	}
	return
}

// check if theme var is valid
func IsThemeValid(str string) bool {
	theme_list := []string{"standard", "shadow", "stars", "thinkertoy"}
	for _, t := range theme_list {
		if str == t {
			return true
		}
	}
	return false
}

// get color & string, returns range index and ANSI color code
// func AsciiColor(str string) (color string) {
// 	t := strings.Split(str, "--color=")
// 	return string(t[1])
// 	// var Reset = "\033[0m"
// 	// var Red = "\033[31m"
// 	// if k > 1 && k < 4 {
// 	// 	// fmt.Printf("%s%s%s", Red, fileLines[index], Reset)
// 	// 	export += Red + fileLines[index] + Reset
// 	// } else {
// 	// 	export += fileLines[index]
// 	// 	// fmt.Print(fileLines[index])
// 	// }
// }

// todo align
// get User terminal width & height
func GetTerminalSize() (width, height int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	height, err = strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}

	width, err = strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return
}

// ?USAGE: GetTerminalSize
// THeight, _ := GetTerminalSize()
// w := THeight
// // h := TWidth
// s := "in the middle"
// res := fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s))/2, s))
// fmt.Printf("%s\n", res)
// ros := fmt.Sprintf("%[1]*s", -w, fmt.Sprintf("%[1]*s", (w+len(s)), s))
// fmt.Printf("%s\n", ros)
