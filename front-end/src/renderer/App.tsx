import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import { useEffect, useState } from 'react';
import AddRom from './AddRom';
import UpdateRom from './UpdateRom';


export default function App() {
  return (
    <div>
      <AddRom />
      <UpdateRom />
    </div>
  );
}
