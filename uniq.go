package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"runtime"
)

func main(){
	fmt.Println("uniq utility in Golang")
	fmt.Println("Runtime " + runtime.GOOS)
	file, err := os.Open("testdata/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
		//fmt.Println()
	}
	
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}