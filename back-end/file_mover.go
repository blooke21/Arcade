package main

import (
	"fmt"
	"os"
	"path/filepath"
)


func handleMoveFile(sourcePath string) error {
	romFolder := filepath.Join("..", "rom")

	ext := filepath.Ext(sourcePath)

	cleanPath := filepath.Clean(sourcePath)

	destinationFolder := filepath.Join(romFolder, "Other")

	switch expression := ext; expression {
	case ".gba":
		destinationFolder = filepath.Join(romFolder, "gba")
	}
	

	return moveFile(cleanPath, destinationFolder)
}

func moveFile(sourcePath, destinationFolder string) error {
	if _, err := os.Stat(sourcePath); os.IsNotExist(err) {
		return fmt.Errorf("source file does not exist: %s", sourcePath)
	}

	fileName := filepath.Base(sourcePath)

	destinationPath := filepath.Join(destinationFolder, fileName)

	if err := os.MkdirAll(destinationFolder, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %v", err)
	}

	if err := os.Rename(sourcePath, destinationPath); err != nil {
		return fmt.Errorf("failed to move file: %v", err)
	}

	return nil;
}