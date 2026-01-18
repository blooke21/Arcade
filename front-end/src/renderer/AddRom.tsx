import React, { useState } from 'react';
import axios from 'axios';
import Popup from './Popup';
import { selectFile } from './fileUtils';

function AddRom() {
  const [selectedFile, setSelectedFile] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [showPopup, setShowPopup] = useState(false);
  const [popupMessage, setMessage] = useState('No file selected.');

  const handleFileChange = async () => {
    try {
      const filePath = await selectFile();

      if (filePath) {
        setSelectedFile(filePath);
        await handleMoveFile(filePath);
      }
    } catch (error) {
      console.error('Error selecting file:', error);
      setSelectedFile(null);
    }
  };

  const handleMoveFile = async (selectedFile: string) => {
    setIsLoading(true);
    let res;

    try {
      // Send the file path to backend
      res = await axios.post('http://localhost:8080/api/add-rom', {
        sourcePath: selectedFile,
      });

      setSelectedFile(null);
      
    } catch (error) {
      console.log((error as any).status);
      console.log('hit on try to pass to api');
    } finally {
      setMessage(res?.data.message || 'File move operation completed.');
      setIsLoading(false);
      setShowPopup(true);
      setTimeout(() => setShowPopup(false), 10000); // Hide popup after 3 seconds
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

      {showPopup && <Popup message={popupMessage} />}
    </div>
  );
}

export default AddRom;
