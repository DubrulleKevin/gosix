package main

import (
	"flag"
	"os"
	"strconv"
	"strings"

	"github.com/DubrulleKevin/gosix/common"
)

func chmodAll(path string, mode os.FileMode) {
	var workingPath string

	for i, el := range strings.Split(path, "/") {
		if i > 0 {
			workingPath += "/"
		}
		workingPath += el
		err := os.Chmod(workingPath, mode)
		common.PanicIfError(err)
	}
}

func main() {
	pFlagPtr := flag.Bool("p", false, "Create any missing intermediate pathname components.")
	mFlagPtr := flag.String("m", "", "Set the file permission bits of the newly-created directory to the specified mode value.")
	flag.Parse()
	args := flag.Args()

	var fileMode os.FileMode

	if *mFlagPtr != "" {
		fileMode64, err := strconv.ParseUint(*mFlagPtr, 8, 32)
		fileMode = os.FileMode(fileMode64)
		common.PanicIfError(err)
	}

	for _, arg := range args {
		var err error
		if *pFlagPtr {
			err = os.MkdirAll(arg, 0777)
			if fileMode != 0 {
				chmodAll(arg, fileMode)
			}
		} else {
			err = os.Mkdir(arg, 0777)
			if fileMode != 0 {
				os.Chmod(arg, fileMode)
			}
		}
		common.PanicIfError(err)
	}
}
