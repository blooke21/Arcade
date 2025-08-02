import React, { useState } from 'react';
import axios from 'axios';

function FileSelector() {
  const [selectedFile, setSelectedFile] = useState(null);
  const [status, setStatus] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  const handleSelectFile = async () => {
    try {
      // Use the exposed Electron API to open the file dialog
      const filePath = await (window as any).electronAPI.openFileDialog();

      if (filePath) {
        setSelectedFile(filePath);
        setStatus('File selected: ' + filePath);
        handleMoveFile(filePath);
      }
    } catch (error) {
      setStatus('Error selecting file: ' + (error as Error).message);
    }
  };

  const handleMoveFile = async (selectedFile: string) => {
    setIsLoading(true);
    setStatus('Moving file...');

    try {
      // Send the file path to backend
      const response = await axios.post('http://localhost:8080/api/move-file', {
        sourcePath: selectedFile,
      });

      setStatus('Success: ' + response.data ? 'File moved' : 'File not moved');
      setSelectedFile(null);
    } catch (error) {
      console.log((error as any).status);
      console.log('hit on try to pass to api');
      setStatus('Error: ' + (error as Error).message);
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="file-selector">
      <h2>File Mover</h2>

      <div className="controls">
        <button onClick={handleSelectFile} disabled={isLoading}>
          Select File
        </button>
      </div>

      {selectedFile && (
        <div className="selected-file">
          <p>Selected file: {selectedFile}</p>
        </div>
      )}

      {status && (
        <div className="status">
          <p>{status}</p>
        </div>
      )}
    </div>
  );
}

export default FileSelector;
