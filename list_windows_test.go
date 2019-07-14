package main

import "testing"

func TestOutputFiles(t *testing.T) {
	out, err := outputFiles("files")
	if err != nil {
		t.Error(err)
		return
	}

	expects := "a.txt\r\nb.txt\r\n\r\nc.txt\r\nd.txt" // failed

	if expects != out {
		t.Errorf("want: %s\r\ngot: %s", expects, out)
	}
}
