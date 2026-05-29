import { useState } from "react";
import api from "../services/api";
import "../styles/login.css";

function Login() {

  const [correo, setCorreo] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const manejarLogin = async (e) => {

    e.preventDefault();

    try {

      const response = await api.post("/login", {
        Correo: correo,
        Password: password,
      });

      localStorage.setItem("token", response.data.token);

      window.location.href = "/";

    } catch (err) {

      setError("Correo o contraseña incorrectos");

    }
  };

  return (
    <div className="login-container">

      <div className="login-box">

        <h1>Iniciar Sesión</h1>

        <form onSubmit={manejarLogin}>

          <input
            type="email"
            placeholder="Correo electrónico"
            value={correo}
            onChange={(e) => {
              setCorreo(e.target.value);
              setError("");
            }}
          />

          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => {
              setPassword(e.target.value);
              setError("");
            }}
          />

          {error && <p className="error">{error}</p>}

          <button type="submit">
            Ingresar
          </button>

        </form>

      </div>

    </div>
  );
}

export default Login;