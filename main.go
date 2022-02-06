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
		fmt.Println("  echo [text] | strconv replace [a:b m:n x:y]")
		fmt.Println("  echo [text] | strconv capitalize")
		fmt.Println("  echo [text] | strconv uppercase")
		fmt.Println("  echo [text] | strconv lowercase")
		fmt.Println("  echo [text] | strconv md5")
		fmt.Println("  echo [text] | strconv sha1")
		fmt.Println("  echo [text] | strconv chunk 64")
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
		os.Exit(2)
	}

	reader := io.LimitReader(os.Stdin, 2<<22)
	input, err := ioutil.ReadAll(reader)

	if err != nil {
		fmt.Println("read err;", err)
		os.Exit(1)
	}
	text := strings.TrimSpace(string(body))

	app.args = flag.Args()

	actions := map[string]Command{
		"replace":    app.Replace,
		"capitalize": app.Capitalize,
		"uppercase":  app.Uppercase,
		"lowercase":  app.Lowercase,
		"md5":        app.Md5,
		"sha1":       app.Sha1,
		"chunk":      app.Chunk,
		"length":     app.Length,
		"b64enc":     app.Base64Encode,
		"b64dec":     app.Base64Decode,
		"urldec":     app.URLDecode,
		"urlenc":     app.URLEncode,
		"rotate":     app.Rotate,
	}

	if function, ok := actions[action]; ok {
		text = function(text)
	}

	fmt.Println(text)
}
