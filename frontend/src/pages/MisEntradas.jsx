import { useEffect, useState } from "react";
import api from "../services/api";
import "../styles/MisEntradas.css";

function MisEntradas() {

  const [entradas, setEntradas] = useState([]);
  const [error, setError] = useState("");
  const [mensaje, setMensaje] = useState("");

  const cargarEntradas = async () => {

    try {

      const token = localStorage.getItem("token");

      const response = await api.get(
        "/mis-entradas",
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setEntradas(response.data || []);
      setError("");

    } catch (err) {

      setEntradas([]);

      setError(
        err.response?.data?.error ||
        "No se pudieron cargar las entradas"
      );

    }

  };

  useEffect(() => {

    const obtenerEntradas = async () => {
      await cargarEntradas();
    };

    obtenerEntradas();

  }, []);

  const cancelarEntrada = async (id) => {

    try {

      const token = localStorage.getItem("token");

      await api.delete(
        `/entradas/${id}`,
        {
          headers: {
            Authorization: `Bearer ${token}`,
          },
        }
      );

      setMensaje("Entrada cancelada correctamente");
      setError("");

      await cargarEntradas();

    } catch (err) {

      setMensaje("");

      setError(
        err.response?.data?.error ||
        "No se pudo cancelar la entrada"
      );

    }

  };

  return (
    <div className="mis-entradas-container">

      <h1>Mis Entradas</h1>

      {mensaje && (
        <p className="mensaje-exito">
          {mensaje}
        </p>
      )}

      {error && (
        <p className="mensaje-error">
          {error}
        </p>
      )}

      {entradas.length === 0 ? (

        <p>No tenés entradas compradas</p>

      ) : (

        <div className="entradas-grid">

          {entradas.map((entrada) => (

            <div
              key={entrada.ID}
              className="entrada-card"
            >

              <h3>
                Entrada #{entrada.ID}
              </h3>

              <p>
                Estado: {entrada.Estado}
              </p>

              <p>
                Usuario ID: {entrada.UsuarioID}
              </p>

              <p>
                Concierto ID: {entrada.ConciertoID}
              </p>

              <button
                className="cancelar-btn"
                onClick={() =>
                  cancelarEntrada(entrada.ID)
                }
              >
                Cancelar Entrada
              </button>

            </div>

          ))}

        </div>

      )}

    </div>
  );
}

export default MisEntradas;