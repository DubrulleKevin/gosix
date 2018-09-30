package main

import (
	"bufio"
	"flag"
	"fmt"
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
		file, err := os.Open(fileName)
		common.PanicIfError(err)
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
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
