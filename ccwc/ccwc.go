package main

import (
	"os"
	"bufio"
	"fmt"
)


func main() {

	args := os.Args

	cmdOption := args[1]
	fileName := args[2]

	// read the file
	f, err := os.Open(fileName)

	if err != nil {
		println(err)
		return
	}
	
	defer f.Close()

	if cmdOption == "-c" {

		// scan the content
		byteCounter := 0
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanBytes)
		for scanner.Scan() {
			b := scanner.Bytes()
			if len(b) != 0 { byteCounter += 1 }
		}

		fmt.Printf("Bytes : %d", byteCounter)

	} else if cmdOption == "-l" {
		
		lineCount := 0
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanLines)
		for scanner.Scan() {
			lineCount += 1
		}

		fmt.Printf("Lines : %d", lineCount)
	} else if cmdOption == "-w" {
		
		wordCount := 0
		scanner := bufio.NewScanner(f)
		scanner.Split(bufio.ScanWords)
		for scanner.Scan() {
			wordCount += 1
		}
		fmt.Printf("Words : %d", wordCount)
	}

	
}
