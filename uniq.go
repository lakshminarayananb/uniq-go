package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strings"
)

func main(){
	fmt.Println("uniq utility in Golang")

	// Open file
	file, err := os.Open("testdata/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string

	// Read opened file using Scanner
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		lineText := scanner.Text()
		lines = append(lines,lineText)
	}

	count := 0
	var newLines []string

	// Loop through each line and check if previous line is same as current line
	// If not, add to a new slice, else don't
	// Obviously not efficient
	for count < len(lines){
		if count > 0 {
			if lines[count] != lines[count-1] {
				newLines = append(newLines, lines[count])
			}
		} else {
			newLines = append(newLines, lines[count])
		}
		count += 1		
	}

	// Create an output file
	fileo, err := os.Create("testdata/output.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fileo.Close()

	// Write the generated slice to file as string separated by newlines
	fileo.WriteString(strings.Join(newLines,"\n"))

	fmt.Println("output saved as output.txt")
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}