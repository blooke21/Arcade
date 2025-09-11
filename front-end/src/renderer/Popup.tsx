import React from "react";

interface PopupProps {
  message: string;
}

const Popup = ({ message }: PopupProps) => {
  return (
    <div className="popup">
      <div className="popup-content">
        <p>{message}</p>
      </div>
    </div>
  );
}

export default Popup;