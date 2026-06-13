import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import DetalleConcierto from "./pages/DetalleConcierto";
import MisEntradas from "./pages/MisEntradas";
import Register from "./pages/Register";
import MisListasEspera from "./pages/MisListasEspera";
import Admin from "./pages/Admin";
import ProtectedAdminRoute from "./components/ProtectedAdminRoute";

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
        <Route
          path="/admin"
          element={
            <ProtectedAdminRoute>
              <Admin />
            </ProtectedAdminRoute>
          }
        />

      </Routes>

    </BrowserRouter>
  );
}

export default App;