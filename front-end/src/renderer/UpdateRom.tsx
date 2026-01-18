import React, { useState } from 'react';
import axios from 'axios';
import Popup from './UpdatePopup';
/**
 * shows a list of roms to update
 * and handles showing the update popup when a rom is selected
 * @returns
 */
function UpdateRom() {
  /**
   * Fetch ROMs from the backend
   */
  const getRoms = async () => {
    let res;

    try {
      res = await axios.post('http://localhost:8080/api/roms');
      if (res) {
        console.log(res.data);
        return res.data;
      }
    } catch (error) {
      console.error('Error getting roms:', error);
      setErrorMessage('Error retrieving ROMs from the database.');
      setSelectedFile(null);
    }
  };

  type Rom = {
    fileName: string;
    // add other fields here as needed
  }

  const [roms, setRoms] = useState<Array<Rom>>([]);
  const [selectedFile, setSelectedFile] = useState<Rom | null>(null);
  const [errorMessage, setErrorMessage] = useState('');
  // Popup state
  const [PopupOpen, setPopupOpen] = useState(false);
  const [PopupFormData, setPopupFormData] = useState({});

  /**
   * Fetch ROMs on component mount
   */
  React.useEffect(() => {
    getRoms().then((data) => {
      if (data) {
        setRoms(data);
      }
    });
  }, []);

  const handleFileChange = async (selectedRom: any) => {
    //ODOT why we doing this again?
    setSelectedFile(selectedRom);

    setPopupOpen(true);
  };

  return (
    <div className="file-selector">
      <h2>Update Rom</h2>

      {roms.length > 0 ? (
        <div>
          <h3>Select a ROM to update:</h3>
          <li>
            {roms.map((rom) => (
              <ul key={rom.fileName}>
                {rom.fileName}
                <button
                  onClick={() => {
                    handleFileChange(rom);
                  }}
                >
                  Update
                </button>
              </ul>
            ))}
          </li>
        </div>
      ) : (
        <p>{errorMessage}</p>
      )}

      {selectedFile && PopupOpen && (
        <Popup
          key = {selectedFile.fileName}
          isOpen={PopupOpen}
          onClose={() => setPopupOpen(false)}
          onSubmit={(formData: any) => {
            setPopupFormData(formData);
            setPopupOpen(false);
          }}
          rom={selectedFile == null ? roms[0] : selectedFile}
        />
      )}
    </div>
  );
}

export default UpdateRom;
