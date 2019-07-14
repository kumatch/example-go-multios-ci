// +build !windows

package main

import "testing"

func TestOutputFiles(t *testing.T) {
	out, err := outputFiles("files")
	if err != nil {
		t.Error(err)
		return
	}

	expects := `.
..
a.txt
b.txt
c.txt`

	if expects != out {
		t.Errorf("want: %s\ngot: %s", expects, out)
	}
}
