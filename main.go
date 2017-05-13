package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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
	text := string(body)

	switch action {
	case "replace":
		fmt.Println(app.Replace(text, flag.Arg(1), flag.Arg(2)))
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

	case "chunk":
		fmt.Println(app.Chunk(text, flag.Arg(1)))
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
		fmt.Println(app.URLDecode(text))
		break

	case "urlenc":
		fmt.Println(app.URLEncode(text))
		break

	case "rotate":
		fmt.Println(app.Rotate(text))
		break
	}
}
