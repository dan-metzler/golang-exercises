package main

import (
	"fmt"
	"io"
	"os"
)

// write to a text file
func main() {

	// fmt.Println prints a line to the terminal in Go
	fmt.Println("This is a printed message in golang using fmt.Println() method")

	// This is the content that will be placed in the text file below
	fileContent := "Lorem ipsum dolor sit amet consectetur adipisicing elit. Maxime deserunt vero hic quaerat consequuntur, omnis alias sint tempora eos, ut eius illum enim quidem ratione dolor! Doloribus a dolores corporis aut! Fugit, non saepe quod nam, commodi ex laborum consectetur fuga quidem autem iusto blanditiis quaerat esse? Exercitationem, consequatur dicta!"

	textFileRelPath := "./Text.txt"

	// crate the text.txt file to write the file content to
	file, err := os.Create(textFileRelPath)

	checkNilErr(err)

	// write the content to the created text file
	length, err := io.WriteString(file, fileContent)

	checkNilErr(err)

	fmt.Println("File Content Length Is", length, "Chars")

	// close tht text file after use
	defer file.Close()

	readFile(textFileRelPath)

}

// read from a text file
func readFile(filename string) {

	// the file content is returned as a databyte
	databyte, err := os.ReadFile(filename)

	checkNilErr(err)

	// this databyte needs to be converted to a string before being output to the terminal
	fmt.Println("File Content:", string(databyte))

}

// error checking function. If an unexpected conditon arises in our program panic will terminate the program
func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
