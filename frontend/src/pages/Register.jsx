import { useState } from "react";
import { useNavigate } from "react-router-dom";
import api from "../services/api";
import "../styles/Register.css";

function Register() {

  const navigate = useNavigate();

  const [nombre, setNombre] = useState("");
  const [correo, setCorreo] = useState("");
  const [password, setPassword] = useState("");

  const [error, setError] = useState("");
  const [mensaje, setMensaje] = useState("");

  const manejarRegistro = async (e) => {

    e.preventDefault();

    if (!nombre.trim() || !correo.trim() || !password.trim()) {
        setError("Todos los campos son obligatorios");
        return;
    }

    if (password.length < 6) {
        setError("La contraseña debe tener al menos 6 caracteres");
        return;
    }

    try {

      await api.post("/register", {
        Nombre: nombre,
        Correo: correo,
        Password: password,
      });

      setError("");
      setMensaje("Usuario registrado correctamente");

      setTimeout(() => {
        navigate("/login");
      }, 1500);

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo registrar el usuario"
      );

    }

  };

  return (
    <div className="register-container">

      <div className="register-box">

        <h1>Crear Cuenta</h1>

        <form onSubmit={manejarRegistro}>

          <input
            type="text"
            placeholder="Nombre"
            value={nombre}
            onChange={(e) => setNombre(e.target.value)}
          />

          <input
            type="email"
            placeholder="Correo electrónico"
            value={correo}
            onChange={(e) => setCorreo(e.target.value)}
          />

          <input
            type="password"
            placeholder="Contraseña"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
          />

          {mensaje && (
            <p className="success">
              {mensaje}
            </p>
          )}

          {error && (
            <p className="error">
              {error}
            </p>
          )}

          <button type="submit">
            Registrarse
          </button>

        </form>

      </div>

    </div>
  );
}

export default Register;