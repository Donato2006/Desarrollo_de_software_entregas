import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";
import "../styles/login.css";

function Login() {
  const [correo, setCorreo] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState("");

  const navigate = useNavigate();

  const manejarLogin = async (e) => {
    e.preventDefault();

    try {
      const response = await api.post("/login", {
        Correo: correo,
        Password: password,
      });

      localStorage.setItem("token", response.data.token);
      localStorage.setItem("rol", response.data.rol);

      navigate("/");
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
            required
          />

          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => {
              setPassword(e.target.value);
              setError("");
            }}
            required
          />

          {error && (
            <p className="error">
              {error}
            </p>
          )}

          <button type="submit">
            Ingresar
          </button>

          <p className="register-link">
            ¿No tenés cuenta?
            {" "}
            <span onClick={() => navigate("/register")}>
              Registrate
            </span>
          </p>

        </form>

      </div>

    </div>
  );
}

export default Login;