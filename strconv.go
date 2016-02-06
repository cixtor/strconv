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
