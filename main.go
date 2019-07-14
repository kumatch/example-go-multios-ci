package main

import "fmt"

func main() {
	out, err := outputFiles("files")
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
