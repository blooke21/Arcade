package com.blook.arcade.helper;

import java.io.File;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;

import javax.swing.JFileChooser;
import javax.swing.filechooser.FileNameExtensionFilter;

public class FileMover {

    public static void main(String[] args) {

        JFileChooser fileChooser = new JFileChooser();
        fileChooser.setDialogTitle("Select File to Move");
        fileChooser.setFileSelectionMode(JFileChooser.FILES_ONLY);

        // Optional: Filter file types
        FileNameExtensionFilter filter = new FileNameExtensionFilter("Text Files", "txt");
        fileChooser.setFileFilter(filter);

        int result = fileChooser.showOpenDialog(null);

        if (result == JFileChooser.APPROVE_OPTION) {
            File selectedFile = fileChooser.getSelectedFile();
            moveFile(selectedFile, "GBA"); //this is where the folder passes 
        } else {
            System.out.println("No file selected.");
        }
    }

    private static void moveFile(File sourceFile, String location) {
        JFileChooser destinationChooser = new JFileChooser();
        // destinationChooser.setDialogTitle("Choose Destination Folder");
        // destinationChooser.setFileSelectionMode(JFileChooser.DIRECTORIES_ONLY);
        File selectedFile = new File("/Users/blakewarnock/Projects/Java/Testing/activeTesting/" + location);
        destinationChooser.setSelectedFile(selectedFile);

        // int result = destinationChooser.showOpenDialog(null);
        // if (result == JFileChooser.APPROVE_OPTION) {
        File destinationFolder = destinationChooser.getSelectedFile();
        Path sourcePath = Paths.get(sourceFile.getAbsolutePath());
        Path destinationPath = Paths.get(destinationFolder.getAbsolutePath(), sourceFile.getName());
        System.out.println("destinationPath = " + destinationPath);
        System.out.println("sourcePAth = " + sourcePath);
        System.out.println("destinationFolder = " + destinationFolder);

        try {
            Files.move(sourcePath, destinationPath, StandardCopyOption.REPLACE_EXISTING);
            System.out.println("File moved successfully to: " + destinationPath);
        } catch (IOException e) {
            System.err.println("Error moving file: " + e.getMessage());
        }
        // } else {
        //     System.out.println("No destination folder selected.");
        // }
    }
}
