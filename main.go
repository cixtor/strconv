package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Println("String Conversion")
		fmt.Println("https://cixtor.com/")
		fmt.Println("https://github.com/cixtor/strconv")
		fmt.Println("https://en.wikipedia.org/wiki/String_functions")
		fmt.Println("https://en.wikipedia.org/wiki/String_(computer_science)")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("  echo [text] | strconv replace [old] [new]")
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
		fmt.Println("  echo [text] | strconv rotate 13")
		fmt.Println()
		fmt.Println("Alias:")
		fmt.Println("  alias upper=\"strconv uppercase\"")
		fmt.Println("  alias lower=\"strconv lowercase\"")
		fmt.Println("  alias len=\"strconv length\"")
	}

	flag.Parse()

	if flag.Arg(0) == "" {
		flag.Usage()
		os.Exit(2)
	}

	reader := io.LimitReader(os.Stdin, 2<<22)
	input, err := ioutil.ReadAll(reader)

	if err != nil {
		fmt.Println("read err;", err)
		os.Exit(1)
	}

	switch flag.Arg(0) {
	case "replace":
		fmt.Printf("%s", replace(input, flag.Arg(1), flag.Arg(2)))
	case "capitalize":
		fmt.Printf("%s", capitalize(input))
	case "uppercase":
		fmt.Printf("%s", uppercase(input))
	case "lowercase":
		fmt.Printf("%s", lowercase(input))
	case "md5":
		fmt.Printf("%s\n", hashMD5(input))
	case "sha1":
		fmt.Printf("%s\n", hashSHA1(input))
	case "chunk":
		fmt.Printf("%s", chunk(input, flag.Arg(1)))
	case "length":
		fmt.Printf("%s\n", length(input, flag.Arg(1)))
	case "b64enc":
		fmt.Printf("%s\n", base64Encode(input))
	case "b64dec":
		fmt.Printf("%s", base64Decode(input))
	case "urlenc":
		fmt.Printf("%s\n", urlEncode(input))
	case "urldec":
		fmt.Printf("%s", urlDecode(input))
	case "rotate":
		fmt.Printf("%s", rotate(input, flag.Arg(1)))
	}
}
