/**
 * Display a single ROM's information and take input for updating
 * Takes in new name and image path
 * image path will then be passed to the backend and stored in the image database
 */

import React, { useState } from 'react';
import { selectFile } from './fileUtils';

interface UpdatePopupProps {
  isOpen: boolean;
  onClose: Function;
  onSubmit: Function;
  rom: any;
}

const Popup = ({ isOpen, onClose, onSubmit, rom }: UpdatePopupProps) => {
  //form state for the inputs passed from update rom
  const [form, setForm] = React.useState({
    name: rom.fileName,
    imagePath: rom.image,
  });

  const [selectedImg, setSelectedImg] = React.useState<string | null>(null);  
  const [isLoading, setIsLoading] = useState(false);

  //sync form state when rom prop changes
  React.useEffect(() => {
    // when rom changes, reset form values to the new rom's data
    setForm({
      name: rom.fileName,
      imagePath: rom.image,
    });
  }, [rom]);

  //handles showing the changes in the form inputs
  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setForm((prev) => ({
      ...prev,
      [name === 'fileName' ? 'name' : name]: value,
    }));
  };

  const handleSubmit = (e: { preventDefault: () => void }) => {
    e.preventDefault();
    onSubmit(form);
    setIsLoading(false);

  };

  const handleFileChange = async () => {
    try {
      const filePath = await selectFile();
    } catch (error) {
      console.error('Error selecting file:', error);
      setSelectedImg(null);
    }
  };

  if (!isOpen || !rom) return null;

  const { name, imagePath } = form;

  return (
    <div className="popup-overlay" onClick={() => onClose()}>
      <div className="popup-content" onClick={(e) => e.stopPropagation()}>
        <h2>Edit Rom</h2>
        <form onSubmit={handleSubmit}>
          <label htmlFor="fileName">Display Name</label>
          <input
            type="text"
            name="fileName"
            placeholder="Display Name"
            value={name}
            onChange={handleChange}
            required
          />
          <img src={imagePath} alt="Current ROM" style={{ width: '200px', height: 'auto', marginBottom: '10px' }} />
          <button onClick={handleFileChange} disabled={isLoading}>
          Add new image
        </button>
          <div className="button-group">
            <button type="submit">Submit</button>
            <button type="button" onClick={() => onClose()}>
              Cancel
            </button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Popup;
