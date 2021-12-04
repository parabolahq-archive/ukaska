package main

import (
	"os"
	"time"
	"ukaska/backup"
	"ukaska/config"
	"ukaska/telegram"
)

func main() {
	config.Init()
	now := time.Now().Format(time.RFC1123)
	backupFiles := backup.BackupMongo()
	nonNullCollections := []string{}
	var fileName string
	for i := 0; i < len(backupFiles); i++ {
		if backupFiles[i] != "" {
			fileName = backupFiles[i]
			file, err := os.Open(fileName)
			backup.HandleIfError(err)
			stat, _ := file.Stat()
			if stat.Size() > 0 {
				nonNullCollections = append(nonNullCollections, backupFiles[i])
			}
			err = file.Close()
			backup.HandleIfError(err)
		}
	}
	if len(nonNullCollections) > 0 {
		telegram.SendMessage(now)
		for i := 0; i < len(nonNullCollections); i++ {
			telegram.SendDocument(nonNullCollections[i])
		}
	}

}
