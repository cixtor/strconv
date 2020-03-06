package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"
)

// Application is the implementation.
type Application struct{}

// Replace returns the given string with all occurrences of a search term
// replaced with a replacement term. This operation is case sensitive. If the
// replacement term is unspecified or empty, all occurrences of the search term
// will be removed from the string.
func (app Application) Replace(text string, old string, new string) string {
	/* support the replacement of new lines */
	old = strings.Replace(old, "\\n", "\n", -1)
	new = strings.Replace(new, "\\n", "\n", -1)
	/* support the replacement of hard tabs */
	old = strings.Replace(old, "\\t", "\t", -1)
	new = strings.Replace(new, "\\t", "\t", -1)
	return strings.Replace(text, old, new, -1)
}

// Capitalize will write a word with its first letter as a capital letter
// (upper-case letter) and the remaining letters in lower case in writing
// systems with a case distinction. The term is also used for the choice of
// case in text.
func (app Application) Capitalize(text string) string {
	return strings.Title(text)
}

// Uppercase - Letter case (or just case) is the distinction between the letters
// that are in larger upper case (also uppercase, capital letters, capitals,
// caps, large letters, or more formally majuscule) and smaller lower case (also
// lowercase, small letters, or more formally minuscule) in the written
// representation of certain languages. Capital letters only. This style can be
// used in headings and special situations, such as for typographical emphasis
// in text made on a typewriter.
func (app Application) Uppercase(text string) string {
	return strings.ToUpper(text)
}

// Lowercase - Letter case (or just case) is the distinction between the letters
// that are in larger upper case (also uppercase, capital letters, capitals,
// caps, large letters, or more formally majuscule) and smaller lower case (also
// lowercase, small letters, or more formally minuscule) in the written
// representation of certain languages. No capital letters. This style is
// sometimes used for artistic effect, such as in poetry. Also commonly seen in
// computer commands, and in SMS language (avoiding the shift key, to type more
// quickly).
func (app Application) Lowercase(text string) string {
	return strings.ToLower(text)
}

// Md5 calculates a message-digest fingerprint (checksum) for a file.
func (app Application) Md5(text string) string {
	hash := md5.New()

	io.WriteString(hash, text)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Sha1 message-digest algorithm. This algorithm takes a message and generates a
// 160-bit digest from the input. The SHA1 algorithm is related to the MD4
// algorithm but has been strengthend against certain types of cryptographic
// attack. SHA1 should be used in preference to MD4 or MD5 in new applications.
func (app Application) Sha1(text string) string {
	hash := sha1.New()

	io.WriteString(hash, text)

	return fmt.Sprintf("%x", hash.Sum(nil))
}

// Chunk splits a string into smaller pieces, default: 64.
func (app Application) Chunk(text string, length string) string {
	limit, err := strconv.Atoi(length)

	if err != nil {
		limit = 64
	}

	var counter int
	var result string
	total := len(text)
	for i := 0; i < total; i++ {
		counter++
		result += text[i : i+1]
		if counter >= limit {
			counter = 0
			result += "\n"
		}
	}
	return result
}

// Length measures the arbitrary (but finite) length of a chain of letters.
// Although formal strings can have an arbitrary (but finite) length, the length
// of strings in real languages is often constrained to an artificial maximum.
// In general, there are two types of string datatypes: fixed-length strings,
// which have a fixed maximum length to be determined at compile time and which
// use the same amount of memory whether this maximum is needed or not, and
// variable-length strings, whose length is not arbitrarily fixed and which can
// use varying amounts of memory depending on the actual requirements at run
// time. Most strings in modern programming languages are variable-length
// strings. Of course, even variable-length strings are limited in length – by
// the number of bits available to a pointer, and by the size of available
// computer memory. The string length can be stored as a separate integer (which
// may put an artificial limit on the length) or implicitly through a
// termination character, usually a character value with all bits zero.
func (app Application) Length(text string) int {
	return len(text)
}

// Base64Encode encodes data with MIME base64. Base64 is a group of similar
// binary-to-text encoding schemes that represent binary data in an ASCII string
// format by translating it into a radix-64 representation. The term Base64
// originates from a specific MIME content transfer encoding. Each base64 digit
// represents exactly 6 bits of data.
func (app Application) Base64Encode(text string) string {
	return base64.StdEncoding.EncodeToString([]byte(text))
}

// Base64Decode decodes data encoded with MIME base64. Base64 is a group of
// similar binary-to-text encoding schemes that represent binary data in an
// ASCII string format by translating it into a radix-64 representation. The
// term Base64 originates from a specific MIME content transfer encoding. Each
// base64 digit represents exactly 6 bits of data.
func (app Application) Base64Decode(text string) string {
	out, err := base64.StdEncoding.DecodeString(text)

	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s", out)
}

// URLEncode encodes a chain of letters using the percent-encoding technique.
// Percent-encoding, also known as URL encoding, is a mechanism for encoding
// information in a Uniform Resource Identifier (URI) under certain
// circumstances. Although it is known as URL encoding it is, in fact, used more
// generally within the main Uniform Resource Identifier (URI) set, which
// includes both Uniform Resource Locator (URL) and Uniform Resource Name (URN).
// As such, it is also used in the preparation of data of the application/x-www-
// form-urlencoded media type, as is often used in the submission of HTML form
// data in HTTP requests.
func (app Application) URLEncode(text string) string {
	return url.QueryEscape(text)
}

// URLDecode decodes a chain of letters using the percent-encoding technique.
// Percent-encoding, also known as URL encoding, is a mechanism for encoding
// information in a Uniform Resource Identifier (URI) under certain
// circumstances. Although it is known as URL encoding it is, in fact, used more
// generally within the main Uniform Resource Identifier (URI) set, which
// includes both Uniform Resource Locator (URL) and Uniform Resource Name (URN).
// As such, it is also used in the preparation of data of the application/x-www-
// form-urlencoded media type, as is often used in the submission of HTML form
// data in HTTP requests.
func (app Application) URLDecode(text string) string {
	out, err := url.QueryUnescape(text)

	if err != nil {
		panic(err)
	}

	return out
}

// Rotate performs the rot13 transform on a string. The ROT13 encoding simply
// shifts every letter by 13 places in the alphabet while leaving non-alpha
// characters untouched. Encoding and decoding are done by the same function,
// passing an encoded string as argument will return the original version.
func (app Application) Rotate(text string) string {
	rotator := func(char rune) rune {
		number := 13
		switch {
		case char >= 'A' && char <= 'Z':
			return 'A' + (char-'A'+rune(number))%26
		case char >= 'a' && char <= 'z':
			return 'a' + (char-'a'+rune(number))%26
		}
		return char
	}

	return strings.Map(rotator, text)
}
