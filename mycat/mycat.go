package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var n bool

func init() {
	flag.BoolVar(&n, "n", false, "行数を出すかどうか")
}

func main() {
	flag.Parse()
	fileNames := flag.Args()
	if len(fileNames) == 0 {
		log.Fatal("ファイル名を入れてください")
	}

	line := 1
	for _, fileName := range fileNames {
		endLine := printContent(fileName, line)
		line = endLine + 1
	}
}

func printContent(fileName string, startLine int) (endLine int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("ファイルを開けませんでした。入力されたファイル名:%s", fileName)
	}
	defer file.Close()

	line := startLine
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if n {
			fmt.Printf("%d: %s\n", line, scanner.Text())
		} else {
			fmt.Println(scanner.Text())
		}
		line++
	}

	return line
}
