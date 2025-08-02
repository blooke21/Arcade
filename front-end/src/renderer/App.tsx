import { MemoryRouter as Router, Routes, Route } from 'react-router-dom';
import './App.css';
import { useEffect, useState } from 'react';
import FileMover from './FileMover';

function Hello() {
  const [message, setMessage] = useState('');

  useEffect(() => {
    fetch('http://localhost:8080/api/hello')
      .then((response) => {
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        return response.json(); // Assuming the response is in JSON format
      })
      .then((data) => {
        setMessage(data.message);
      });
  }, []);
  return (
    <div>
      <h1>Hello from React!</h1>
      <p>{message}</p>
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
