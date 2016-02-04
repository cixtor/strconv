package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var text string
	var app Application

	flag.Usage = func() {
		fmt.Println("String Conversion")
		fmt.Println("https://cixtor.com/")
		fmt.Println("https://github.com/cixtor/strconv")
		fmt.Println("https://en.wikipedia.org/wiki/String_functions")
		fmt.Println("https://en.wikipedia.org/wiki/String_(computer_science)")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  echo [text] | strconv replace")
		fmt.Println("  echo [text] | strconv capitalize")
		fmt.Println("  echo [text] | strconv uppercase")
		flag.PrintDefaults()
		os.Exit(2)
	}

	action := flag.Arg(1)

	if action == "" {
		flag.Usage()
	}

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text += scanner.Text()
	}

	switch action {
	case "replace":
		fmt.Println(app.Replace(text, flag.Arg(2), flag.Arg(3)))
		break
	case "capitalize":
		fmt.Println(app.Capitalize(text))
		break
	case "uppercase":
		fmt.Println(app.Uppercase(text))
		break
	}
}
