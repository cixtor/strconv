package main

import "testing"

const text = "tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"

func TestReplace(t *testing.T) {
	var app Application

	if app.Replace(text, "n", "@") != "tf GO1Q@@7S60 WqsxKdVFQ7 @@@FQ@sVD1@5MPpVqB@NrraQd2" {
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

	if app.Length(text) != 51 {
		t.Fatalf("Length did not run as expected")
	}
}
