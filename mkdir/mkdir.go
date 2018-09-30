package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/DubrulleKevin/gosix/common"
)

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
		fmt.Println(fileMode)
	} else {
		fileMode = 0777
	}

	for _, arg := range args {
		var err error
		if *pFlagPtr {
			err = os.MkdirAll(arg, fileMode)
		} else {
			err = os.Mkdir(arg, fileMode)
		}
		common.PanicIfError(err)
	}
}
