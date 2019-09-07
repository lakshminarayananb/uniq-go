package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main(){
	fmt.Println("uniq utility in Golang")
	content, err := ioutil.ReadFile("testdata/sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", content)
	fmt.Println()
}