import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import { useEffect, useState } from 'react';
import AddRom from './AddRom';


export default function App() {
  return (
    <Router>
      <Routes>
        {/* <Route path="/" element={<Hello />} /> */}
        <Route path="/" element={<AddRom />} />
      </Routes>
    </Router>
  );
}
