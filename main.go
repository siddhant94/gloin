package main

import (
	errLog "bing-wod/logging"
	"bing-wod/utils"
	"fmt"
	//"log"
	"os"
)

func main() {
	// argsWithProg := os.Args
	userDirPath := os.Args[1] // Get arguments aside from pragram path
	fmt.Println("If the directory mentioned is not present, this will create one")
	check, err := utils.Exists(userDirPath) // Checks whether file/dir exists or not
	if err != nil {
		//log.Fatalf("error creating/opening log file: %v", err)
		errLog.Logger.Printf("error creating/opening log file: %v \n", err)
		return
	}
	if !check {
		err := os.Mkdir(userDirPath, 0755)
		if err != nil {
			//log.Fatalf("error creating directory: %v", err)
			errLog.Logger.Printf("error creating directory: %v \n", err)
			return
		}
	}
}
