import React from 'react';
import './App.css';
import { Route, Routes } from 'react-router-dom';

import Home from "./pages/home";
import PDF from "./pages/pdf";
import About from "./pages/about";

const App: React.FC = () => {
  return (
    <>
      <Routes>
        <Route index element={<Home />} />
        <Route path="/home" element={<Home />} />
        <Route path="/pdf" element={<PDF />} />
        <Route path="/about" element={<About />} />
      </Routes>
    </>
  );
}

export default App;
