/**
 * Display a single ROM's information and take input for updating
 * Takes in new name and image path
 * image path will then be passed to the backend and stored in the image database
 */

import React from 'react';

interface UpdatePopupProps {
  isOpen: boolean;
  onClose: Function;
  onSubmit: Function;
  rom: any;
}

const Popup = ({ isOpen, onClose, onSubmit, rom }: UpdatePopupProps) => {
  const [form, setForm] = React.useState({
    name: rom.fileName,
    imagePath: '',
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({
      ...form,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = (e: { preventDefault: () => void; }) => {
    e.preventDefault();
    onSubmit(form);
  };

  if(!isOpen || !rom) return null;

  const {name, imagePath} = form;

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
        <label htmlFor="imagePath">Image Path</label>
          <input
            type="text"
            name="imagePath"
            placeholder="Image Path"
            value={imagePath}
            onChange={handleChange}
            required
          />
          <div className="button-group">
            <button type="submit">Submit</button>
            <button type="button" onClick={() => onClose()}>Cancel</button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default Popup;
