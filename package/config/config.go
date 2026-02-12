package config

import (
	"library/package/log"
	"fmt"
	"os"
)

func Downloads() string {
	path, err := os.Getwd()
	if err != nil {
		log.Errorf("error getting working directory %v:", err)
		return ""
	}
	path = fmt.Sprintf("%s/.storage/downloads/", path)
	err = CreateFolderIfDoesntExist(path)
	if err != nil {
		log.Error("error getting downloads directory: %v", err)
		return ""
	}
	return path
}

func CreateFolderIfDoesntExist(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
