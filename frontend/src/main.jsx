<<<<<<< HEAD
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./styles/index.css";
import App from "./App.jsx";
=======
import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './styles/index.css'
import App from './App.jsx'
>>>>>>> 94d7605619bbb0b3d5cdb1839265c18d00e14741

createRoot(document.getElementById("root")).render(
  <StrictMode>
    <App />
  </StrictMode>
);