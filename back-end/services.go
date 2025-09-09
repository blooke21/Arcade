// Business logic
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"encoding/json"
)

func updateRomDatabase(fileMap map[string]string) error {
	dataFilePath := filepath.Join("..", "rom", "rom_database.json")
	
	var roms []map[string]string
	
	// Try to read existing file
	if data, err := os.ReadFile(dataFilePath); err == nil {
		// File exists, try to parse it
		if len(data) > 0 {
			if err := json.Unmarshal(data, &roms); err != nil {
				return fmt.Errorf("failed to parse existing database: %v", err)
			}
		}
	} else if !os.IsNotExist(err) {
		// Some other error occurred
		return fmt.Errorf("failed to read database file: %v", err)
	}
	// If file doesn't exist, roms remains as zero value (empty slice)
	
	// Check for duplicate by fileName
	for _, rom := range roms {
		if rom["fileName"] == fileMap["fileName"] {
			return nil // Already exists, skip
		}
	}
	
	// Add new ROM
	roms = append(roms, fileMap)
	
	// Ensure directory exists
	if err := os.MkdirAll(filepath.Dir(dataFilePath), 0755); err != nil {
		return fmt.Errorf("failed to create directory: %v", err)
	}
	
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