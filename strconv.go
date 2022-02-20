package main

import (
	"bytes"
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"net/url"
	"strconv"
	"unicode/utf8"
)

// replace returns the given string with all occurrences of a search term
// replaced with a replacement term. This operation is case sensitive. If the
// replacement term is unspecified or empty, all occurrences of the search term
// will be removed from the string.
func replace(text []byte, viejo string, nuevo string) []byte {
	return bytes.ReplaceAll(text, []byte(viejo), []byte(nuevo))
}

// capitalize will write a word with its first letter as a capital letter
// (upper-case letter) and the remaining letters in lower case in writing
// systems with a case distinction. The term is also used for the choice of
// case in text.
func capitalize(text []byte) []byte {
	for _, b := range text {
		if b >= utf8.RuneSelf {
			return bytes.Title(text)
		}
	}
	result := make([]byte, len(text))
	start := true
	for i, b := range text {
		if start && b >= 'a' && b <= 'z' {
			result[i] = b - 'a' + 'A'
		} else {
			result[i] = b
		}
		switch b {
		case ' ', '\t', '\n', '\r', '\f', '\v':
			start = true
		default:
			start = false
		}
	}
	return result
}

// uppercase - Letter case (or just case) is the distinction between the letters
// that are in larger upper case (also uppercase, capital letters, capitals,
// caps, large letters, or more formally majuscule) and smaller lower case (also
// lowercase, small letters, or more formally minuscule) in the written
// representation of certain languages. Capital letters only. This style can be
// used in headings and special situations, such as for typographical emphasis
// in text made on a typewriter.
func uppercase(text []byte) []byte {
	hasLower := false
	for _, b := range text {
		if b >= utf8.RuneSelf {
			return bytes.ToUpper(text)
		}
		if b >= 'a' && b <= 'z' {
			hasLower = true
		}
	}
	if !hasLower {
		return append([]byte(nil), text...)
	}
	result := make([]byte, len(text))
	for i, b := range text {
		if b >= 'a' && b <= 'z' {
			result[i] = b - 'a' + 'A'
		} else {
			result[i] = b
		}
	}
	return result
}

// lowercase - Letter case (or just case) is the distinction between the letters
// that are in larger upper case (also uppercase, capital letters, capitals,
// caps, large letters, or more formally majuscule) and smaller lower case (also
// lowercase, small letters, or more formally minuscule) in the written
// representation of certain languages. No capital letters. This style is
// sometimes used for artistic effect, such as in poetry. Also commonly seen in
// computer commands, and in SMS language (avoiding the shift key, to type more
// quickly).
func lowercase(text []byte) []byte {
	hasUpper := false
	for _, b := range text {
		if b >= utf8.RuneSelf {
			return bytes.ToLower(text)
		}
		if b >= 'A' && b <= 'Z' {
			hasUpper = true
		}
	}
	if !hasUpper {
		return append([]byte(nil), text...)
	}
	result := make([]byte, len(text))
	for i, b := range text {
		if b >= 'A' && b <= 'Z' {
			result[i] = b - 'A' + 'a'
		} else {
			result[i] = b
		}
	}
	return result
}

// MD5 calculates a message-digest fingerprint (checksum) for a file.
func hashMD5(text []byte) []byte {
	hash := md5.Sum(text)
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash[:])
	return dst
}

// SHA1 message-digest algorithm. This algorithm takes a message and generates a
// 160-bit digest from the input. The SHA1 algorithm is related to the MD4
// algorithm but has been strengthend against certain types of cryptographic
// attack. SHA1 should be used in preference to MD4 or MD5 in new applications.
func hashSHA1(text []byte) []byte {
	hash := sha1.Sum(text)
	dst := make([]byte, hex.EncodedLen(len(hash)))
	hex.Encode(dst, hash[:])
	return dst
}

// chunk splits a string into smaller pieces, default: 64.
func chunk(text []byte, number string) []byte {
	limit, err := strconv.Atoi(number)
	if err != nil || limit <= 0 {
		limit = 64
	}
	total := len(text)
	newlineCount := total / limit
	result := make([]byte, 0, total+newlineCount)
	for i := 0; i < total; i += limit {
		end := i + limit
		if end > total {
			end = total
		}
		result = append(result, text[i:end]...)
		if i+limit <= total {
			result = append(result, '\n')
		}
	}
	return result
}

// length measures the arbitrary (but finite) length of a chain of letters.
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
func length(text []byte, verbose string) []byte {
	if verbose == "-v" {
		buf := make([]byte, 0, len(text)+16)
		buf = append(buf, '[')
		buf = strconv.AppendInt(buf, int64(len(text)), 10)
		buf = append(buf, ']', 's', 't', 'r', 'i', 'n', 'g', '{')
		buf = strconv.AppendQuote(buf, string(text))
		buf = append(buf, '}')
		return buf
	}
	return []byte(strconv.Itoa(len(text)))
}

// base64Encode encodes data with MIME base64. Base64 is a group of similar
// binary-to-text encoding schemes that represent binary data in an ASCII string
// format by translating it into a radix-64 representation. The term Base64
// originates from a specific MIME content transfer encoding. Each base64 digit
// represents exactly 6 bits of data.
func base64Encode(text []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(text)))
	base64.StdEncoding.Encode(dst, text)
	return dst
}

// base64Decode decodes data encoded with MIME base64. Base64 is a group of
// similar binary-to-text encoding schemes that represent binary data in an
// ASCII string format by translating it into a radix-64 representation. The
// term Base64 originates from a specific MIME content transfer encoding. Each
// base64 digit represents exactly 6 bits of data.
func base64Decode(text []byte) []byte {
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(text)))
	n, err := base64.StdEncoding.Decode(dst, text)
	if err != nil {
		return []byte(err.Error())
	}
	return dst[:n]
}

// urlEncode encodes a chain of letters using the percent-encoding technique.
// Percent-encoding, also known as URL encoding, is a mechanism for encoding
// information in a Uniform Resource Identifier (URI) under certain
// circumstances. Although it is known as URL encoding it is, in fact, used more
// generally within the main Uniform Resource Identifier (URI) set, which
// includes both Uniform Resource Locator (URL) and Uniform Resource Name (URN).
// As such, it is also used in the preparation of data of the application/x-www-
// form-urlencoded media type, as is often used in the submission of HTML form
// data in HTTP requests.
func urlEncode(text []byte) []byte {
	if text[len(text)-1] == '\n' {
		text = text[0 : len(text)-1]
	}
	return []byte(url.QueryEscape(string(text)))
}

// urlDecode decodes a chain of letters using the percent-encoding technique.
// Percent-encoding, also known as URL encoding, is a mechanism for encoding
// information in a Uniform Resource Identifier (URI) under certain
// circumstances. Although it is known as URL encoding it is, in fact, used more
// generally within the main Uniform Resource Identifier (URI) set, which
// includes both Uniform Resource Locator (URL) and Uniform Resource Name (URN).
// As such, it is also used in the preparation of data of the application/x-www-
// form-urlencoded media type, as is often used in the submission of HTML form
// data in HTTP requests.
func urlDecode(text []byte) []byte {
	out, err := url.QueryUnescape(string(text))
	if err != nil {
		return []byte(err.Error())
	}
	return []byte(out)
}

// rotate performs the rot13 transform on a string. The ROT13 encoding simply
// shifts every letter by 13 places in the alphabet while leaving non-alpha
// characters untouched. Encoding and decoding are done by the same function,
// passing an encoded string as argument will return the original version.
func rotate(text []byte, number string) []byte {
	n, err := strconv.Atoi(number)
	if err != nil {
		n = 13
	}
	shift := byte((n%26 + 26) % 26)
	if shift == 0 {
		return append([]byte(nil), text...)
	}
	result := make([]byte, len(text))
	for i, b := range text {
		switch {
		case b >= 'A' && b <= 'Z':
			result[i] = 'A' + (b-'A'+shift)%26
		case b >= 'a' && b <= 'z':
			result[i] = 'a' + (b-'a'+shift)%26
		default:
			result[i] = b
		}
	}
	return result
}
