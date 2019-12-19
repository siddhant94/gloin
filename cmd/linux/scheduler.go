package linux

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

//crontab -l > my-crontab
const userCrontabPerm = 0600
const binaryName = "gloin"
const gloinLog = "gloinCron.log"

func AddCronEntry(freq string) {
	fmt.Println("Adding CRON")
	// Get User Crontab entries
	crontabContent, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		log.Printf("Error executing command : %v \n", err)
	}
	//gloinJob := "* * * * * /home/sid/Desktop/Workspace/gloin/gloin >> /home/sid/Desktop/Workspace/gloin/crontab.log\n"
	cronCmd, err := buildCronCmd(freq)
	if err != nil {
		return
	}
	crontabContent = append(crontabContent, cronCmd...) // use "..."
	// Store in temp file
	file, err := ioutil.TempFile("", "temp_cron")
	if err != nil {
		log.Printf("Error in creating temp file for saving cron : %v", err)
		return
	}
	defer os.Remove(file.Name())
	// Write cron content to temp file
	_ , err = file.Write(crontabContent)
	if err != nil {
		log.Printf("Error writing to temp file : %v\n", err)
		return
	}
	_, err = exec.Command("crontab", file.Name()).Output()
	if err != nil {
		log.Printf("Error executing cmd : %v", err)
		return
	}
	fmt.Println("Successfully saved new cron")
}

func buildCronCmd(freq string) ([]byte, error) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory : %v \n", err)
		return []byte{}, err
	}
	//TODO : Make this configurable
	if len(freq) < 5 {
		freq = "* * * * *" // Malformed, fallback to default
	}
	fullBinaryPath := currDir + "/" + binaryName
	gloinCronLog := currDir + "/" + gloinLog
	jobCmd := []byte(freq + " " + fullBinaryPath + " >> " + gloinCronLog + "\n")
	return jobCmd, nil
}