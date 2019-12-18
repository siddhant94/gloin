package linux

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

//crontab -l > my-crontab
const userCrontabPerm = 0600
const binaryName = "gloin"
const gloinLog = "gloinCron.log"

func AddCronEntry() {
	fmt.Println("Adding CRON")
	// Get User Crontab entries
	crontabContent, err := exec.Command("crontab", "-l").Output()
	if err != nil {
		log.Printf("Error executing command : %v \n", err)
	}
	//gloinJob := "* * * * * /home/sid/Desktop/Workspace/gloin/gloin >> /home/sid/Desktop/Workspace/gloin/crontab.log\n"
	cronCmd, err := buildCronCmd()
	if err != nil {
		return
	}
	crontabContent = append(crontabContent, cronCmd...) // use "..."
	fmt.Println(string(crontabContent))
}

func buildCronCmd() ([]byte, error) {
	currDir, err := os.Getwd()
	if err != nil {
		log.Printf("Error getting working directory : %v \n", err)
		return []byte{}, err
	}
	//TODO : Make this configurable
	frequency := "@daily"// "* * * * *"
	fullBinaryPath := currDir + "/" + binaryName
	gloinCronLog := currDir + "/" + gloinLog
	jobCmd := []byte(frequency + " " + fullBinaryPath + " >> " + gloinCronLog + "\n")
	return jobCmd, nil
}
