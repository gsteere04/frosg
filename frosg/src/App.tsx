import './App.css'
import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import Donation from './pages/Donation';
import Projects from './pages/Projects';
import Media from './pages/Media';
import Contact from './pages/Contact';
import Header from './components/Header';
import Footer from './components/Footer';

const App: React.FC = () => {
  return (
    <Router>
      <Header />
      <main>
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/donations" element={<Donation />} />
          <Route path="/projects" element={<Projects />} />
          <Route path="/media" element={<Media />} />
          <Route path="/contact" element={<Contact />} />
        </Routes>
      </main>
      <Footer />
    </Router>
  );
};

export default App
