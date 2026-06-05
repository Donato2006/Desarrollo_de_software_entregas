import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import DetalleConcierto from "./pages/DetalleConcierto";
import MisEntradas from "./pages/MisEntradas";
import Register from "./pages/Register";
import MisListasEspera from "./pages/MisListasEspera";

function App() {
  return (
    <BrowserRouter>

      <Routes>

        <Route
          path="/"
          element={<Home />}
        />

        <Route
          path="/login"
          element={<Login />}
        />

        <Route
          path="/conciertos/:id"
          element={<DetalleConcierto />}
        />

        <Route
          path="/mis-entradas"
          element={<MisEntradas />}
        />

        <Route
          path="/register"
          element={<Register />}
        />

        <Route
          path="/mis-listas-espera"
          element={<MisListasEspera />}
        />

      </Routes>

    </BrowserRouter>
  );
}

export default App;