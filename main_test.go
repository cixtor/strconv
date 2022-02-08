package main

import (
	"bytes"
	"testing"
)

func TestReplace(t *testing.T) {
	table := map[string][]string{
		"basic": {
			"tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2",
			"n",
			"@",
			"tf GO1Q@@7S60 WqsxKdVFQ7 @@@FQ@sVD1@5MPpVqB@NrraQd2",
		},
		"n_to_at": {
			"tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2",
			"n",
			"@",
			"tf GO1Q@@7S60 WqsxKdVFQ7 @@@FQ@sVD1@5MPpVqB@NrraQd2",
		},
		"n_to_empty": {
			"tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2",
			"n",
			"",
			"tf GO1Q7S60 WqsxKdVFQ7 FQsVD15MPpVqBNrraQd2",
		},
		"n_to_equal": {
			"tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2",
			"n",
			"=",
			"tf GO1Q==7S60 WqsxKdVFQ7 ===FQ=sVD1=5MPpVqB=NrraQd2",
		},
		"colon_to_semicolon": {
			"L:r:m :ps:m d:l:r s:t :m:t, c:ns:ct:t:r :d:p:s:c:ng :l:t",
			":",
			";",
			"L;r;m ;ps;m d;l;r s;t ;m;t, c;ns;ct;t;r ;d;p;s;c;ng ;l;t",
		},
		"new_line1": {
			"hello\nworld",
			"\n",
			"@",
			"hello@world",
		},
		"new_line2": {
			"hello@world",
			"@",
			"\n",
			"hello\nworld",
		},
		"tabulation1": {
			"hello\tworld",
			"\t",
			"@",
			"hello@world",
		},
		"tabulation2": {
			"hello@world",
			"@",
			"\t",
			"hello\tworld",
		},
		"world_to_rabbit": {
			"hello world hello world hello world hello world",
			"world",
			"rabbit",
			"hello rabbit hello rabbit hello rabbit hello rabbit",
		},
	}
	for name, v := range table {
		t.Run(name, func(t *testing.T) {
			expected := []byte(v[3])
			actual := replace([]byte(v[0]), v[1], v[2])
			if !bytes.Equal(expected, actual) {
				t.Fatalf("incorrect replacement\n- %s\n+ %s", expected, actual)
			}
		})
	}
}

type TableTest struct {
	handler  func([]byte) []byte
	expected string
}

func TestOtherFunctions(t *testing.T) {
	table := map[string]TableTest{
		"capitalize": {handler: capitalize, expected: "Tf GO1Qnn7S60 WqsxKdVFQ7 NnnFQnsVD1n5MPpVqBnNrraQd2"},
		"uppercase":  {handler: uppercase, expected: "TF GO1QNN7S60 WQSXKDVFQ7 NNNFQNSVD1N5MPPVQBNNRRAQD2"},
		"lowercase":  {handler: lowercase, expected: "tf go1qnn7s60 wqsxkdvfq7 nnnfqnsvd1n5mppvqbnnrraqd2"},
		"md5":        {handler: hashMD5, expected: "03e32e975dc9565535b854908a0e8624"},
		"sha1":       {handler: hashSHA1, expected: "c9e1d6ee08eb9f0355a4c2ee58cc37c536e57182"},
	}
	for name, v := range table {
		t.Run(name, func(t *testing.T) {
			expected := []byte(v.expected)
			actual := v.handler([]byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"))
			if !bytes.Equal(expected, actual) {
				t.Fatalf("incorrect replacement\n- %s\n+ %s", expected, actual)
			}
		})
	}
}

func TestLength(t *testing.T) {
	expected := []byte("51")
	actual := length([]byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"), "")
	if !bytes.Equal(actual, expected) {
		t.Fatalf("length did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestBase64Encode(t *testing.T) {
	expected := []byte("dGYgR08xUW5uN1M2MCBXcXN4S2RWRlE3IG5ubkZRbnNWRDFuNU1QcFZxQm5OcnJhUWQy")
	actual := base64Encode([]byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"))
	if !bytes.Equal(actual, expected) {
		t.Fatalf("base64Encode did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestBase64Decode(t *testing.T) {
	expected := []byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2")
	actual := base64Decode([]byte("dGYgR08xUW5uN1M2MCBXcXN4S2RWRlE3IG5ubkZRbnNWRDFuNU1QcFZxQm5OcnJhUWQy"))
	if !bytes.Equal(actual, expected) {
		t.Fatalf("base64Decode did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestURLDecode(t *testing.T) {
	expected := []byte("https://cixtor.com/?foo=bar&lorem=ipsum")
	actual := urlDecode([]byte("https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum"))
	if !bytes.Equal(actual, expected) {
		t.Fatalf("urlDecode did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestURLEncode(t *testing.T) {
	expected := []byte("https%3A%2F%2Fcixtor.com%2F%3Ffoo%3Dbar%26lorem%3Dipsum")
	actual := urlEncode([]byte("https://cixtor.com/?foo=bar&lorem=ipsum"))
	if !bytes.Equal(actual, expected) {
		t.Fatalf("urlEncode did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestRotate13(t *testing.T) {
	expected := []byte("gs TB1Daa7F60 JdfkXqISD7 aaaSDafIQ1a5ZCcIdOaAeenDq2")
	actual := rotate([]byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"), "13")
	if !bytes.Equal(actual, expected) {
		t.Fatalf("rotate did not run as expected\n- %q\n+ %q", expected, actual)
	}
}

func TestRotate5(t *testing.T) {
	expected := []byte("yk LT1Vss7X60 BvxcPiAKV7 sssKVsxAI1s5RUuAvGsSwwfVi2")
	actual := rotate([]byte("tf GO1Qnn7S60 WqsxKdVFQ7 nnnFQnsVD1n5MPpVqBnNrraQd2"), "5")
	if !bytes.Equal(actual, expected) {
		t.Fatalf("rotate did not run as expected\n- %q\n+ %q", expected, actual)
	}
}
