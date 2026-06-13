import { useEffect, useState } from "react";
import { Navigate } from "react-router-dom";
import api from "../services/api";
import "../styles/Admin.css";

function Admin() {

  const rol = localStorage.getItem("rol");
  const token = localStorage.getItem("token");

  const [conciertos, setConciertos] = useState([]);

  const [nombre, setNombre] = useState("");
  const [fecha, setFecha] = useState("");
  const [lugar, setLugar] = useState("");
  const [cupos, setCupos] = useState("");

  const [editandoId, setEditandoId] = useState(null);

  const [mensaje, setMensaje] = useState("");
  const [error, setError] = useState("");

  const cargarConciertos = async () => {

    try {

      const response = await api.get("/conciertos");

      setConciertos(response.data);

    } catch {

      setError("No se pudieron cargar los conciertos");

    }

  };

  useEffect(() => {

    let activo = true;

    const obtenerConciertos = async () => {

      try {

        const response = await api.get("/conciertos");

        if (activo) {
          setConciertos(response.data);
        }

      } catch {

        if (activo) {
          setError("No se pudieron cargar los conciertos");
        }

      }

    };

    obtenerConciertos();

    return () => {
      activo = false;
    };

  }, []);

  const crearConcierto = async (e) => {

    e.preventDefault();

    try {

      await api.post(
        "/conciertos",
        {
          Nombre: nombre,
          Fecha: new Date(fecha).toISOString(),
          Lugar: lugar,
          CupoTotal: Number(cupos),
          CuposDisponibles: Number(cupos),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setNombre("");
      setFecha("");
      setLugar("");
      setCupos("");

      setMensaje("Concierto creado correctamente");
      setError("");

      cargarConciertos();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo crear el concierto"
      );

    }

  };

  const actualizarConcierto = async (e) => {

    e.preventDefault();

    try {

      await api.put(
        `/conciertos/${editandoId}`,
        {
          Nombre: nombre,
          Fecha: new Date(fecha).toISOString(),
          Lugar: lugar,
          CupoTotal: Number(cupos),
          CuposDisponibles: Number(cupos),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setMensaje("Concierto actualizado correctamente");
      setError("");

      setEditandoId(null);

      setNombre("");
      setFecha("");
      setLugar("");
      setCupos("");

      cargarConciertos();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo actualizar el concierto"
      );

    }

  };

  const editarConcierto = (concierto) => {

    setEditandoId(concierto.ID);

    setNombre(concierto.Nombre);

    setFecha(
      new Date(concierto.Fecha)
        .toISOString()
        .slice(0, 16)
    );

    setLugar(concierto.Lugar);

    setCupos(concierto.CupoTotal);

    window.scrollTo({
      top: 0,
      behavior: "smooth",
    });

  };

  const cancelarEdicion = () => {

    setEditandoId(null);

    setNombre("");
    setFecha("");
    setLugar("");
    setCupos("");

  };

  const eliminarConcierto = async (id) => {

    if (!window.confirm("¿Eliminar concierto?")) {
      return;
    }

    try {

      await api.delete(
        `/conciertos/${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setMensaje("Concierto eliminado correctamente");
      setError("");

      cargarConciertos();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo eliminar el concierto"
      );

    }

  };

  if (rol !== "admin") {
    return <Navigate to="/" />;
  }

  return (
    <div className="admin-container">

      <h1>Panel Administrador</h1>

      {mensaje && (
        <p className="admin-success">
          {mensaje}
        </p>
      )}

      {error && (
        <p className="admin-error">
          {error}
        </p>
      )}

      <form
        className="admin-form"
        onSubmit={
          editandoId
            ? actualizarConcierto
            : crearConcierto
        }
      >

        <input
          type="text"
          placeholder="Nombre"
          value={nombre}
          onChange={(e) => setNombre(e.target.value)}
          required
        />

        <input
          type="datetime-local"
          value={fecha}
          onChange={(e) => setFecha(e.target.value)}
          required
        />

        <input
          type="text"
          placeholder="Lugar"
          value={lugar}
          onChange={(e) => setLugar(e.target.value)}
          required
        />

        <input
          type="number"
          placeholder="Cupos"
          value={cupos}
          onChange={(e) => setCupos(e.target.value)}
          required
        />

        <button type="submit">

          {editandoId
            ? "Actualizar Concierto"
            : "Crear Concierto"}

        </button>

        {editandoId && (

          <button
            type="button"
            onClick={cancelarEdicion}
          >
            Cancelar Edición
          </button>

        )}

      </form>

      <div className="admin-grid">

        {conciertos.map((concierto) => (

          <div
            key={concierto.ID}
            className="admin-card"
          >

            <h3>
              {concierto.Nombre}
            </h3>

            <p>
              📍 {concierto.Lugar}
            </p>

            <p>
              🎫 Cupos disponibles:
              {" "}
              {concierto.CuposDisponibles}
            </p>

            <button
              onClick={() =>
                editarConcierto(concierto)
              }
            >
              Editar
            </button>

            <button
              onClick={() =>
                eliminarConcierto(concierto.ID)
              }
            >
              Eliminar
            </button>

          </div>

        ))}

      </div>

    </div>
  );
}

export default Admin;