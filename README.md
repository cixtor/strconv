# String Conversion (strconv)  [![GoReport](https://goreportcard.com/badge/github.com/cixtor/strconv)](https://goreportcard.com/report/github.com/cixtor/strconv) [![GoDoc](https://godoc.org/github.com/cixtor/strconv?status.svg)](https://godoc.org/github.com/cixtor/strconv)

String functions are used in computer programming languages to manipulate a string or query information about a string (some do both). Most programming languages that have a string datatype will have some string functions although there may be other low-level ways within each language to handle strings directly. In object-oriented languages, string functions are often implemented as properties and methods of string objects. In functional and list-based languages a string is represented as a list (of character codes), therefore all list-manipulation procedures could be considered string functions. However such languages may implement a subset of explicit string- specific functions as well.

— From WikiPedia [String Functions, by Programming Languages](http://en.wikipedia.org/wiki/String_functions)

## Features

| Command | Description |
|---------|-------------|
| `replace` | Replace a text string with another |
| `capitalize` | Convert a text string into a capitalized version of its words |
| `uppercase` | Convert all the characters in a text string into their capital form |
| `lowercase` | Convert all the characters in a text string into their lower form |
| `md5` | Calculate the md5 hash of the string specified |
| `sha1` | Calculate the sha1 hash of the string specified |
| `chunk` | Splits a string into smaller pieces of the same size |
| `length` | Returns the length of the string specified |
| `b64enc` | Encodes data with MIME base64 |
| `b64dec` | Decodes data encoded with MIME base64 |
| `urldec` | Decodes URL-encoded string |
| `urlenc` | Encodes URL string with their correspondent hex digits |
| `rotate` | Perform a rotation on a string by the value specified |

## Usage

```sh
$ echo "Hello, World" | strconv replace o @
Hell@, W@rld

$ echo "hello, world" | strconv capitalize
Hello, World

$ echo "hello, world" | strconv uppercase
HELLO, WORLD

$ echo "HELLO, WORLD" | strconv lowercase
hello, world

$ echo "hello, world" | strconv md5
e4d7f1b4ed2e42d15898f4b27b019da4

$ echo "hello, world" | strconv sha1
b7e23ec29af22b0b4e41da31e868d57226121c84

$ echo "hello, world" | strconv length
12

$ echo "hello, world" | strconv b64enc
aGVsbG8sIHdvcmxk

$ echo "aGVsbG8sIHdvcmxk" | strconv b64dec
hello, world

$ echo "https://cixtor.com/?foo=bar&lorem=ipsum" | strconv urlenc
https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum

$ echo "https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum" | strconv urldec
https://cixtor.com/?foo=bar&lorem=ipsum

$ echo "hello, world" | strconv rotate
uryyb, jbeyq

$ echo "uryyb, jbeyq" | strconv rotate
hello, world
```

## Aliases

Add these aliases to your `~/.profile` file to maximize usability.

```sh
alias replace="strconv replace"
alias capitalize="strconv capitalize"
alias upper="strconv uppercase"
alias lower="strconv lowercase"
alias md5="strconv md5"
alias sha1="strconv sha1"
alias len="strconv length"
alias b64enc="strconv b64enc"
alias b64dec="strconv b64dec"
alias urldec="strconv urldec"
alias urlenc="strconv urlenc"
alias rotate="strconv rotate"
```
