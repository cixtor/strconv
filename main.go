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
		fmt.Println("  echo [text] | strconv lowercase")
		fmt.Println("  echo [text] | strconv md5")
		fmt.Println("  echo [text] | strconv sha1")
		fmt.Println("  echo [text] | strconv length")
		fmt.Println("  echo [text] | strconv b64enc")
		fmt.Println("  echo [text] | strconv b64dec")
		fmt.Println("  echo [text] | strconv urldec")
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
	case "lowercase":
		fmt.Println(app.Lowercase(text))
		break
	case "md5":
		fmt.Println(app.Md5(text))
		break
	case "sha1":
		fmt.Println(app.Sha1(text))
		break
	case "length":
		fmt.Println(app.Length(text))
		break
	case "b64enc":
		fmt.Println(app.Base64Encode(text))
		break
	case "b64dec":
		fmt.Println(app.Base64Decode(text))
		break
	case "urldec":
		fmt.Println(app.UrlDecode(text))
		break
	}
}
