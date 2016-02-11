package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strings"
)

// Converter is the abstract object.
type Converter interface {
	Replace(string) string
	Capitalize(string) string
	Uppercase(string) string
	Lowercase(string) string
	Md5(string) string
	Sha1(string) string
	Length(string) int
	Base64Encode(string) string
	Base64Decode(string) string
	UrlDecode(string) string
}

// Application is the implementation.
type Application struct{}

func (app Application) Replace(text string, old string, _new string) string {
	return strings.Replace(text, old, _new, -1)
}

func (app Application) Capitalize(text string) string {
	return strings.Title(text)
}

func (app Application) Uppercase(text string) string {
	return strings.ToUpper(text)
}

func (app Application) Lowercase(text string) string {
	return strings.ToLower(text)
}

func (app Application) Md5(text string) string {
	hash := md5.New()

	io.WriteString(hash, text)

	return fmt.Sprintf("%x", hash.Sum(nil))

}

func (app Application) Sha1(text string) string {
	hash := sha1.New()

	io.WriteString(hash, text)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

func (app Application) Length(text string) int {
	return len(text)
}

func (app Application) Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

func (app Application) Base64Decode(text string) string {
	out, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", out)
}

func (app Application) UrlDecode(text string) string {
	out, err := url.QueryUnescape(text)

	if err != nil {
		panic(err)
	}

	return out
}
