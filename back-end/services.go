// Business logic
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

func addRomDatabase(fileMap map[string]string) error {
	dataFilePath := filepath.Join("..", "rom", "rom_database.json")
	
	var roms []map[string]string
	
	database, err := getRomDatabase();
	if err != nil {
		return fmt.Errorf("failed to get ROM database: %v", err)
	}

	databaseErr := json.Unmarshal(database, &roms)
	if databaseErr != nil {
		return fmt.Errorf("failed to decode database: %v", databaseErr)
	}
	
	fmt.Printf("Current ROM database: %+v\n", roms)

	fmt.Printf("Adding new ROM entry: %+v\n", fileMap)
	
	// Add new ROM
	roms = append(roms, fileMap)
	
	// Write back to file
	jsonData, err := json.MarshalIndent(roms, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to encode database: %v", err)
	}
	
	if err := os.WriteFile(dataFilePath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write database file: %v", err)
	}
	
	return nil
}

func getRomDatabase() ([]byte, error) {
	dataFilePath := filepath.Join("..", "rom", "rom_database.json")

	_, err := os.Stat(dataFilePath);
	if err != nil {
		file := []byte("[]");
		if writeErr := os.WriteFile(dataFilePath, file, 0644); writeErr != nil {
            fmt.Printf("Error creating file: %v\n", writeErr)
            return nil, writeErr
        }
		return file, nil;
	}

	file, err := os.ReadFile(dataFilePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err;
	}
	return file, nil;
}