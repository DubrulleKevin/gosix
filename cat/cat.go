package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func panicIfError(e error) {
	if e != nil {
		panic(e)
	}
}

func stringSliceContains(s []string, e string) bool {
	for _, el := range s {
		if el == e {
			return true
		}
	}

	return false
}

func main() {
	flag.Bool("u", false, "Not implemented")

	flag.Parse()

	if len(flag.Args()) == 0 || stringSliceContains(flag.Args(), "-") {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		for _, fileName := range flag.Args() {
			fileInfo, err := os.Stat(fileName)
			panicIfError(err)
			if !fileInfo.IsDir() {
				fileContent, err := ioutil.ReadFile(fileName)
				panicIfError(err)
				fmt.Print(string(fileContent))
			}
		}
	}
}
