package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DubrulleKevin/gosix/common"
)

func readFromStdin() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func readFromFile(fileName string) {
	fileInfo, err := os.Stat(fileName)
	common.PanicIfError(err)
	if !fileInfo.IsDir() {
		fileContent, err := ioutil.ReadFile(fileName)
		common.PanicIfError(err)
		fmt.Print(string(fileContent))
	}
}

func main() {
	flag.Bool("u", false, "Not implemented")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		readFromStdin()
	} else {
		for _, fileName := range args {
			if fileName == "-" {
				readFromStdin()
			} else {
				readFromFile(fileName)
			}
		}
	}
}
