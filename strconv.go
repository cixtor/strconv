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
}

// Application is the implementation.
type Application struct{}

func (app Application) Replace(text string, old string, _new string) string {
	return strings.Replace(text, old, _new, -1)
}
