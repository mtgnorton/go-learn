package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func generate() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func writeFile(filename string) {
	intGen := generate()
	if file, err := os.Create(filename); err != nil {
		fmt.Println(err)
	} else {
		defer file.Close()
		writer := bufio.NewWriter(file)

		defer writer.Flush()

		for i := 0; i < 20; i++ {
			fmt.Fprintln(writer, intGen())
		}
	}

}

func main() {
	f := generate()
	printFileContents(f)
	writeFile("abc.txt")
}
