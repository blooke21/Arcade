import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import { useEffect, useState } from 'react';
import FileMover from './FileMover';

function Hello() {
  fetch('http://localhost:8080/api/hello')
    .then((response) => {
      if (!response.ok) {
        throw new Error('Network response was not ok');
      }
      return response.json(); // Assuming the response is in JSON format
    })
    .then((data) => {
      // Handle the API response data
      console.log(data);
    })
    .catch((error) => {
      // Handle any errors
      console.error('There was a problem with the fetch operation:', error);
    });
  return (
    <div>
      <h1>Hello World</h1>
    </div>
  );
}

export default function App() {
  return (
    <Router>
      <Routes>
        {/* <Route path="/" element={<Hello />} /> */}
        <Route path="/" element={<FileMover />} />
      </Routes>
    </Router>
  );
}
