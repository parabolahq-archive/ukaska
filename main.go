package main

import (
	"fmt"
	"time"
	"ukaska/backup"
	"ukaska/config"
	"ukaska/telegram"
)

func main() {
	config.Init()
	now := time.Now().Format(time.RFC1123)
	backupFiles := backup.BackupMongo()
	if len(backupFiles) > 0 {
		telegram.SendMessage(now)
		for i := 0; i < len(backupFiles); i++ {
			telegram.SendMessage(fmt.Sprintf("%s(%d)", backupFiles[i].CollectionName, backupFiles[i].DocumentsCount))
			telegram.SendDocument(backupFiles[i].Filename)
		}
	}
}
