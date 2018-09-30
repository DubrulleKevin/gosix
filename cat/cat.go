package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/DubrulleKevin/gosix/common"
)

func printLines(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func printStdin() {
	printLines(os.Stdin)
}

func printFile(fileName string) {
	fileInfo, err := os.Stat(fileName)
	common.PanicIfError(err)
	if !fileInfo.IsDir() {
		file, err := os.Open(fileName)
		common.PanicIfError(err)
		defer file.Close()

		printLines(file)
	}
}

func main() {
	flag.Bool("u", false, "Not implemented")
	flag.Parse()
	args := flag.Args()

	if len(args) == 0 {
		printStdin()
	} else {
		for _, fileName := range args {
			if fileName == "-" {
				printStdin()
			} else {
				printFile(fileName)
			}
		}
	}
}
