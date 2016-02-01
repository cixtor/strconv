package main

import "testing"

const text = "tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"

func TestReplace(t *testing.T) {
	var app Application

	if app.Replace(text, "n", "@") != "tf GO1Q@@7S60 WqsxKdVFQ7 @@@FQ@sVD1@5MPpVqB@NrraQd2" {
		t.Fatalf("Replace did not run as expected")
	}
}
