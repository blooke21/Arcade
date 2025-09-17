import React, { useState } from 'react';
import axios from 'axios';
import Popup from './Popup';
import { json } from 'node:stream/consumers';

function UpdateRom() {
    const [selectedFile, setSelectedFile] = useState(null);

    const handleFileChange = async () => {
        handleSelectdbEntry();
        if (selectedFile) {
            // handleUpdateFile(selectedFile);
        }
    }

    const handleSelectdbEntry = async () => {
        let res;
        let list;

        try {
            res = await axios.post('http://localhost:8080/api/roms');
            if (res) {
                console.log(res.data);
                list = res.data;
                // // Populate a dropdown with the list of ROMs
                console.log(list);
            }
        }
        catch (error) {
            console.error('Error getting roms:', error);
            setSelectedFile(null);
        }
    }

    return (
        <div className="file-selector">
            <h2>Update Rom</h2>

            <div className="controls">
                <button onClick={handleFileChange} disabled={false}>
                    Select ROM to Update
                </button>
            </div>

        </div>
    );
}

export default UpdateRom;