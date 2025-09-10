package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)


func handleMoveFile(sourcePath string) (map[string]string, error) {
	var fileMap map[string]string = make(map[string]string);

	ext := filepath.Ext(sourcePath)

	cleanedSourcePath := filepath.Clean(sourcePath)

	//set default destination folder to "Other" for unrecognized file types
	fileMap["type"] = "Other"

	fileMap["fileName"] = strings.TrimSuffix(filepath.Base(sourcePath), ext)

	switch expression := ext; expression { //add more cases as needed for different ROM types
	case ".gba":
		fileMap["type"] = "gba"
	}
	
	destinationFolder := filepath.Join("../rom", fileMap["type"])
	destinationPath := filepath.Join(destinationFolder, fileMap["fileName"] + ext)
	fileMap["source"] = destinationPath

	// Check if destination file already exists
    if _, err := os.Stat(destinationPath); err == nil {
        return nil, ErrDuplicateROM // Use your custom error here
    }

	moveFile(cleanedSourcePath, destinationFolder, destinationPath)
	fileMap["image"] = "defualt_image.png" // Placeholder for image path, can be updated later
	return fileMap, nil;
}

func moveFile(sourcePath, destinationFolder, destinationPath string) error {
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("source file does not exist: %s", sourcePath)
	}

	if err := os.MkdirAll(destinationFolder, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %v", err)
	}

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		return fmt.Errorf("failed to move file: %v", err)
	}

	return nil;
}