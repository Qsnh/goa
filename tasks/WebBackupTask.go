package tasks

import (
	"archive/zip"
	"github.com/Qsnh/goa/utils"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var BACKUP_EXCLUDE  = map[string]bool{
	".git": true,
	"node_modules": true,
}

func AllFiles(dist string) []string {
	var files []string
	files, err := filepath.Glob(dist + string(os.PathSeparator) + "*")
	if err != nil {
		return files
	}
	for _, path := range files {
		baseName := filepath.Base(path)
		if BACKUP_EXCLUDE[baseName] == true {
			continue
		}
		fileInfo, err := os.Stat(path)
		if err != nil {
			continue
		}
		if fileInfo.IsDir() {
			children := AllFiles(path)
			files = append(files, path)
			files = append(files, children...)
			continue
		}
		files = append(files, path)
	}
	return files
}

func Backup(dist string, files []string) error {
	backupFile, err := os.Create(dist)
	if err != nil {
		return err
	}
	zipWriter := zip.NewWriter(backupFile)
	defer zipWriter.Close()
	for _, file := range files {
		fileInfo, err := os.Stat(file)
		if err != nil {
			continue
		}
		fileHeader, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			continue
		}
		if fileInfo.IsDir() {
			fileHeader.Name = strings.TrimLeft(file, "/") + "/"
			fileHeader.Method = zip.Store
		} else {
			fileHeader.Name = strings.TrimLeft(file, "/")
			fileHeader.Method = zip.Deflate
		}
		zipWriter, err := zipWriter.CreateHeader(fileHeader)
		if err != nil {
			continue
		}
		if fileInfo.IsDir() {
			continue
		}
		// 写入文件数据
		fileContent, err := ioutil.ReadFile(file)
		if err != nil {
			continue
		}
		zipWriter.Write(fileContent)
	}
	return nil
}

func WebBackupTask() error {
	dist := strings.TrimRight(os.Getenv("BACKUP_SAVE_PATH"), "/") + "/" + time.Now().Format("2006-01-01 15-04-05") + "_backup.zip"
	pwd := utils.Pwd()
	return Backup(dist, AllFiles(pwd))
}