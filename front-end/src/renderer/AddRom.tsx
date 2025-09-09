import React, { useState } from 'react';
import axios from 'axios';

function AddRom() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [moveStatus, setMoveStatus] = useState<boolean>(false);
  const [isLoading, setIsLoading] = useState(false);

  const handleFileChange = async () => {
      handleSelectFile();
      if (selectedFile) {
        handleMoveFile(selectedFile);
      }
  }

  const handleSelectFile = async () => {
    try {
      // Use the exposed Electron API to open the file dialog
      const filePath = await (window as any).electronAPI.openFileDialog();

      if (filePath) {
        setSelectedFile(filePath);
        handleMoveFile(filePath);
      }
    } catch (error) {
      console.error('Error selecting file:', error);
      setSelectedFile(null);
    }
  };

  const handleMoveFile = async (selectedFile: string) => {
    setIsLoading(true);

    try {
      // Send the file path to backend
      const response = await axios.post('http://localhost:8080/api/add-rom', {
        sourcePath: selectedFile,
      });

      setMoveStatus(!!response.data);
      setSelectedFile(null);
    } catch (error) {
      console.log((error as any).status);
      console.log('hit on try to pass to api');
      setMoveStatus(false);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="file-selector">
      <h2>File Mover</h2>

      <div className="controls">
        <button onClick={handleFileChange} disabled={isLoading}>
          Select File
        </button>
      </div>

      {selectedFile && (
        <div className="selected-file">
          <p>Selected file: {selectedFile}</p>
        </div>
      )}

      {moveStatus && (
        <div className="status">
          <p>{moveStatus}</p>
        </div>
      )}
    </div>
  );
}

export default AddRom;
