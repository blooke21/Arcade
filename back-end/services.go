// Business logic, utility functions
package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func returnRomList() ([]byte, error) {
	fmt.Println("Retrieving ROM database...")
	database, err := getAllRomDatabase()
	if err != nil {
		return nil, fmt.Errorf("failed to get ROM database: %v", err)
	}

	return database, nil
}

/**
 * Adds a new ROM entry to the rom_database.json file
 * @param fileMap A map containing ROM details (e.g., type, fileName, source, image)
 * @return The ID of the newly added ROM and an error if any
 */
func addRomDatabase(fileMap map[string]string) (int, error) {
	dataFilePath := filepath.Join(romDB)
	
	var roms []map[string]string
	
	database, err := getAllRomDatabase();
	if err != nil {
		return -1, fmt.Errorf("failed to get ROM database: %v", err)
	}

	databaseErr := json.Unmarshal(database, &roms)
	if databaseErr != nil {
		return -1, fmt.Errorf("failed to decode database: %v", databaseErr)
	}
	
	fmt.Printf("Current ROM database: %+v\n", roms)

	fmt.Printf("Adding new ROM entry: %+v\n", fileMap)
	
	// Add new ROM
	roms = append(roms, fileMap)
	
	// Write back to file
	jsonData, err := json.MarshalIndent(roms, "", "  ")

	romID := len(roms) - 1;

	if err != nil {
		return romID, fmt.Errorf("failed to encode database: %v", err)
	}
	
	if err := os.WriteFile(dataFilePath, jsonData, 0644); err != nil {
		return romID, fmt.Errorf("failed to write database file: %v", err)
	}
	
	return romID, nil
}

/**
 * Retrieves the ROM database from rom_database.json
 * @return The content of the ROM database file and an error if any
 */
func getAllRomDatabase() ([]byte, error) {
	dataFilePath := filepath.Join(romDB)

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

/**
 * Retrieves a ROM entry by its ID
 * @param romID The ID of the ROM to retrieve
 * @return A map containing the ROM details and an error if any
 */
func getRomByID(romID int) (map[string]string, error) {
	database, err := getAllRomDatabase();
	if err != nil {
		return nil, fmt.Errorf("failed to get ROM database: %v", err)
	}

	var roms []map[string]string
	databaseErr := json.Unmarshal(database, &roms)
	if databaseErr != nil {
		return nil, fmt.Errorf("failed to decode database: %v", databaseErr)
	}

	if romID < 0 || romID >= len(roms) {
		return nil, fmt.Errorf("ROM ID %d out of range", romID)
	}

	return roms[romID], nil;
}

/**
 * Retrieves the number of ROM entries in the database
 * @return The count of ROM entries and an error if any
 */
func getAllRomDatabaseLength() (int, error) {
	database, err := getAllRomDatabase();
	if err != nil {
		return -1, fmt.Errorf("failed to get ROM database: %v", err)
	}

	var roms []map[string]string
	databaseErr := json.Unmarshal(database, &roms)
	if databaseErr != nil {
		return -1, fmt.Errorf("failed to decode database: %v", databaseErr)
	}

	return len(roms), nil;
}

/**
 * Moves a file from sourcePath to destinationPath, creating destinationFolder if it doesn't exist
 * @param sourcePath The path of the file to move
 * @param destinationFolder The folder to move the file into
 * @param destinationPath The full path (including filename) for the moved file
 * @return An error if the operation fails
 */
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