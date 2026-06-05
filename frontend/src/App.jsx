import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import DetalleConcierto from "./pages/DetalleConcierto";
import MisEntradas from "./pages/MisEntradas";
import Register from "./pages/Register";

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

      </Routes>

    </BrowserRouter>
  );
}

export default App;