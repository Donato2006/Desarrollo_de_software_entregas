import { BrowserRouter, Routes, Route } from "react-router-dom";

import Home from "./pages/Home";
import Login from "./pages/Login";
import DetalleConcierto from "./pages/DetalleConcierto";

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

      </Routes>

    </BrowserRouter>
  );
}

export default App;