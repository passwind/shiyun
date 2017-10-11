package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
)

func main() {
	b, err := ioutil.ReadFile("shiyun.txt") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	// fmt.Println(b) // print the content as 'bytes'

	// str := string(b) // convert content to a 'string'

	// fmt.Println(str) // print the content as a 'string'

	n, lines, err := bufio.ScanLines(b, true)
	if err != nil {
		fmt.Printf("error %s\n", err)
	}
	fmt.Println(n)
	fmt.Println(len(lines))
	fmt.Println(string(lines[0]))
}
