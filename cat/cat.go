package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/DubrulleKevin/gosix/common"
)

func main() {
	flag.Bool("u", false, "Not implemented")

	flag.Parse()

	if len(flag.Args()) == 0 || common.StringSliceContains(flag.Args(), "-") {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	} else {
		for _, fileName := range flag.Args() {
			fileInfo, err := os.Stat(fileName)
			common.PanicIfError(err)
			if !fileInfo.IsDir() {
				fileContent, err := ioutil.ReadFile(fileName)
				common.PanicIfError(err)
				fmt.Print(string(fileContent))
			}
		}
	}
}
