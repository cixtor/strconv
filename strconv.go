package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

var action = flag.String("action", "none", "String convertion that will be executed")
var text = flag.String("text", "", "Text string that will be processed")
var old_str = flag.String("old", "", "Text string that will be replaced")
var new_str = flag.String("new", "", "Text string that will replace the old one")
var number = flag.Int("num", 13, "Positions to shift the text in the alphabet")

var replace = flag.Bool("replace", false, "Replace a text string with another")
var capitalize = flag.Bool("capitalize", false, "Convert a text string into a capitalized version of its words")
var uppercase = flag.Bool("uppercase", false, "Convert all the characters in a text string into their capital form")
var lowercase = flag.Bool("lowercase", false, "Convert all the characters in a text string into their lower form")
var hash_md5 = flag.Bool("md5", false, "Calculate the md5 hash of the string specified")
var hash_sha1 = flag.Bool("sha1", false, "Calculate the sha1 hash of the string specified")
var length = flag.Bool("length", false, "Returns the length of the string specified")
var base64_enc = flag.Bool("b64enc", false, "Encodes data with MIME base64")
var base64_dec = flag.Bool("b64dec", false, "Decodes data encoded with MIME base64")
var url_decode = flag.Bool("urldec", false, "Decodes URL-encoded string")
var url_encode = flag.Bool("urlenc", false, "Encodes URL string with their correspondent hex digits")
var rotate = flag.Bool("rotate", false, "Perform a rotation on a string by the value specified")

func main() {
	flag.Usage = func() {
		fmt.Println("String Conversion")
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://github.com/cixtor/mamutools")
		fmt.Println("  http://en.wikipedia.org/wiki/String_(computer_science)")
		fmt.Println("  http://en.wikipedia.org/wiki/String_functions")
		fmt.Println("Usage:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *action == "replace" || *replace == true {
		fmt.Printf("%s\n", strings.Replace(*text, *old_str, *new_str, -1))
		os.Exit(0)
	}

	if *action == "capitalize" || *capitalize == true {
		fmt.Printf("%s\n", strings.Title(*text))
		os.Exit(0)
	}

	if *action == "uppercase" || *uppercase == true {
		fmt.Printf("%s\n", strings.ToUpper(*text))
		os.Exit(0)
	}

	if *action == "lowercase" || *lowercase == true {
		fmt.Printf("%s\n", strings.ToLower(*text))
		os.Exit(0)
	}

	if *action == "md5" || *hash_md5 == true {
		hash := md5.New()
		io.WriteString(hash, *text)
		fmt.Printf("%x\n", hash.Sum(nil))
		os.Exit(0)
	}

	if *action == "sha1" || *hash_sha1 == true {
		hash := sha1.New()
		io.WriteString(hash, *text)
		fmt.Printf("%x\n", hash.Sum(nil))
		os.Exit(0)
	}

	if *action == "length" || *length == true {
		fmt.Printf("%d\n", len(*text))
		os.Exit(0)
	}

	if *action == "b64enc" || *base64_enc == true {
		fmt.Printf("%s\n", base64.StdEncoding.EncodeToString([]byte(*text)))
		os.Exit(0)
	}

	if *action == "b64dec" || *base64_dec == true {
		data, err := base64.StdEncoding.DecodeString(*text)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Printf("%s\n", data)
		os.Exit(0)
	}

	if *action == "urldec" || *url_decode == true {
		result, err := url.QueryUnescape(*text)
		if err == nil {
			fmt.Printf("%s\n", result)
			os.Exit(0)
		} else {
			fmt.Printf("Error decoding url: %s\n", err)
			os.Exit(1)
		}
	}

	if *action == "urlenc" || *url_encode == true {
		fmt.Printf("%s\n", url.QueryEscape(*text))
		os.Exit(0)
	}

	if *action == "rotate" || *rotate == true {
		rotator := func(letter rune) rune {
			switch {
			case letter >= 'A' && letter <= 'Z':
				return 'A' + (letter-'A'+rune(*number))%26
			case letter >= 'a' && letter <= 'z':
				return 'a' + (letter-'a'+rune(*number))%26
			}
			return letter
		}
		fmt.Printf("%s\n", strings.Map(rotator, *text))
		os.Exit(0)
	}

	flag.Usage()
	fmt.Printf("Error. Action specified is not allowed\n")
	os.Exit(1)
}
