package main

import (
	errLog "bing-wod/logging"
	"bing-wod/utils"
	"fmt"
	//"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Directory name missing")
		return
	}
	userDirPath := os.Args[1] // Get arguments aside from pragram path
	fmt.Println("If the directory mentioned is not present, this will create one. Here all wallpapers will be" +
		"downloaded in future...:)")
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
	// We have directory path, now time for calling rest endpoint and gfetting wallpaper meta
	metaData := utils.GetBingWallpaperMeta()
	success, savedFile := utils.DownloadWallpaper(metaData, userDirPath)
	if(success) {
		fmt.Println("Finished ...Your wallpaper is downloaded and with filename "+ savedFile +" and is present at " + userDirPath)
	} else {
		fmt.Println("Uh oh!! Could not save Image file. Check error log")
	}
}
