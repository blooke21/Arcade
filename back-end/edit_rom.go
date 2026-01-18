package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// editRomDatabase edits an existing ROM entry in the rom_database.json file
// @param romID The ID of the ROM to edit
// @param newName The new name for the ROM
// @param imgPath The new image path for the ROM
// @return An error if any
func editRomDatabase(romID string, newName string, imgPath string) error {
	dataFilePath := filepath.Join("..", "rom", "rom_database.json")

	var roms []map[string]string

	database, err := getAllRomDatabase()
	if err != nil {
		return fmt.Errorf("failed to get ROM database: %v", err)
	}

	databaseErr := json.Unmarshal(database, &roms)
	if databaseErr != nil {
		return fmt.Errorf("failed to decode database: %v", databaseErr)
	}

	fmt.Printf("Current ROM database: %+v\n", roms)

	// Find and update the ROM entry
	for _, rom := range roms {
		if rom["id"] == romID {
			rom["fileName"] = newName
			rom["image"] = imgPath
			break
		}
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