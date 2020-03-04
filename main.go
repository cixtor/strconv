package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
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
		fmt.Println("  echo [text] | strconv chunk")
		fmt.Println("  echo [text] | strconv length")
		fmt.Println("  echo [text] | strconv b64enc")
		fmt.Println("  echo [text] | strconv b64dec")
		fmt.Println("  echo [text] | strconv urldec")
		fmt.Println("  echo [text] | strconv urlenc")
		fmt.Println("  echo [text] | strconv rotate")
		fmt.Println()
		fmt.Println("Alias:")
		fmt.Println("  alias upper=\"strconv uppercase\"")
		fmt.Println("  alias lower=\"strconv lowercase\"")
		fmt.Println("  alias len=\"strconv length\"")
		os.Exit(2)
	}

	flag.Parse()

	action := flag.Arg(0)

	if action == "" {
		flag.Usage()
	}

	body, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println("read err;", err)
		os.Exit(1)
	}
	text := strings.TrimSpace(string(body))

	switch action {
	case "replace":
		text = app.Replace(text, flag.Arg(1), flag.Arg(2))
	case "capitalize":
		text = app.Capitalize(text)
	case "uppercase":
		text = app.Uppercase(text)
	case "lowercase":
		text = app.Lowercase(text)
	case "md5":
		text = app.Md5(text)
	case "sha1":
		text = app.Sha1(text)
	case "chunk":
		text = app.Chunk(text, flag.Arg(1))
	case "length":
		text = fmt.Sprintf("%d", app.Length(text))
	case "b64enc":
		text = app.Base64Encode(text)
	case "b64dec":
		text = app.Base64Decode(text)
	case "urldec":
		text = app.URLDecode(text)
	case "urlenc":
		text = app.URLEncode(text)
	case "rotate":
		text = app.Rotate(text)
	}

	fmt.Println(text)
}
