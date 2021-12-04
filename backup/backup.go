package backup

import (
	"fmt"
	"log"
	"os/exec"
	"ukaska/config"
)

func BackupMongo() (backupFilesNames []string) {
	for i := 0; i < len(config.CollectionNames); i++ {
		collectionName := config.CollectionNames[i]
		command, params := getExecutableCommand(collectionName)
		cmd := exec.Command(command, params...)
		log.Println(command, params)
		output, err := cmd.Output()
		HandleIfError(err)
		log.Println(output)
		backupFilesNames = append(backupFilesNames, fmt.Sprintf("%s.bak", collectionName))
	}
	return
}

func getExecutableCommand(collectionName string) (command string, params []string) {
	command = "/usr/bin/mongoexport"
	params = []string{"--forceTableScan",
		fmt.Sprintf("--uri=\"%s\"", config.DbUrl),
		fmt.Sprintf("-c=\"%s\"", collectionName),
		fmt.Sprintf("-o=%s.bak", collectionName)}
	return
}

func HandleIfError(err error) {
	if err != nil {
		panic(err)
	}
}
