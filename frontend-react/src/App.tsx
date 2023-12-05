import React from 'react';
import './App.css';
import { Route, Routes } from 'react-router-dom';

import BasicContainer from "./containers/basic";
import HomeContent from "./pages/home";
import PDFContent from "./pages/pdf";
import AboutContent from "./pages/about";
import useUser from "./auth/user";

const App: React.FC = () => {
  const user = useUser()

  return (
    <>
      <Routes>
        <Route index element={<BasicContainer user={user}><HomeContent /></BasicContainer>} />
        <Route path="/home" element={<BasicContainer user={user}><HomeContent /></BasicContainer>} />
        <Route path="/pdf" element={<BasicContainer user={user}><PDFContent user={user} /></BasicContainer>} />
        <Route path="/about" element={<BasicContainer user={user}><AboutContent /></BasicContainer>} />
      </Routes>
    </>
  );
}

export default App;
