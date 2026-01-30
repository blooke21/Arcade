package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func buildMoveFile(sourcePath string) (map[string]string, error) {
	var fileMap map[string]string = make(map[string]string);

	ext := filepath.Ext(sourcePath)

	cleanedSourcePath := filepath.Clean(sourcePath)

	index, err := getAllRomDatabaseLength()
	if err != nil {
		return nil, err
	}

	fileMap["id"] = fmt.Sprint(index);

	//set default destination folder to "Other" for unrecognized file types
	fileMap["type"] = "Other"

	fileMap["fileName"] = strings.TrimSuffix(filepath.Base(sourcePath), ext)

	switch expression := ext; expression { //add more cases as needed for different ROM types
	case ".gba":
		fileMap["type"] = "gba"
	}
	
	// Create new file path for rom
	destinationFolder := filepath.Join("../rom", fileMap["type"])
	destinationPath := filepath.Join(destinationFolder, fileMap["fileName"] + ext)
	fileMap["source"] = destinationPath

	// Check if file already exists
    if _, err := os.Stat(destinationPath); err == nil {
        return nil, ErrDuplicateROM // Use your custom error here
    }

	moveFile(cleanedSourcePath, destinationFolder, destinationPath)

	fileMap["image"] = romImgs
	return fileMap, nil;
}