package backup

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"ukaska/config"
)

type BackedUpFile struct {
	Filename       string
	DocumentsCount int
	CollectionName string
}

func BackupMongo() (backupFilesNames []BackedUpFile) {
	for i := 0; i < len(config.CollectionNames); i++ {
		collectionName := config.CollectionNames[i]
		command, params := getExecutableCommand(collectionName)
		cmd := exec.Command(command, params...)
		log.Println(command, params)
		output, err := cmd.Output()
		HandleIfError(err)
		log.Println(output)
		filename := fmt.Sprintf("%s.bak", collectionName)
		stat, _ := os.Stat(filename)
		file, _ := os.Open(filename)
		lines, _ := lineCounter(file)
		file.Close()
		if stat.Size() > 0 {
			backupFilesNames = append(backupFilesNames, BackedUpFile{
				Filename:       filename,
				DocumentsCount: lines,
				CollectionName: collectionName,
			})
		}
	}
	return
}

// https://stackoverflow.com/questions/24562942/golang-how-do-i-determine-the-number-of-lines-in-a-file-efficiently
func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
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
