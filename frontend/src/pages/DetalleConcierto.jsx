import { useEffect, useState } from "react";
import { useParams } from "react-router-dom";
import api from "../services/api";
import "../styles/DetalleConcierto.css";

function DetalleConcierto() {

  const { id } = useParams();

  const [concierto, setConcierto] = useState(null);

  const [mensaje, setMensaje] = useState("");
  const [error, setError] = useState("");

  const token = localStorage.getItem("token");

  useEffect(() => {

    const cargarConcierto = async () => {

      try {

        const response = await api.get(`/conciertos/${id}`);

        setConcierto(response.data);

      } catch {

        setError("No se pudo cargar el concierto");

      }

    };

    cargarConcierto();

  }, [id]);

  const recargarConcierto = async () => {

    const response = await api.get(`/conciertos/${id}`);

    setConcierto(response.data);

  };

  const comprarEntrada = async () => {

    try {

      await api.post(
        "/entradas",
        {
          ConciertoID: Number(id),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setError("");
      setMensaje("Entrada comprada correctamente");

      await recargarConcierto();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo comprar la entrada"
      );

    }

  };

  const unirseListaEspera = async () => {

    try {

      await api.post(
        "/lista-espera",
        {
          ConciertoID: Number(id),
        },
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setError("");
      setMensaje("Te anotaste correctamente en la lista de espera");

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo ingresar a la lista de espera"
      );

    }

  };

  if (!concierto) {
    return (
      <div className="detalle-container">
        <h2>Cargando...</h2>
      </div>
    );
  }

  const fecha = new Date(concierto.Fecha).toLocaleDateString(
    "es-AR",
    {
      day: "2-digit",
      month: "long",
      year: "numeric",
    }
  );

  const hora = new Date(concierto.Fecha).toLocaleTimeString(
    "es-AR",
    {
      hour: "2-digit",
      minute: "2-digit",
    }
  );

  return (
    <div className="detalle-container">

      <div className="detalle-card">

        <h1 className="detalle-titulo">
          {concierto.Nombre}
        </h1>

        <p>📍 {concierto.Lugar}</p>

        <p>📅 {fecha}</p>

        <p>🕘 {hora}</p>

        <p>
          🎫 Cupos disponibles:
          {" "}
          {concierto.CuposDisponibles}
        </p>

        {mensaje && (
          <p className="detalle-exito">
            {mensaje}
          </p>
        )}

        {error && (
          <p className="detalle-error">
            {error}
          </p>
        )}

        {!token && (
          <p>
            Iniciá sesión para comprar entradas o unirte a la lista de espera.
          </p>
        )}

        {token && (
          <div className="detalle-botones">

            <button
              className="detalle-boton"
              onClick={comprarEntrada}
            >
              Comprar Entrada
            </button>

            {concierto.CuposDisponibles === 0 && (
              <button
                className="detalle-boton"
                onClick={unirseListaEspera}
              >
                Unirse a Lista de Espera
              </button>
            )}

          </div>
        )}

      </div>

    </div>
  );
}

export default DetalleConcierto;