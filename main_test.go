package main

import "testing"

const text = "tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"

func TestReplace(t *testing.T) {
	var app Application

	app.args = []string{"n", "@"}

	if app.Replace(text) != "tf GO1Q@@7S60 WqsxKdVFQ7 @@@FQ@sVD1@5MPpVqB@NrraQd2" {
		t.Fatalf("Replace did not run as expected")
	}
}

func TestCapitalize(t *testing.T) {
	var app Application

	if app.Capitalize(text) != "Tf GO1Qnn7S60 WqsxKdVFQ7 NnnFQnsVD1n5MPpVqBnNrraQd2" {
		t.Fatalf("Capitalize did not run as expected")
	}
}

func TestUppercase(t *testing.T) {
	var app Application

	if app.Uppercase(text) != "TF GO1QNN7S60 WQSXKDVFQ7 NNNFQNSVD1N5MPPVQBNNRRAQD2" {
		t.Fatalf("Uppercase did not run as expected")
	}
}

func TestLowercase(t *testing.T) {
	var app Application

	if app.Lowercase(text) != "tf go1qnn7s60 wqsxkdvfq7 nnnfqnsvd1n5mppvqbnnrraqd2" {
		t.Fatalf("Lowercase did not run as expected")
	}
}

func TestMd5(t *testing.T) {
	var app Application

	if app.Md5(text) != "03e32e975dc9565535b854908a0e8624" {
		t.Fatalf("Md5 did not run as expected: %s", app.Md5(text))
	}
}

func TestSha1(t *testing.T) {
	var app Application

	if app.Sha1(text) != "c9e1d6ee08eb9f0355a4c2ee58cc37c536e57182" {
		t.Fatalf("Sha1 did not run as expected")
	}
}

func TestLength(t *testing.T) {
	var app Application

	if app.Length(text) != "51" {
		t.Fatalf("Length did not run as expected")
	}
}

func TestBase64Encode(t *testing.T) {
	var app Application

	expected := "dGYgR08xUW5uN1M2MCBXcXN4S2RWRlE3IG5ubkZRbnNWRDFuNU1QcFZxQm5OcnJhUWQy"

	if app.Base64Encode(text) != expected {
		t.Fatalf("Base64Encode did not run as expected")
	}
}

func TestBase64Decode(t *testing.T) {
	var app Application

	initial := "dGYgR08xUW5uN1M2MCBXcXN4S2RWRlE3IG5ubkZRbnNWRDFuNU1QcFZxQm5OcnJhUWQy"

	if app.Base64Decode(initial) != text {
		t.Fatalf("Base64Decode did not run as expected")
	}
}

func TestURLDecode(t *testing.T) {
	var app Application

	initial := "https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum"
	expected := "https://cixtor.com/?foo=bar&lorem=ipsum"

	if app.URLDecode(initial) != expected {
		t.Fatalf("URLDecode did not run as expected")
	}
}

func TestURLEncode(t *testing.T) {
	var app Application

	initial := "https://cixtor.com/?foo=bar&lorem=ipsum"
	expected := "https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum"

	if app.URLEncode(initial) != expected {
		t.Fatalf("URLEncode did not run as expected")
	}
}

func TestRotate(t *testing.T) {
	var app Application

	if app.Rotate(text) != "gs TB1Daa7F60 JdfkXqISD7 aaaSDafIQ1a5ZCcIdOaAeenDq2" {
		t.Fatalf("Rotate did not run as expected")
	}
}
