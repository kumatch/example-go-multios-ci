package main

import "testing"

func TestOutputFilesOnWindows(t *testing.T) {
	out, err := outputFiles("files")
	if err != nil {
		t.Error(err)
		return
	}

	expects := "a.txt\r\nb.txt\r\nc.txt"

	if expects != out {
		t.Errorf("want: %s\r\ngot: %s", expects, out)
	}
}
